package orchestration

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/delivery/graphql"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	lookoutGraphQl "github.com/KarnerTh/query-lookout/usecase/lookout/delivery/graphql"
)

func setupDelivery(lookoutManager lookout.LookoutManager, lookoutService lookout.LookoutService) {
	log.Info("Start delivery")

	graphql.Setup(
		"/query",
		&graphql.CombinedResolver{
			LookoutResolver: lookoutGraphQl.NewLookoutResolver(lookoutManager, lookoutService),
		},
	)

	// TODO: port config?
	go log.Fatal(http.ListenAndServe(":8080", nil))
}
