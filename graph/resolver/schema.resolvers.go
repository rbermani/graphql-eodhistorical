package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"gqlgen-cape/graph/generated"
	"gqlgen-cape/graph/model"
	"net/http"
	"time"
)

func (r *queryResolver) GetFundamentals(ctx context.Context, ticker *string) (*model.EquityFundamentals, error) {
	const eod_fundamental_api string = "https://eodhistoricaldata.com/api/fundamentals/AAPL.US?api_token="
	const eod_api_key string = "OeAFFmMliFG5orCUuwAKQ8l4WWFQ67YX"
	var http_client = &http.Client{Timeout: 10 * time.Second}
	result, err := http_client.Get(eod_fundamental_api + eod_api_key)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	fundamental_data := new(model.EquityFundamentals)

	err = json.NewDecoder(result.Body).Decode(&fundamental_data)
	if err != nil {
		return nil, err
	}
	return fundamental_data, nil
}

func (r *technicalsResolver) WeekHigh52(ctx context.Context, obj *model.Technicals) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *technicalsResolver) WeekLow52(ctx context.Context, obj *model.Technicals) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *technicalsResolver) DayMa50(ctx context.Context, obj *model.Technicals) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *technicalsResolver) DayMa200(ctx context.Context, obj *model.Technicals) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Technicals returns generated.TechnicalsResolver implementation.
func (r *Resolver) Technicals() generated.TechnicalsResolver { return &technicalsResolver{r} }

type queryResolver struct{ *Resolver }
type technicalsResolver struct{ *Resolver }
