package review

import "github.com/KarnerTh/query-lookout/notifier"

type ReviewResultPublisher = notifier.Publisher[ReviewResult]
type ReviewResultSubscriber = notifier.Subscriber[ReviewResult]

type ReviewResult struct {
	Rule    ReviewRule
	Success bool
}
