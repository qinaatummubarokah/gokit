package reportcsv

import "context"

// Service describes the Account service.
type Service interface {
	GetData(ctx context.Context, id string) ([]Csvdata, error)
}
