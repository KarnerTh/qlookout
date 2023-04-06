package graphql

type LookoutResolver struct{}

func (_ LookoutResolver) Lookout() string { return "lookout works" }
