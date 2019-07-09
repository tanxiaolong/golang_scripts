package main

import (
    "github.com/graphql-go/graphql"
    "errors"
)

type Goods struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64`json:"price"`
    Url   string `json:"url"`
}

var goodsType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Goods",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.String,
            },
            "name": &graphql.Field{
                Type: graphql.String,
            },
            "price": &graphql.Field{
                Type: graphql.Float,
            },
            "url": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)
var goodsListType = graphql.NewList(goodsType)
