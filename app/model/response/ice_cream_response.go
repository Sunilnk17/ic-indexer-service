package response

import "time"

type BulkIcecreamIndexResponse struct {
	Icecreams    []IcecreamIndexResponse
	TotalResults int64
}

type IcecreamIndexResponse struct {
	Id                    *int64     `json:"icecream_id"`
	Name                  *string    `json:"name,omitempty"`
	ImageClosed           *string    `json:"image_closed,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	LastUpdatedAt         *time.Time `json:"last_updated_at,omitempty"`
	ImageOpen             *string    `json:"image_open,omitempty"`
	Description           *string    `json:"description,omitempty"`
	Story                 *time.Time `json:"story,omitempty"`
	AllergyInfo           *string    `json:"allergy_info,omitempty"`
	SourcingValues        *[]string  `json:"sourcing_values,omitempty"`
	Ingredients           *[]string  `json:"ingredients,omitempty"`
	DietaryCertifications *string    `json:"dietary_certifications,omitempty"`
	ProductId             *string    `json:"product_id,omitempty"`
	//13
}
