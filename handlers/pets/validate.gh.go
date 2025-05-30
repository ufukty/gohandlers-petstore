// Code generated by gohandlers . DO NOT EDIT.

package pets

func (bq CreatePetRequest) Validate() (errs map[string]error) {
	errs = map[string]error{}
	if err := bq.Name.Validate(); err != nil {
		errs["name"] = err
	}
	if err := bq.Tag.Validate(); err != nil {
		errs["tag"] = err
	}
	return
}

func (bq DeletePetRequest) Validate() (errs map[string]error) {
	errs = map[string]error{}
	if err := bq.ID.Validate(); err != nil {
		errs["id"] = err
	}
	return
}

func (bq GetPetRequest) Validate() (errs map[string]error) {
	errs = map[string]error{}
	if err := bq.ID.Validate(); err != nil {
		errs["id"] = err
	}
	return
}

func (bq ListPetsRequest) Validate() (errs map[string]error) {
	errs = map[string]error{}
	if err := bq.Limit.Validate(); err != nil {
		errs["limit"] = err
	}
	return
}
