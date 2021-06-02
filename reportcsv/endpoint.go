package reportcsv

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetData endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetData: makeGetDataEndpoint(s),
	}
}

func makeGetDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDataRequest)
		data, err := s.GetData(ctx, req.Id)
	}

	return GetDataResponse{
		Data: data,
	}, err
}
