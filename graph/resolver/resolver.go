//go:generate go run github.com/99designs/gqlgen generate

package resolver

//import (
//	"gqlgen-cape/graph/model"
//)

import (
	"gqlgen-cape/graph/dbmodel"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *dbmodel.DbModel
	//	equityFundamentals *model.EquityFundamentals
}

func NewResolver() (*Resolver, error) {
	db := new(dbmodel.DbModel)
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return &Resolver{db: db}, nil
}

func (r *Resolver) Close() error {
	return r.db.Disconnect()
}
