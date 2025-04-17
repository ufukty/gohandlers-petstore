package pets

import (
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
		//
	}

	if err := bq.Validate(); err != nil {
		//
	}

	// implementation here
}
