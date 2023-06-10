package graphql

import (
	_ "embed"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/core/delivery"
	lookoutResolver "github.com/KarnerTh/query-lookout/core/usecase/lookout/delivery/graphql"
	notifyResolver "github.com/KarnerTh/query-lookout/core/usecase/notify/delivery/graphql"
	reviewResolver "github.com/KarnerTh/query-lookout/core/usecase/review/delivery/graphql"
)

//go:embed schema.graphql
var schemaContent string

type CombinedResolver struct {
	lookoutResolver.LookoutResolver
	reviewResolver.ReviewResolver
	notifyResolver.NotifyResolver
}

func Setup(endpoint string, resolver *CombinedResolver) {
	log.Info("Setup graphql schema and handler")
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schemaContent, resolver, opts...)
	graphQlHandler := graphqlws.NewHandlerFunc(schema, delivery.CorsMiddleware(&relay.Handler{Schema: schema}))
	http.HandleFunc(endpoint, graphQlHandler)
	setupGraphqlIde(endpoint)
}

func setupGraphqlIde(endpoint string) {
	// First argument must be same as graphql handler path
	graphiqlHandler, err := graphiql.NewGraphiqlHandler(endpoint)
	if err != nil {
		panic(err)
	}
	http.Handle("/query-ide", graphiqlHandler)
}
