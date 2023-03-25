package review

import "github.com/KarnerTh/query-lookout/observer"


type ReviewResultPublisher = observer.Publisher[ReviewResult]
type ReviewResultSubscriber = observer.Subscriber[ReviewResult]

type ReviewResult struct {
	Rule    ReviewRule
	Success bool
}
