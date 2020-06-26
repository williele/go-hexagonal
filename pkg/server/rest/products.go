package rest

import (
	. "demo/pkg/services/products"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type productsEndpoint struct {
	productsService Service
}

func newProductsEndpoint(productsService Service) productsEndpoint {
	return productsEndpoint{
		productsService: productsService,
	}
}

func (e *productsEndpoint) routers() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", e.getAll())
	r.Get("/{id:[0-9]+}", e.getByID())
	r.Get("/{slug}", e.getBySlug())

	return r
}

// routers
// get all
func (e *productsEndpoint) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := &[]Product{}
		e.productsService.GetAll(products)

		response(http.StatusOK, w, r, products)
	}
}

// get by id
func (e *productsEndpoint) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responseError(http.StatusBadRequest, w, r, err)
		}

		product := &Product{}
		if err := e.productsService.GetByID(product, int64(id)); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
		}

		response(http.StatusOK, w, r, product)
	}
}

// get by slug
func (e *productsEndpoint) getBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get slug
		slug := chi.URLParam(r, "slug")

		product := &Product{}
		if err := e.productsService.GetBySlug(product, slug); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
		}

		response(http.StatusOK, w, r, product)
	}
}
