package pets

import (
	"log/slog"
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/types"
)

type DeletePetRequest struct {
	ID types.PetId `route:"id"`
}

// DELETE /pets/{id}
func (p *Pets) DeletePet(w http.ResponseWriter, r *http.Request) {
	bq := &DeletePetRequest{}

	if err := bq.Parse(r); err != nil {
		slog.Debug("user error on parsing request", "content", err.Error())
		http.Error(w, "error on parsing request", 400)
		return
	}

	if issues := bq.Validate(); issues != nil {
		if err := serialize(w, issues); err != nil {
			slog.Error("error on serialization validation error", "content", err)
			http.Error(w, "error on serialization validation error", 500)
		} else {
			slog.Debug("user error on validation", "length", len(issues))
		}
		return
	}

	// ...
}
