package product

import (
	"context"
	"database/sql"
)

// Product structure
type Product struct {
	id                   int32          `json:"id"`
	productType          sql.NullString `json:"productType"`
	productLabel         sql.NullString `json:"productLabel"`
	internalName         sql.NullString `json:"internalName"`
	brand_name           string         `json:"brand_name"`
	productName          sql.NullString `json:"productName"`
	description          sql.NullString `json:"description"`
	virtualProduct       sql.NullString `json:"virtualProduct"`
	variantProduct       sql.NullString `json:"variantProduct"`
	virtualVariantMethod sql.NullString `json:"virtualVariantMethod"`
	primaryCategoryId    sql.NullString `json:"primaryCategoryId"`
	product_date         sql.NullTime   `json:"product_date"`
	productPrice         sql.NullString `json:"productPrice"`
}

type Repository interface {
	GetProduct(ctx context.Context, id int32) (Product, error)
}
