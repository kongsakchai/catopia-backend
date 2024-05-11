package usecase

import (
	"context"

	"github.com/kongsakchai/catopia-backend/domain"
)

type recommendUsecase struct {
	catUsecase  domain.CatUsecase
	userUsecase domain.UserUsecase
}

func NewRecommendUsecase(catUsecase domain.CatUsecase, userUsecase domain.UserUsecase) domain.RecommendUsecase {
	return &recommendUsecase{catUsecase, userUsecase}
}

func (u *recommendUsecase) RecommendByCat(ctx context.Context, id int64, userID int64) ([]string, error) {
	return u.catUsecase.GetBreedingByCat(ctx, id, userID)
}

func (u *recommendUsecase) RecommendByUser(ctx context.Context, userID int64) ([]string, error) {
	ids, err := u.userUsecase.GetUserIDsInSameGroup(ctx, userID)
	if err != nil || len(ids) == 0 {
		return u.catUsecase.GetBreedingByRandom(ctx)
	}

	return u.catUsecase.GetBreedingByUser(ctx, ids)
}
