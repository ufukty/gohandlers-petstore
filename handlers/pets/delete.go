package pets

import (
	"net/http"

	"github.com/ufukty/gohandlers/pkg/types/basics"
)

type DeletePetRequest struct {
	ID basics.String `route:"id"`
}

// DELETE /pets/{id}
func (p *Pets) DeletePet(w http.ResponseWriter, r *http.Request) {
	bq := &DeletePetRequest{}

	if err := bq.Parse(r); err != nil {
		//
	}

	// implementation here
}
