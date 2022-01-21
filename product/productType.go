package product

type ProductType string

const (
	SIMPLE       ProductType = "Simple"
	CONFIGURABLE ProductType = "Configurable"
	GROUPED      ProductType = "Grouped"
	VIRTUAL      ProductType = "Virtual"
	BUNDLE       ProductType = "Bundle"
)
