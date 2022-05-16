package dbmodel

import (
	"testing"
)

func TestHttpReqFuncs(t *testing.T) {
	t.Run("GetSplitsDataHttp", func(t *testing.T) {
		_, err := getSplitsDataHttp("AAPL.US")
		if err != nil {
			t.Error(err)
		}
		// for _, s := range rec {
		// 	fmt.Printf("%+v", s)
		// }
	})

	t.Run("GetFundamentalDataHttp", func(t *testing.T) {
		_, err := getFundamentalDataHttp("AAPL.US")
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%+v", rec)
	})
}

func TestCachedDbRequests(t *testing.T) {
	dbClient := new(DbModel)
	err := dbClient.Connect()
	if err != nil {
		t.Errorf("Failed to Connect to Database: %+v", err)
	}

	t.Run("ReqCachedFundamentals", func(t *testing.T) {
		_, err := dbClient.ReqCachedFundamentals("AAPL.US")
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%+v", rec)
	})

	t.Run("ReqCachedSplitData", func(t *testing.T) {
		_, err := dbClient.ReqCachedSplitData("AAPL.US")
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%+v", rec)
	})

	dbClient.Disconnect()
	return
}
