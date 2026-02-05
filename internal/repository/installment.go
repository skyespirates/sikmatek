package repository

import "context"

type InstallmentRepository interface {
	CreateN(context.Context, QueryExecutor, string, int) error
}
