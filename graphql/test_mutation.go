package main

import (
	"errors"
	"github.com/graphql-go/graphql"
)

var goodsInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "goodsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// 无需处理参数
			"goodsList": &graphql.Field{
				Type: goodsListType,
				// 处理结构体的回调函数，直接返回处理完成的结构体即可
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return result, err
				},
			},
			// 参数是id
			"goods": &graphql.Field{
				Type: goodsType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 获取参数
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return result, nil
					}
					err := errors.New("Field 'goods' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	},
)
