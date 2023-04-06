package graphql

type ReviewResolver struct{}

func (_ ReviewResolver) Rule() string { return "rule works" }
