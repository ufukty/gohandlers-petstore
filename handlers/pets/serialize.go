package pets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serialize(w http.ResponseWriter, issues map[string]any) error {
	err := json.NewEncoder(w).Encode(issues)
	if err != nil {
		return fmt.Errorf("encoding json: %w", err)
	}
	return nil
}
