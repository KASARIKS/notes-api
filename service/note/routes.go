package note

import (
	"net/http"
	"time"

	"github.com/kasariks/notes_api/service/auth"
	"github.com/kasariks/notes_api/types"
	"github.com/kasariks/notes_api/utils"
)

type Handler struct {
	store     types.NotesStore
	userStore types.UserStore
}

func NewHandler(store types.NotesStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(routes *http.ServeMux) {
	routes.HandleFunc("POST /create_note", auth.WithJWTAuth(h.HandleCreateNote, h.userStore))
	routes.HandleFunc("POST /get_notes", auth.WithJWTAuth(h.HandleGetNotes, h.userStore))
}

func (h *Handler) HandleCreateNote(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())
	var payload types.NotePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	newNote := &types.Note{
		UserID:    userID,
		Name:      payload.Name,
		Value:     payload.Value,
		CreatedAt: time.Now().String(),
	}

	err := h.store.CreateNote(userID, *newNote)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
	return
}

func (h *Handler) HandleGetNotes(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIdFromContext(r.Context())
	notes, err := h.store.GetNotesByUserId(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, notes)
}
