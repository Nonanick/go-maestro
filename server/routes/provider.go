package routes

type Provider interface {
	Routes() []Resolver
}
