package routes

type Provider interface {
	Routes() []ExecutionPlan
}
