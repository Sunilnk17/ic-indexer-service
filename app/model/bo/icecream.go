package bo

import "time"

type ESIcecream struct {
	Id                    *int64     `json:"id"`
	Name                  *string    `json:"name,omitempty"`
	ImageClosed           *string    `json:"image_closed,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	LastUpdatedAt         *time.Time `json:"last_updated_at,omitempty"`
	ImageOpen             *string    `json:"image_open,omitempty"`
	Description           *string    `json:"description,omitempty"`
	Story                 *string    `json:"story,omitempty"`
	AllergyInfo           *string    `json:"allergy_info,omitempty"`
	SourcingValues        *[]string  `json:"sourcing_values,omitempty"`
	Ingredients           *[]string  `json:"ingredients,omitempty"`
	DietaryCertifications *string    `json:"dietary_certifications,omitempty"`
	ProductId             *string    `json:"product_id,omitempty"`
}

func NewESIcecream() *ESIcecream {
	return &ESIcecream{}
}

func (e *ESIcecream) GetCombinationKey() string {
	return *e.Name + "-" + *e.ProductId
}

func (e *ESIcecream) GetKey() string {
	return *e.ProductId
}
