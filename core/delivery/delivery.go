package delivery

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/delivery/graphql"
)

func Start() {
	log.Info("Start delivery")
	graphql.Start("/query")

	// TODO: port config?
	log.Fatal(http.ListenAndServe(":8080", nil))
}
