package loadsmile_app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/goph/idgen/ulidgen"
	"github.com/gorilla/mux"
	appkitendpoint "github.com/sagikazarmark/appkit/endpoint"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	"github.com/sagikazarmark/kitx/correlation"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	kitxtransport "github.com/sagikazarmark/kitx/transport"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"
	"google.golang.org/grpc"

	// "github.com/prasetyowira/loadsmile/internal/app/mga/httpbin"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo/lunchadapter"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo/lunchadapter/ent"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo/lunchadapter/ent/migrate"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo/tododriver"
	// "github.com/prasetyowira/loadsmile/internal/app/mga/todo/todogen"

	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/httpbin"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/lunchadapter/ent/migrate"
	"github.com/prasetyowira/loadsmile/internal/app/loadsmile-app/lunch/tododriver"
)

const todoTopic = "todo"

// InitializeApp initializes a new HTTP and a new gRPC application.
func InitializeApp(
	httpRouter *mux.Router,
	grpcServer *grpc.Server,
	publisher message.Publisher,
	storage string,
	db *sql.DB,
	logger Logger,
	errorHandler ErrorHandler,
) {
	endpointMiddleware := []endpoint.Middleware{
		correlation.Middleware(),
		opencensus.TraceEndpoint("", opencensus.WithSpanName(func(ctx context.Context, _ string) string {
			name, _ := kitxendpoint.OperationName(ctx)

			return name
		})),
		appkitendpoint.LoggingMiddleware(logger),
	}

	transportErrorHandler := kitxtransport.NewErrorHandler(errorHandler)

	httpServerOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transportErrorHandler),
		kithttp.ServerErrorEncoder(kitxhttp.NewJSONProblemErrorEncoder(appkithttp.NewDefaultProblemConverter())),
		kithttp.ServerBefore(correlation.HTTPToContext()),
	}

	{

		client := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
		err := client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			panic(err)
		}

		var store todo.Store = lunchadapter.NewEntStore(client)

		service := todo.NewService(
			ulidgen.NewGenerator(),
			store,
			todogen.NewEventDispatcher(eventBus),
		)
		service = tododriver.LoggingMiddleware(logger)(service)
		service = tododriver.InstrumentationMiddleware()(service)

		endpoints := tododriver.MakeEndpoints(
			service,
			kitxendpoint.Combine(endpointMiddleware...),
		)

		tododriver.RegisterHTTPHandlers(
			endpoints,
			httpRouter.PathPrefix("/todos").Subrouter(),
			kitxhttp.ServerOptions(httpServerOptions),
		)

		httpRouter.PathPrefix("/graphql").Handler(tododriver.MakeGraphQLHandler(endpoints, errorHandler))
	}

	httpRouter.PathPrefix("/httpbin").Handler(http.StripPrefix(
		"/httpbin",
		httpbin.MakeHTTPHandler(logger.WithFields(map[string]interface{}{"module": "httpbin"})),
	))
}
