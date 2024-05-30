package sourceafis

import (
	"github.com/miqdadyyy/go-sourceafis/extractor"
	"github.com/miqdadyyy/go-sourceafis/extractor/logger"
	"github.com/miqdadyyy/go-sourceafis/primitives"
	"github.com/miqdadyyy/go-sourceafis/templates"
)

type Extractor interface {
	Extract(raw *primitives.Matrix, dpi float64) (*templates.FeatureTemplate, error)
}

type TemplateCreator struct {
	logger    logger.TransparencyLogger
	extractor Extractor
}

func NewTemplateCreator(logger logger.TransparencyLogger) *TemplateCreator {
	return &TemplateCreator{
		logger:    logger,
		extractor: extractor.New(logger),
	}
}

func (c *TemplateCreator) Template(img *Image) (*templates.SearchTemplate, error) {
	ft, err := c.extractor.Extract(img.matrix, img.dpi)
	if err != nil {
		return nil, err
	}
	if c.logger != nil {
		err := c.logger.Log("shuffled-minutiae", ft)
		if err != nil {
			return nil, err
		}
	}

	return templates.NewSearchTemplate(c.logger, ft), nil
}
