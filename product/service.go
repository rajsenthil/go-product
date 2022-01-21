package product

import "context"

type Service interface {
	GetProduct(ctx context.Context, id int32) (Product, error)
}
