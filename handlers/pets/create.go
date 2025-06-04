package pets

import (
	"log/slog"
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/types"
)

type CreatePetRequest struct {
	Name types.PetName `json:"name"`
	Tag  types.PetTag  `json:"tag"`
}

type CreatePetResponse struct {
	ID string `json:"id"`
}

func (p *Pets) CreatePet(w http.ResponseWriter, r *http.Request) {
	bq := &CreatePetRequest{}

	if err := bq.Parse(r); err != nil {
		slog.Debug("error on parsing request", "content", err.Error())
		http.Error(w, "error on parsing request", 400)
		return
	}

	if errs := bq.Validate(); errs != nil {
		if err := serialize(w, errs); err != nil {
			slog.Error("error on serialization validation error", "content", err)
			http.Error(w, "error on serialization validation error", 500)
		} else {
			slog.Debug("error on validation", "length", len(errs))
		}
		return
	}

	// ...

	bs := &CreatePetResponse{
		ID: "",
	}
	if err := bs.Write(w); err != nil {
		slog.Debug("error on serializing response", "content", err.Error())
		http.Error(w, "serializing response", 400)
		return
	}
}
