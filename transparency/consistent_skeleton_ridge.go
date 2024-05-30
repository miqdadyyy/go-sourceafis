package transparency

import "github.com/miqdadyyy/go-sourceafis/primitives"

type ConsistentSkeletonRidge struct {
	Start, End int
	Points     primitives.List[primitives.IntPoint]
}
