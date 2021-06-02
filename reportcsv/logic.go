package reportcsv

import (
	"context"
	"log"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) GetData(ctx context.Context, id string) ([]Csvdata, error) {
	logger := log.With(s.logger, "method", "GetData")

	data, err := s.repostory.GetData(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return data, err
	}

	logger.Log("Get data")

	return data, nil
}
