package pets

import (
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/dto"
	"github.com/ufukty/gohandlers/pkg/types/basics"
)

type GetPetRequest struct {
	ID basics.String `route:"id"`
}

type GetPetResponse struct {
	Pet dto.Pet `json:"pet"`
}

// GET /pets/{id}
func (p *Pets) GetPet(w http.ResponseWriter, r *http.Request) {
	// implementation here

	_ = &GetPetRequest{}
}
