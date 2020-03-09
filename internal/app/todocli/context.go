package todocli

import (
	todov1beta1 "github.com/prasetyowira/loadsmile/.gen/api/proto/todo/v1beta1"
)

type context struct {
	client todov1beta1.TodoListClient
}

func (c *context) GetTodoClient() todov1beta1.TodoListClient {
	return c.client
}
