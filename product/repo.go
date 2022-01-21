package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/log"
)

var RepoErr = errors.New("Unable to handle Repo request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{db: db, logger: log.With(logger, "repo", "sql")}
}

func (repo *repo) GetProduct(ctx context.Context, id int32) (Product, error) {
	logger := log.With(repo.logger, "method", "Repo GetProduct", ", id: ", id)
	var product Product
	err := repo.db.QueryRow("SELECT * FROM product WHERE id=$1", id).Scan(&product.id, &product.brand_name, &product.description, &product.internalName, &product.primaryCategoryId, &product.product_date, &product.productLabel, &product.productName, &product.productPrice, &product.productType, &product.variantProduct, &product.virtualProduct, &product.virtualVariantMethod)
	if err != nil {
		return product, RepoErr
	}
	logger.Log("Repo Get Product", product.brand_name)
	return product, nil
}
