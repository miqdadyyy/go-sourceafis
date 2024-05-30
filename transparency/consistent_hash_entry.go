package transparency

import "github.com/miqdadyyy/go-sourceafis/features"

type ConsistentHashEntry struct {
	Key   int
	Edges []*features.IndexedEdge
}
