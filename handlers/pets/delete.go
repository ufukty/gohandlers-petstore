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
	// implementation here

	_ = &DeletePetRequest{}
}
