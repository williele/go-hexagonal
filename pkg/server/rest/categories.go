package rest

import (
	. "demo/pkg/services/categories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type categoriesEndpoint struct {
	categoriesService Service
}

func newCategoriesEndpoint(categoriesService Service) categoriesEndpoint {
	return categoriesEndpoint{
		categoriesService: categoriesService,
	}
}

func (e *categoriesEndpoint) routers() *chi.Mux {
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
func (e *categoriesEndpoint) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories := &[]Category{}
		e.categoriesService.GetAll(categories)

		response(http.StatusOK, w, r, categories)
	}
}

// get by id
func (e *categoriesEndpoint) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		category := &Category{}
		if err := e.categoriesService.GetByID(category, int64(id)); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, category)
	}
}

// get by slug
func (e *categoriesEndpoint) getBySlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get slug
		slug := chi.URLParam(r, "slug")

		// get category
		category := &Category{}
		if err := e.categoriesService.GetBySlug(category, slug); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, category)
	}
}

// create category
func (e *categoriesEndpoint) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse body
		input := &CategoryCreateInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// create category
		category := &Category{}
		if err := e.categoriesService.Create(category, input); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusCreated, w, r, category)
	}
}

// update category
func (e *categoriesEndpoint) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// parse body
		input := &CategoryUpdateInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			responseError(http.StatusBadRequest, w, r, err)
			return
		}

		// get category
		category := &Category{}
		if err := e.categoriesService.GetByID(category, int64(id)); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		// update category
		if err := e.categoriesService.Update(category, input); err != nil {
			responseError(http.StatusInternalServerError, w, r, err)
			return
		}

		response(http.StatusOK, w, r, category)
	}
}
