package orchestration

import "github.com/KarnerTh/query-lookout/usecase/lookout"

func setupLookout(lookoutRepo lookout.LookoutRepo) {
	l := lookout.NewLookoutService(lookoutRepo)
	l.Start()
}
