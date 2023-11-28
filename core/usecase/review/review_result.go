package review

import "github.com/KarnerTh/qlookout/core/observer"

type ReviewResultPublisher = observer.Publisher[ReviewResult]
type ReviewResultSubscriber = observer.Subscriber[ReviewResult]

type ReviewResult struct {
	LookoutId int
	Rule      ReviewRule
	Result    ValidationResult
	Error     error
}
