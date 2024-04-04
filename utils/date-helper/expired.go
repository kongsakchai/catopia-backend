package dhelper

import (
	"time"

	"github.com/kongsakchai/catopia-backend/domain/date"
)

func CreateExpiredAtDay(day int) date.JSONDate {
	now := time.Now()
	expiredAt := now.AddDate(0, 0, day)
	return date.JSONDate(expiredAt)
}
