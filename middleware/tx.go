package middleware

import (
	"context"
	"fmt"

	"github.com/marutaku/epub-index-creator/ent"
	goa "goa.design/goa/v3/pkg"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := client.Tx(ctx)
	ctx = ent.NewTxContext(ctx, tx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	v, err := fn(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return v, err
	}
	if err := tx.Commit(); err != nil {
		return v, fmt.Errorf("committing transaction: %w", err)
	}
	return v, nil
}

func NewTransaction(ctx context.Context, client *ent.Client) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		// A Goa endpoint is itself a function.
		return goa.Endpoint(func(ctx context.Context, req interface{}) (interface{}, error) {
			return WithTx(ctx, client, func(ctx context.Context) (interface{}, error) {
				// Call the original endpoint.
				return e(ctx, req)
			})
		})
	}
}
