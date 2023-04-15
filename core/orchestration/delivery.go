package orchestration

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/delivery/graphql"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	lookoutGraphQl "github.com/KarnerTh/query-lookout/usecase/lookout/delivery/graphql"
	"github.com/KarnerTh/query-lookout/usecase/review"
	reviewGraphQl "github.com/KarnerTh/query-lookout/usecase/review/delivery/graphql"
)

func setupDelivery(lookoutManager lookout.LookoutManager, lookoutRepo lookout.LookoutRepo, reviewRepo review.ReviewRepo) {
	log.Info("Start delivery")

	reviewResolver := reviewGraphQl.NewReviewResolver(reviewRepo)
	lookoutResolver := lookoutGraphQl.NewLookoutResolver(lookoutManager, lookoutRepo, reviewResolver)

	graphql.Setup(
		"/query",
		&graphql.CombinedResolver{
			LookoutResolver: lookoutResolver,
			ReviewResolver:  reviewResolver,
		},
	)

	// TODO: port config?
	go log.Fatal(http.ListenAndServe(":8080", nil))
}
