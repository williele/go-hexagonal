package categories

import (
	. "demo/pkg/services"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

// create new category
func (p *service) Create(category *Category, input *CategoryCreateInput) error {
	// validate input
	if err := Validate.Struct(input); err != nil {
		return NewErrValidate("create category", err.(validator.ValidationErrors))
	}

	// name
	category.Name = strings.ToLower(strings.Trim(input.Name, " "))
	// generate slug
	category.Slug = slug.Make(category.Name)
	// check if slug already exists
	if ok, err := p.repo.CheckSlugExists(category.Slug); ok || err != nil {
		if err != nil {
			return err
		}
		return NewErrInputInvalid("create category", map[string]string{"name": "validate-duplicate"})
	}

	return nil
}

// update category
func (p *service) Update(category *Category, input *CategoryUpdateInput) error {
	// validate input
	if err := Validate.Struct(input); err != nil {
		return NewErrValidate("update category", err.(validator.ValidationErrors))
	}

	// title
	if input.Name != "" {
		category.Name = strings.ToLower(strings.Trim(input.Name, " "))
		// update slug name
		category.Slug = slug.Make(category.Name)

		// validate slug name duplicate
		if ok, err := p.repo.CheckSlugExists(category.Slug); ok || err != nil {
			if err != nil {
				return err
			}
			return NewErrInputInvalid("update product", map[string]string{"name": "validate-duplicate"})
		}
	}

	// update
	if err := p.repo.Update(category); err != nil {
		return err
	}

	return nil
}
