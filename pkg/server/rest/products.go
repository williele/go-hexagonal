package rest

import (
	. "demo/pkg/services/products"
	"encoding/json"
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
	r.Post("/", e.create())
	r.Patch("/{id:[0-9]+}", e.update())

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
			return
		}

		product := &Product{}
		if err := e.productsService.GetByID(product, int64(id)); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, product)
	}
}

// get by slug
func (e *productsEndpoint) getBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get slug
		slug := chi.URLParam(r, "slug")

		// get product
		product := &Product{}
		if err := e.productsService.GetBySlug(product, slug); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, product)
	}
}

// create product
func (e *productsEndpoint) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse body
		input := &ProductCreateInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// create product
		product := &Product{}
		if err := e.productsService.Create(product, input); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusCreated, w, r, product)
	}
}

// update product
func (e *productsEndpoint) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// parse body
		input := &ProductUpdateInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// get product
		product := &Product{}
		if err := e.productsService.GetByID(product, int64(id)); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		// update product
		if err := e.productsService.Update(product, input); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, product)
	}
}
