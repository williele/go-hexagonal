package categories

// database repository
type Repository interface {
	GetAll(*[]Category) error
	GetByID(*Category, int64) error
	GetBySlug(*Category, string) error

	CheckSlugExists(string) (bool, error)

	Create(*Category) error
	Update(*Category) error
	Delete(*Category) error
}

// service
type Service interface {
	GetAll(*[]Category) error
	GetByID(*Category, int64) error
	GetBySlug(*Category, string) error

	Create(*Category, *CategoryCreateInput) error
	Update(*Category, *CategoryUpdateInput) error
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
