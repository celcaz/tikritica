package handler

import (
	"errors"
	"net/http"

	bookuc "github.com/celio/tikritica-api/internal/usecase/book"
	"github.com/celio/tikritica-api/pkg/response"
)

// ListBooks godoc
// @Summary      List all books
// @Description  Get a list of all books
// @Tags         books
// @Produce      json
// @Success      200  {array}   entity.Book
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/books [get]
func (h *Handler) listBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.BookUC.ListBooks(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to list books")
		return
	}

	response.JSON(w, http.StatusOK, books)
}

// GetBook godoc
// @Summary      Get book by slug
// @Description  Get details of a specific book by its slug
// @Tags         books
// @Produce      json
// @Param        slug   path      string  true  "Book Slug"
// @Success      200    {object}  entity.Book
// @Failure      404    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/books/{slug} [get]
func (h *Handler) getBook(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	book, err := h.BookUC.GetBookBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, bookuc.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "book not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to get book")
		return
	}

	response.JSON(w, http.StatusOK, book)
}
