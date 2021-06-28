package GQL

import (
	"github.com/aws/amazon-ecs-agent/agent/api"
	apicontainer "github.com/aws/amazon-ecs-agent/agent/api/container"
	apitask "github.com/aws/amazon-ecs-agent/agent/api/task"
	"github.com/aws/amazon-ecs-agent/agent/engine/dockerstate"
	"github.com/aws/amazon-ecs-agent/agent/handlers/utils"
	v3 "github.com/aws/amazon-ecs-agent/agent/handlers/v3"
	"github.com/aws/amazon-ecs-agent/agent/stats"
	"github.com/cihub/seelog"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"log"
)

var ContainerMetadataPath = "/graphql/" + utils.ConstructMuxVar(v3.V3EndpointIDMuxName, utils.AnythingButSlashRegEx)

func CreateSchema(
	state dockerstate.TaskEngineState,
	ecsClient api.ECSClient,
	statsEngine stats.Engine,
	cluster string,
	availabilityZone string,
	containerInstanceArn string) (graphql.Schema){
	/*
	   Create Query object type with field "DockerID" with type dockerid by using GraphQLObjectTypeConfig:
	       - Name: name of object type
	       - Fields: a map of fields by using GraphQLFields
	   Setup type of field use GraphQLFieldConfig to define:
	       - Type: type of field
	       - Args: arguments to query with current field
	       - Resolve: function to query data using params from [Args] and return value with current type
	*/

	//containerMetadataType = graphql.NewObject(graphql.ObjectConfig{
	//	Name:        "ECS_CONTAINER_METADATA",
	//	Description: "The version 4 Container Metadata",
	//	Fields: graphql.Fields{
	//		"DockerID": &graphql.Field{
	//			Type:        graphql.String,
	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//				return "MockID", nil
	//			},
	//		},
	//
	//	},
	//})

	//
	//queryType := graphql.NewObject(graphql.ObjectConfig{
	//	Name: "Query",
	//	Fields: graphql.Fields{
	//		"ECS_CONTAINER_METADATA": &graphql.Field{
	//			Type: graphql.String,
	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//				return "MockID", nil
	//			},
	//		},
	//	},
	//})
	//GQLSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
	//	Query: queryType,
	//})

	/*
	-Query Examples: getTasks(), getContainers(), etc.
	 */

	type Container struct {
		dockerID              string
		DockerName            string
		KnownStatus			  string
	}

	test := Container{
		dockerID: "ea32192c8553fbff06c9340478a2ff089b2bb5646fb718b4ee206641c9086d66",
		DockerName: "curl",
		KnownStatus: "RUNNING",
	}
	MockData := map[string]Container{
		"curl": test,
	}

	containerType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Container",
		Fields: graphql.Fields{
			"DockerId": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.DockerID, nil
					}
					return nil, errors.Errorf("Unable to retrieve DockerID for container")
				},
			},
			"Name": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.Name, nil
					}
					return nil, nil
				},
			},
			"DockerName": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.DockerName, nil
					}
					return nil, nil
				},
			},
			"Image": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.Image, nil
					}
					return nil, nil
				},
			},
			"ImageID": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.ImageID, nil
					}
					return nil, nil
				},
			},
			//TODO: Fix Labels Resolver
			"Labels": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetLabels(), nil
					}
					return nil, nil
				},
			},
			"DesiredStatus": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetDesiredStatus().String(), nil
					}
					return nil, nil
				},
			},
			"KnownStatus": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetKnownStatus().String(), nil
					}
					return nil, nil
				},
			},
			"ExitCode": &graphql.Field{
				Type:		 graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetKnownExitCode(), nil
					}
					return nil, nil
				},
			},
			"CreatedAt": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetCreatedAt().UTC().String(), nil
					}
					return nil, nil
				},
			},
			"StartedAt": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetStartedAt().UTC().String(), nil
					}
					return nil, nil
				},
			},
			"FinishedAt": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetFinishedAt().UTC().String(), nil
					}
					return nil, nil
				},
			},
			"Type": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.Type.String(), nil
					}
					return nil, nil
				},
			},
			"LogDriver": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetLogDriver(), nil
					}
					return nil, nil
				},
			},
			//TODO: Fix LogOptions
			"LogOptions": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.GetLogOptions(), nil
					}
					return nil, nil
				},
			},
			"ContainerARN": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if container, ok := p.Source.(*apicontainer.DockerContainer); ok {
						return container.Container.ContainerArn, nil
					}
					return nil, nil
				},
			},
		},
	})

	taskType := graphql.NewObject(graphql.ObjectConfig{
		Name: 		"Task",
		Fields: graphql.Fields{
			"Cluster": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return cluster, nil
				},
			},
			"TaskARN": &graphql.Field{
				Type:		 graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*apitask.Task); ok {
						return task.Arn, nil
					}
					return nil, nil
				},
			},
		},
	})

	fields := graphql.Fields{
		"Container": &graphql.Field{
			Type: containerType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				v3EndpointID := p.Context.Value("V3EndpointID").(string)
				dockerID, ok := state.DockerIDByV3EndpointID(v3EndpointID)
				if !ok {
					seelog.Errorf("unable to get docker ID from v3 endpoint ID: %s", v3EndpointID)
					return nil, errors.Errorf("unable to get docker ID from v3 endpoint ID: %s", v3EndpointID)
				}

				dockerContainer, ok := state.ContainerByID(dockerID)
				if !ok {
					seelog.Errorf("v2 container response: unable to find container '%s'", dockerID)
					return nil, errors.Errorf("v2 container response: unable to find container '%s'", dockerID)
				}
				return dockerContainer, nil
			},
		},
		"Task": &graphql.Field{
			Type: taskType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v3EndpointID := p.Context.Value("V3EndpointID").(string)
				taskARN, ok := state.TaskARNByV3EndpointID(v3EndpointID)
				if !ok {
					return "", errors.Errorf("unable to get task Arn from v3 endpoint ID: %s", v3EndpointID)
				}
				task, ok := state.TaskByArn(taskARN)
				return task, nil
			},
		},
		"containerByName" : &graphql.Field{
			Type: containerType,
			Args: graphql.FieldConfigArgument{
				"DockerName" : &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["DockerName"].(string)
				return MockData[name], nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	GQLSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return GQLSchema
}