/*
 * Lunch API
 *
 * Manage Lunch
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type Recipe struct {

	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Ingredients []Ingredient `json:"ingredients,omitempty"`
}
