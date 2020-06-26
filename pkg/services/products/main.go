package products

// database repository
type Repository interface {
	GetAll(*[]Product) error
	GetByID(*Product, int64) error
	GetBySlug(*Product, string) error

	CheckSlugExists(string) (bool, error)

	Create(*Product) error
	Update(*Product) error
	Delete(*Product) error
}

// service
type Service interface {
	GetAll(*[]Product) error
	GetByID(*Product, int64) error
	GetBySlug(*Product, string) error

	Create(*Product, *ProductCreateInput) error
	Update(*Product, *ProductUpdateInput) error
}

// service implemnet
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
