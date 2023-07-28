package domain

type IMathService interface {
	ProcessExpression(expression string) (int, error)
}
