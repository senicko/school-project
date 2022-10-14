package word_set

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

type WordSet struct {
	ID    int      `json:"id"`
	Title string   `json:"tite"`
	Words []string `json:"words"`
}

// Scan query into WorsSet struct
func Scan(r pgx.Row) (*WordSet, error) {
	ws := &WordSet{}

	if err := r.Scan(&ws.ID, &ws.Title, &ws.Words); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return ws, nil
}
