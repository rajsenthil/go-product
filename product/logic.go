package product

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{repository: repository, logger: logger}
}

func (s service) GetProduct(ctx context.Context, id int32) (Product, error) {
	logger := log.With(s.logger, "method", "GetProduct")
	product, err := s.repository.GetProduct(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return product, err
	}

	logger.Log("logic Get Product: ", product.brand_name)
	return product, nil
}
