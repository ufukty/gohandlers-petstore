package pets

import (
	"log/slog"
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/dto"
	"github.com/ufukty/gohandlers-petstore/handlers/pets/types"
)

type GetRequest struct {
	ID types.PetId `route:"id"`
}

type GetResponse struct {
	Pet dto.Pet `json:"pet"`
}

// GET /pets/{id}
func (p *Pets) Get(w http.ResponseWriter, r *http.Request) {
	bq := &GetRequest{}

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

	bs := &GetResponse{
		Pet: dto.Pet{},
	}
	if err := bs.Write(w); err != nil {
		slog.Debug("user error on serializing response", "content", err.Error())
		http.Error(w, "serializing response", 400)
		return
	}
}
