package GQL

import (
	"encoding/json"
	"github.com/aws/amazon-ecs-agent/agent/engine/dockerstate"
	"net/http"

	"github.com/graphql-go/graphql"
)







func ContainerMetadataHandler(state dockerstate.TaskEngineState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//containerID, err := v3.GetContainerIDByRequest(r, state)
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        GQLSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	}
}