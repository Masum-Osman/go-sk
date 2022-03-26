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

type Book struct {
	ID       int
	Title    string
	Price    float32
	IsbnNo   string
	Author   Author
	Comments []Comment
}

type Author struct {
	Id        int
	Name      string
	Biography string
	Books     []int
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

func populateBooks() []Book {
	authors := &Author{
		Id:        1,
		Name:      "Dan Brown",
		Biography: "All time Best Seller",
		Books:     []int{1},
	}
	book := Book{
		ID:     1,
		Title:  "Angels and Demons",
		Price:  500.00,
		IsbnNo: "978-3-16-148410-0",
		Author: *authors,
	}

	var books []Book
	books = append(books, book)

	return books
}

func main() {

	// tutorials := populate()
	books := populateBooks()

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

	var booksType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Books",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
				"isbn_no": &graphql.Field{
					Type: graphql.String,
				},
				"authors": &graphql.Field{
					Type: authorType,
				},
			},
		},
	)

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get tutorial by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Full Tutorial List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
		"books": &graphql.Field{
			Type:        graphql.NewList(booksType),
			Description: "Get Full Book List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
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
		list {
            id
            title
            comment {
                body
            }
            author {
                Name
                Tutorials
            }
        }
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
