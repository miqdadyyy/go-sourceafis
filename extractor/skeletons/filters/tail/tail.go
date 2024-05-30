package tail

import (
	"github.com/miqdadyyy/go-sourceafis/config"
	"github.com/miqdadyyy/go-sourceafis/extractor/logger"
	"github.com/miqdadyyy/go-sourceafis/extractor/skeletons/filters/dot"
	"github.com/miqdadyyy/go-sourceafis/extractor/skeletons/filters/knot"
	"github.com/miqdadyyy/go-sourceafis/features"
)

type SkeletonTailFilter struct {
	logger logger.TransparencyLogger
}

func New(logger logger.TransparencyLogger) *SkeletonTailFilter {
	return &SkeletonTailFilter{
		logger: logger,
	}
}

func (f *SkeletonTailFilter) Apply(skeleton *features.Skeleton) error {
	for _, minutia := range skeleton.Minutiae {
		if len(minutia.Ridges) == 1 && len(minutia.Ridges[0].End().Ridges) >= 3 {
			if minutia.Ridges[0].Points.Size() < config.Config.MinTailLength {
				minutia.Ridges[0].Detach()
			}
		}
	}
	if err := dot.Apply(skeleton); err != nil {
		return err
	}

	if err := knot.Apply(skeleton); err != nil {
		return err
	}

	return f.logger.LogSkeleton("removed-tails", skeleton)
}
