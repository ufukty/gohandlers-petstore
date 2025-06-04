package pets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serialize(w http.ResponseWriter, errs map[string]error) error {
	s := map[string]string{}
	for param, err := range errs {
		s[param] = err.Error()
	}
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return fmt.Errorf("encoding json: %w", err)
	}
	return nil
}
