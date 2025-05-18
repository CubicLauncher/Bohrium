package handlers

import (
	"encoding/json"
	"net/http"

	"Bohrium/models"
	"Bohrium/services"
)

type SkinHandler struct {
	minecraftService *services.MinecraftService
}

func NewSkinHandler(minecraftService *services.MinecraftService) *SkinHandler {
	return &SkinHandler{
		minecraftService: minecraftService,
	}
}

func (h *SkinHandler) HandleSkinRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	parts := splitPath(r.URL.Path)
	if len(parts) != 2 || parts[0] != "skin" || parts[1] == "" {
		http.Error(w, `{"error":"Username requerido en la ruta /skin/{username}"}`, http.StatusBadRequest)
		return
	}
	username := parts[1]

	uuid, err := h.minecraftService.GetUUID(username)
	if err != nil {
		http.Error(w, `{"error":"No se encontró el jugador"}`, http.StatusBadRequest)
		return
	}

	headOnly := r.URL.Query().Get("head") == "true"

	var skinURL string
	if headOnly {
		skinURL = h.minecraftService.GetHeadURL(uuid)
	} else {
		skinURL, err = h.minecraftService.GetSkinURL(uuid)
		if err != nil {
			http.Error(w, `{"error":"No se pudo obtener la skin"}`, http.StatusInternalServerError)
			return
		}
	}

	resp := models.SkinURLResponse{
		Username: username,
		UUID:     uuid,
		SkinURL:  skinURL,
	}
	json.NewEncoder(w).Encode(resp)
}

func splitPath(path string) []string {
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	return split(path, '/')
}

func split(s string, sep rune) []string {
	var res []string
	curr := ""
	for _, c := range s {
		if c == sep {
			res = append(res, curr)
			curr = ""
		} else {
			curr += string(c)
		}
	}
	res = append(res, curr)
	return res
} 
