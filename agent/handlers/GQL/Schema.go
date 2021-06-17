package GQL

import (
	"github.com/graphql-go/graphql"
)

var (
	GQLSchema graphql.Schema
	containerMetadataType *graphql.Object
)
func init() {

	/*
	   Create Query object type with field "DockerID" with type dockerid by using GraphQLObjectTypeConfig:
	       - Name: name of object type
	       - Fields: a map of fields by using GraphQLFields
	   Setup type of field use GraphQLFieldConfig to define:
	       - Type: type of field
	       - Args: arguments to query with current field
	       - Resolve: function to query data using params from [Args] and return value with current type
	*/

	containerMetadataType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ECS_CONTAINER_METADATA",
		Description: "The version 4 Container Metadata",
		Fields: graphql.Fields{
			"DockerID": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "MockID", nil
				},
			},

		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"ECS_CONTAINER_METADATA": &graphql.Field{
				Type: containerMetadataType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "MockContainer", nil
				},
			},
		},
	})
	GQLSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}