package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

func populate() []Tutorial {
	author := &Author{Name: "E Forbes", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQL",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "Good"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

func main() {

	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorials": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					Type: authorType,
				},
				"comment": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get All tutorials",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
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
