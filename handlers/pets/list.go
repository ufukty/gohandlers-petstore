package pets

import (
	"net/http"

	"github.com/ufukty/gohandlers-petstore/handlers/pets/dto"
	"github.com/ufukty/gohandlers-petstore/handlers/pets/types"
)

type ListPetsRequest struct {
	Limit types.ListLimit `query:"limit"` // optional
}

type ListPetsResponse struct {
	Pets []dto.Pet `json:"pets"`
}

// GET /pets
func (p *Pets) ListPets(w http.ResponseWriter, r *http.Request) {
	bq := &ListPetsRequest{}

	if err := bq.Parse(r); err != nil {
		//
	}

	if err := bq.Validate(); err != nil {
		//
	}

	// implementation here

	bs := &ListPetsResponse{}
	if err := bs.Write(w); err != nil {
		//
	}
}
