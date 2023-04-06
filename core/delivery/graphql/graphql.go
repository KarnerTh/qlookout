package graphql

import (
	_ "embed"
	"net/http"

	lookoutResolver "github.com/KarnerTh/query-lookout/usecase/lookout/delivery/graphql"
	reviewResolver "github.com/KarnerTh/query-lookout/usecase/review/delivery/graphql"
	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"
)

//go:embed schema.graphql
var schemaContent string

type CombinedResolver struct {
	lookoutResolver.LookoutResolver
	reviewResolver.ReviewResolver
}

func Setup(endpoint string) {
	log.Info("Setup graphql schema and handler")
	schema := graphql.MustParseSchema(schemaContent, &CombinedResolver{})
	http.Handle(endpoint, &relay.Handler{Schema: schema})
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
