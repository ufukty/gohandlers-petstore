package pets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serialize(w http.ResponseWriter, errs map[string]any) error {
	err := json.NewEncoder(w).Encode(errs)
	if err != nil {
		return fmt.Errorf("encoding json: %w", err)
	}
	return nil
}
