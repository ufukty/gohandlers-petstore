package pets

import (
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
		//
	}

	if err := bq.Validate(); err != nil {
		//
	}

	// implementation here

	bs := &CreatePetResponse{}
	if err := bs.Write(w); err != nil {
		//
	}
}
