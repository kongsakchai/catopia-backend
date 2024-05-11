package domain

import "context"

type RecommendUsecase interface {
	RecommendByCat(ctx context.Context, id int64, userID int64) ([]string, error)
	RecommendByUser(ctx context.Context, userID int64) ([]string, error)
}
