package products

// get all products
func (s *service) GetAll(products *[]Product) error {
	return s.repo.GetAll(products)
}

// get product by id
func (s *service) GetByID(product *Product, id int64) error {
	return s.repo.GetByID(product, id)
}

// get product by slug name
func (s *service) GetBySlug(product *Product, slug string) error {
	return s.repo.GetBySlug(product, slug)
}
