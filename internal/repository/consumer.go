package repository

import "context"

type ConsumerRepository interface {
	GetIdByUserId(context.Context, QueryExecutor, int) (int, error)
}
