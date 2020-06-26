package products

import (
	. "demo/pkg/services"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

// create new product
func (p *service) Create(product *Product, input *ProductCreateInput) error {
	// validate input
	if err := Validate.Struct(input); err != nil {
		return NewErrValidate("create product", err.(validator.ValidationErrors))
	}

	// title
	product.Title = strings.ToLower(strings.Trim(input.Title, " "))
	// generate slug
	product.Slug = slug.Make(product.Title)
	// check if slug already exists
	if ok, err := p.repo.CheckSlugExists(product.Slug); ok || err != nil {
		if err != nil {
			return err
		}
		return NewErrInputInvalid("create product", map[string]string{"title": "validate-duplicate"})
	}
	// description
	product.Description = strings.Trim(input.Description, " ")
	// price
	product.Price = input.Price
	// create
	if err := p.repo.Create(product); err != nil {
		return err
	}

	return nil
}

// update product
func (p *service) Update(product *Product, input *ProductUpdateInput) error {
	// validate input
	if err := Validate.Struct(input); err != nil {
		return NewErrValidate("update product", err.(validator.ValidationErrors))
	}

	// title
	if input.Title != "" {
		product.Title = strings.ToLower(strings.Trim(input.Title, " "))
		// update slug name
		product.Slug = slug.Make(product.Title)

		// validate slug name duplicate
		if ok, err := p.repo.CheckSlugExists(product.Slug); ok || err != nil {
			if err != nil {
				return err
			}
			return NewErrInputInvalid("update product", map[string]string{"title": "validate-duplicate"})
		}
	}
	// description
	if input.Description != "" {
		product.Description = strings.Trim(input.Description, " ")
	}
	// price
	if input.Price != nil {
		product.Price = *input.Price
	}
	// updated timestamp
	product.UpdatedAt = time.Now().UTC()

	// update
	if err := p.repo.Update(product); err != nil {
		return err
	}

	return nil
}
