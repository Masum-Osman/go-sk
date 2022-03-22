package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schamaConfig := graphql.SchemaConfig{Query: graphql.NewObject((rootQuery))}
	schema, err := graphql.NewSchema(schamaConfig)

	if err != nil {
		fmt.Println("Err in Graph Schema")
	}

	query := `
	{
		hello
	}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		fmt.Println(r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
