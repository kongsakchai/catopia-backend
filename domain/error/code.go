package errs

const (
	ErrNotFound     = 1
	ErrInternal     = 2
	ErrConflict     = 3
	ErrBadRequest   = 4
	ErrUnauthorized = 5
	ErrForbidden    = 6

	ErrTransaction = 7
	ErrRollback    = 8
	ErrCommit      = 9

	ErrUserGetByEmail     = 1000
	ErrUserGetByID        = 1001
	ErrUserCreate         = 1002
	ErrUserUpdate         = 1003
	ErrUserGetByUsername  = 1004
	ErrUserUpdatePassword = 1005
	ErrUserUpdateGroup    = 1006
	ErrUserGetByGroup     = 1007

	ErrTreatmentTypeGet    = 2000
	ErrTreatmentGetByID    = 2001
	ErrTreatmentGetByCatID = 2001
	ErrTreatmentCreate     = 2002
	ErrTreatmentUpdate     = 2003
	ErrTreatmentDelete     = 2004

	ErrSessionCreate = 3000
	ErrSessionGet    = 3001
	ErrSessionDelete = 3002

	ErrCatGetByID     = 4000
	ErrCatGetByUserID = 4001
	ErrCatCreate      = 4002
	ErrCatUpdate      = 4003
	ErrCatDelete      = 4004
	ErrCatUpdateGroup = 4005
	ErrCatGetCluster  = 4006
	ErrCatGetRandom   = 4007

	ErrCatGroup = 5000
)
