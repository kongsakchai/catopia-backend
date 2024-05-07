package domain

type ModelUsecae interface {
	CatGroup(input []float64) (int64, error)
}
