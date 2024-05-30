package matcher

import (
	"github.com/miqdadyyy/go-sourceafis/features"
	"github.com/miqdadyyy/go-sourceafis/templates"
)

type Probe struct {
	template *templates.SearchTemplate
	hash     map[int][]*features.IndexedEdge
}

func NewProbe(template *templates.SearchTemplate, hash map[int][]*features.IndexedEdge) *Probe {
	return &Probe{
		template: template,
		hash:     hash,
	}
}
