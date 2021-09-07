package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type Phone struct {
	Number string `json:"number"`         // номер телефона
	Name   string `json:"name,omitempty"` // имя владельца (необязательное)
}

var phones = []Phone{
	{
		Name:   "Ivan",
		Number: "74212340077",
	},
	{
		Name:   "John",
		Number: "18002003122",
	},
}

var phoneType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Phone",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"number": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single phone by number
			   http://localhost:8080/phone?query={check(number:"74212340077"){name,number}}
			*/
			"check": &graphql.Field{
				Type:        phoneType,
				Description: "Get phone by number",
				Args: graphql.FieldConfigArgument{
					"number": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					number, ok := p.Args["number"].(string)
					if ok {
						for _, phone := range phones {
							if phone.Number == number {
								return phone, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) phones list
			   http://localhost:8080/phone?query={list{name,number}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(phoneType),
				Description: "Get phones list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return phones, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/phone", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Println("can't send HTTP-response")
		}
	})

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
