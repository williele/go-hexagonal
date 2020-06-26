package categories

// get all categories
func (s *service) GetAll(categories *[]Category) error {
	return s.repo.GetAll(categories)
}

// get category by id
func (s *service) GetByID(category *Category, id int64) error {
	return s.repo.GetByID(category, id)
}

// get category by slug name
func (s *service) GetBySlug(category *Category, slug string) error {
	return s.repo.GetBySlug(category, slug)
}
