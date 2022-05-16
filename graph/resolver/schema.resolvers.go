package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-cape/graph/generated"
	"gqlgen-cape/graph/model"
)

func (r *queryResolver) GetFundamentals(ctx context.Context, ticker *string) (*model.EquityFundamentals, error) {
	// preRecord, err := r.Resolver.db.ReqCachedFundamentals(*ticker)
	// if err != nil {
	// 	return nil, err
	// }
	// record := model.EquityFundamentals{General: preRecord.Record.General}

	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCompanyOfficers(ctx context.Context, ticker *string) ([]*model.OfficerMap, error) {
	preRecord, err := r.Resolver.db.ReqCachedFundamentals(*ticker)
	if err != nil {
		return nil, err
	}

	officersRec := make([]model.OfficerMap, 0, len(preRecord.Record.General.Officers))
	officersRecPtrs := make([]*model.OfficerMap, len(preRecord.Record.General.Officers))
	for key, officer := range preRecord.Record.General.Officers {
		tmpString := key
		tmpOfficer := model.Officer{Name: officer.Name, YearBorn: officer.YearBorn, Title: officer.Title}
		officersRec = append(officersRec, model.OfficerMap{ItemNumber: &tmpString, Value: &tmpOfficer})
	}

	for i, _ := range officersRec {
		officersRecPtrs[i] = &officersRec[i]
	}
	return officersRecPtrs, nil
}

func (r *queryResolver) GetSplits(ctx context.Context, ticker *string) (*model.EquitySplits, error) {
	preRecord, err := r.Resolver.db.ReqCachedSplitData(*ticker)
	if err != nil {
		return nil, err
	}

	const layoutISO = "2006-01-02"
	time_formatted := preRecord.InsertDate.Format(layoutISO)

	splitsRec := make([]model.SplitsRecord, 0, len(preRecord.Record))
	splitsRecPtrs := make([]*model.SplitsRecord, len(preRecord.Record))
	for i, _ := range preRecord.Record {
		modrec := model.SplitsRecord(preRecord.Record[i])
		splitsRec = append(splitsRec, modrec)
		splitsRecPtrs[i] = &splitsRec[i]
	}

	record := model.EquitySplits{Ticker: preRecord.Ticker, InsertDate: time_formatted, Record: splitsRecPtrs}
	return &record, nil
}

func (r *queryResolver) GetCompanyTechnicals(ctx context.Context, ticker *string) (*model.Technicals, error) {
	preRecord, err := r.Resolver.db.ReqCachedFundamentals(*ticker)
	if err != nil {
		return nil, err
	}

	techrec := model.Technicals(preRecord.Record.Technicals)

	return &techrec, nil
}

func (r *queryResolver) GetCompanyHighlights(ctx context.Context, ticker *string) (*model.Highlights, error) {
	preRecord, err := r.Resolver.db.ReqCachedFundamentals(*ticker)
	if err != nil {
		return nil, err
	}

	hlightsrec := model.Highlights(preRecord.Record.Highlights)

	return &hlightsrec, nil
}

func (r *queryResolver) GetCompanyEarningsHistory(ctx context.Context, ticker *string) ([]*model.HistoryMapTuple, error) {
	preRecord, err := r.Resolver.db.ReqCachedFundamentals(*ticker)
	if err != nil {
		return nil, err
	}

	eHistoryRec := make([]model.HistoryMapTuple, 0, len(preRecord.Record.Earnings.History))
	eHistoryRecPtrs := make([]*model.HistoryMapTuple, len(preRecord.Record.Earnings.History))
	for key, val := range preRecord.Record.Earnings.History {
		modelrec := model.History(val)
		modrec := model.HistoryMapTuple{Key: key, Value: &modelrec}
		eHistoryRec = append(eHistoryRec, modrec)
	}
	for i, _ := range eHistoryRec {
		eHistoryRecPtrs[i] = &eHistoryRec[i]
	}
	return eHistoryRecPtrs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
