package main

/*
how to use:

postman

url   : http://localhost:3000/graphql
method: post
body  : choose "GraphQL"
body content:
{
  post(id: 6) {
    userId
    id
    body
    title
    comments {
      id
      email
      name
    }
  }
}
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"io/ioutil"
	"log"
	"net/http"
)

// Post is a post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Comment is a comment
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(
				createPostType(
					createCommentType(),
				),
			),
		),
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func createQueryType(postType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"post": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching post with id: %d", v)
					return fetchPostByiD(v)
				},
			},
			//"comment": &graphql.Field{
			//	Type: postType,
			//	Args: graphql.FieldConfigArgument{
			//		"id": &graphql.ArgumentConfig{},
			//	},
			//	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			//		id := p.Args["id"]
			//		v, _ := id.(int)
			//		log.Printf("fetching comment with id: %d", v)
			//		return fetchCommentByiD(v)
			//	},
			//},
		}}
}

func createPostType(commentType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post, _ := p.Source.(*Post)
					log.Printf("fetching comments of post with id: %d", post.ID)
					return fetchCommentsByPostID(post.ID)
				},
			},
		},
	})
}

func createCommentType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"postId": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func fetchPostByiD(id int) (*Post, error) {
	//resp, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d", id))
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != 200 {
	//	return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	//}
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, errors.New("could not read data")
	//}
	result := Post{
		UserID: 6,
		ID:     6,
		Title:  "good good study",
		Body:   "day day up",
	}
	//err = json.Unmarshal(b, &result)
	//if err != nil {
	//	return nil, errors.New("could not unmarshal data")
	//}
	return &result, nil
}

func fetchCommentsByPostID(id int) ([]Comment, error) {
	resp, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d/comments", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read data")
	}
	result := []Comment{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	return result, nil
}

func fetchCommentByiD(id int) (*Comment, error) {
	//resp, err := http.Get(fmt.Sprintf("http://jsonplaceholder.typicode.com/posts/%d", id))
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != 200 {
	//	return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	//}
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, errors.New("could not read data")
	//}
	result := Comment{
		PostID: 6,
		ID:     5,
		Name:   "什么玩意儿",
		Email:  "261337699@qq.com",
		Body:   "文明看球",
	}
	return &result, nil
}
