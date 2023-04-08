package graphql

import (
	_ "embed"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/delivery"
	lookoutResolver "github.com/KarnerTh/query-lookout/usecase/lookout/delivery/graphql"
	reviewResolver "github.com/KarnerTh/query-lookout/usecase/review/delivery/graphql"
)

//go:embed schema.graphql
var schemaContent string

type CombinedResolver struct {
	lookoutResolver.LookoutResolver
	reviewResolver.ReviewResolver
}

func Setup(endpoint string, resolver *CombinedResolver) {
	log.Info("Setup graphql schema and handler")
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schemaContent, resolver, opts...)
	http.Handle(endpoint, delivery.CorsMiddleware(&relay.Handler{Schema: schema}))
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
