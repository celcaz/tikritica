package handler

import (
	"net/http"
)

// GetDiary godoc
// @Summary      Get user diary
// @Description  Get diary entries for the current user
// @Tags         diary
// @Produce      json
// @Success      200  {array}   entity.DiaryEntry
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/diary [get]
func (h *Handler) getDiary(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// CreateDiaryEntry godoc
// @Summary      Create diary entry
// @Description  Create a new diary entry
// @Tags         diary
// @Accept       json
// @Produce      json
// @Param        entry  body      entity.DiaryEntry  true  "Diary entry data"
// @Success      201    {object}  entity.DiaryEntry
// @Failure      400    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/diary [post]
func (h *Handler) createDiaryEntry(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
