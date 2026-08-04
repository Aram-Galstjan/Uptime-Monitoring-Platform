package postgres

import "context"

type Repository interface {
	Ping(context.Context) error
}
