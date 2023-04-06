package orchestration

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/delivery/graphql"
)

func setupDelivery() {
	log.Info("Start delivery")
	graphql.Setup("/query")

	// TODO: port config?
	go log.Fatal(http.ListenAndServe(":8080", nil))
}
