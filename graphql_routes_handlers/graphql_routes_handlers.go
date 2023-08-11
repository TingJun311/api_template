package graphqlrouteshandlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/graphql-go/graphql"
)

func HandleGraphQLRequest(w http.ResponseWriter, r *http.Request) {
    var reqBody map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    query, _ := reqBody["query"].(string)
	newSchema, err := graphqlSchema()
	if err != nil {
		// handler error
	}

    params := graphql.Params{
        Schema:        newSchema,
        RequestString: query,
    }

    result := graphql.Do(params)
    if len(result.Errors) > 0 {
        http.Error(w, fmt.Sprintf("GraphQL query error: %v", result.Errors), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func graphqlSchema() (graphql.Schema, error) {
	// Configure GraphQL API
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: buildRootQuery(),
	})
	if err != nil {
		return schema, err
	}
	return schema, nil
}

func buildRootQuery() *graphql.Object {
    fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "Hello, GraphQL!", nil
            },
        },
    }

    return graphql.NewObject(graphql.ObjectConfig{
        Name:   "RootQuery",
        Fields: fields,
    })
}