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

	bq := &GetPetRequest{}

	if err := bq.Parse(r); err != nil {
		//
	}

	// implementation here

	bs := &GetPetResponse{}
	if err := bs.Write(w); err != nil {
		//
	}
}
