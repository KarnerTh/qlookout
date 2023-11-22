package orchestration

import (
	"log/slog"
	"net/http"

	"github.com/KarnerTh/query-lookout/core/delivery/graphql"
	"github.com/KarnerTh/query-lookout/core/usecase/lookout"
	lookoutGraphQl "github.com/KarnerTh/query-lookout/core/usecase/lookout/delivery/graphql"
	"github.com/KarnerTh/query-lookout/core/usecase/notify"
	notifyGraphQl "github.com/KarnerTh/query-lookout/core/usecase/notify/delivery/graphql"
	"github.com/KarnerTh/query-lookout/core/usecase/review"
	reviewGraphQl "github.com/KarnerTh/query-lookout/core/usecase/review/delivery/graphql"
)

func setupDelivery(lookoutManager lookout.LookoutManager, lookoutRepo lookout.LookoutRepo, reviewRepo review.ReviewRepo, notificationSubscriber notify.NotificationSubscriber) {
	slog.Info("Start delivery")

	reviewResolver := reviewGraphQl.NewReviewResolver(reviewRepo)
	lookoutResolver := lookoutGraphQl.NewLookoutResolver(lookoutManager, lookoutRepo, reviewResolver)
	notifyResolver := notifyGraphQl.NewNotifyResolver(notificationSubscriber)

	graphql.Setup(
		"/query",
		&graphql.CombinedResolver{
			LookoutResolver: lookoutResolver,
			ReviewResolver:  reviewResolver,
			NotifyResolver:  notifyResolver,
		},
	)

	go slog.Error("Delivery failed", http.ListenAndServe(":63001", nil))
}
