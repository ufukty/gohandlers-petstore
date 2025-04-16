package pets

import (
	"net/http"

	"github.com/ufukty/gohandlers/pkg/types/basics"
)

type CreatePetRequest struct {
	Name basics.String `json:"name"`
	Tag  basics.String `json:"tag"`
}

type CreatePetResponse struct {
	ID string `json:"id"`
}

func (p *Pets) CreatePet(w http.ResponseWriter, r *http.Request) {
	bq := &CreatePetRequest{}

	if err := bq.Parse(r); err != nil {
		//
	}

	// implementation here

	bs := &CreatePetResponse{}
	if err := bs.Write(w); err != nil {
		//
	}
}
