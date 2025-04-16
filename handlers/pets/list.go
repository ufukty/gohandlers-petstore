package pets

import (
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/dto"
	"github.com/ufukty/gohandlers/pkg/types/basics"
)

type ListPetsRequest struct {
	Limit basics.Int `query:"limit"` // optional
}

type ListPetsResponse struct {
	Pets []dto.Pet `json:"pets"`
}

// GET /pets
func (p *Pets) ListPets(w http.ResponseWriter, r *http.Request) {
	// implementation here

	_ = &ListPetsRequest{}
}
