package types

import (
	"time"
)

type RetrieveItem struct {
	Word     string      `json:"word,omitempty"`
	EN       string      `json:"en,omitempty"`
	CN       string      `json:"cn,omitempty"`
	Examples []string    `json:"examples,omitempty"`
	Hits     []time.Time `json:"hits,omitempty"`
	Synomyms []string    `json:"synomyms,omitempty"`
}
