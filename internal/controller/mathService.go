package controller

type IMathService interface {
	ProcessExpression(expression string) (int, error)
}
