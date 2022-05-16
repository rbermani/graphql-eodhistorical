package dbmodel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gqlgen-cape/graph/servertype"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI              string        = "mongodb://localhost:27017"
	MONGO_DB               string        = "cape"
	MONGO_FUNDA_COL        string        = "fundamentals"
	MONGO_SPLITS_COL       string        = "splits"
	QUERY_TIMEOUT_S        time.Duration = 5
	HTTP_REQ_TIMEOUT_S     time.Duration = 5
	EOD_API_KEY            string        = "OeAFFmMliFG5orCUuwAKQ8l4WWFQ67YX"
	BASE_API_URI           string        = "https://eodhistoricaldata.com"
	FUNDAMENTALS_API_ROUTE string        = "/api/fundamentals/"
	SPLITS_API_ROUTE       string        = "/api/splits/"
	DIVIDENDS_API_ROUTE    string        = "/api/div/"
	API_TOKEN_PARAM        string        = "?api_token=" + EOD_API_KEY
	FORMAT_JSON            string        = "&fmt=json"
)

type DbModel struct {
	client     *mongo.Client
	funda_col  *mongo.Collection
	splits_col *mongo.Collection
}

func getSplitsDataHttp(ticker string) ([]servertype.SplitsRecordRoot, error) {
	record := make([]servertype.SplitsRecordRoot, 4)
	eod_request_string := BASE_API_URI + SPLITS_API_ROUTE + ticker + API_TOKEN_PARAM + FORMAT_JSON
	var http_client = &http.Client{Timeout: HTTP_REQ_TIMEOUT_S * time.Second}
	resp, err := http_client.Get(eod_request_string)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return record, nil
}

func getFundamentalDataHttp(ticker string) (*servertype.EODFundamentalsRoot, error) {
	var fundamental_data = new(servertype.EODFundamentalsRoot)
	eod_request_string := BASE_API_URI + FUNDAMENTALS_API_ROUTE + ticker + API_TOKEN_PARAM
	var http_client = &http.Client{Timeout: HTTP_REQ_TIMEOUT_S * time.Second}
	resp, err := http_client.Get(eod_request_string)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(fundamental_data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return fundamental_data, err
}

func (r *DbModel) insertFundamentalsDbRecord(record *servertype.FundamentalsRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), QUERY_TIMEOUT_S*time.Second)
	defer cancel()
	_, err := r.funda_col.InsertOne(ctx, record)
	return err
}

func (r *DbModel) insertSplitsDbRecord(record *servertype.SplitsRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), QUERY_TIMEOUT_S*time.Second)
	defer cancel()
	_, err := r.splits_col.InsertOne(ctx, record)
	return err
}

func (r *DbModel) findRecentDbSplitsRecord(ticker string) (*servertype.SplitsRecord, error) {
	result := new(servertype.SplitsRecord)
	const layoutISO = "2006-01-02"
	time_now := time.Now()

	t, err := time.Parse(layoutISO, time_now.Format(layoutISO))
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"insertDate": bson.M{"$gte": t}, "ticker": "AAPL"}

	ctx, cancel := context.WithTimeout(context.Background(), QUERY_TIMEOUT_S*time.Second)
	defer cancel()
	err = r.splits_col.FindOne(ctx, filter).Decode(result)
	return result, err
}

func (r *DbModel) findRecentDbFundamentalsRecord(ticker string) (*servertype.FundamentalsRecord, error) {
	result := new(servertype.FundamentalsRecord)
	const layoutISO = "2006-01-02"
	time_now := time.Now()

	t, err := time.Parse(layoutISO, time_now.Format(layoutISO))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	filter := bson.M{"insertDate": bson.M{"$gte": t}, "ticker": "AAPL"}

	ctx, cancel := context.WithTimeout(context.Background(), QUERY_TIMEOUT_S*time.Second)
	defer cancel()
	err = r.funda_col.FindOne(ctx, filter).Decode(result)
	return result, err
}

func (r *DbModel) IsConnectionAlive() error {
	var err error
	if r.client != nil {
		err = r.client.Ping(context.TODO(), nil)
	} else {
		err = errors.New("Client pointer is Nil")
	}
	return err
}

func (r *DbModel) Connect() error {
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	var err error
	// Connect to MongoDB
	r.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Connected to MongoDB!")
	r.funda_col = r.client.Database(MONGO_DB).Collection(MONGO_FUNDA_COL)
	r.splits_col = r.client.Database(MONGO_DB).Collection(MONGO_SPLITS_COL)
	return nil
}

func (r *DbModel) Disconnect() error {
	var err error
	if r.client != nil {
		err = r.client.Disconnect(context.TODO())
	} else {
		err = errors.New("Client pointer is Nil")
	}

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}

func (r *DbModel) ReqCachedFundamentals(ticker string) (*servertype.FundamentalsRecord, error) {
	var record *servertype.FundamentalsRecord

	record, err := r.findRecentDbFundamentalsRecord(ticker)
	if err == mongo.ErrNoDocuments {

		fundamental_data, err := getFundamentalDataHttp(ticker)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tmprecord := servertype.FundamentalsRecord{
			Ticker:     ticker,
			InsertDate: time.Now(),
			Record:     fundamental_data,
		}
		err = r.insertFundamentalsDbRecord(&tmprecord)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &tmprecord, err
	} else if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, err
}

func (r *DbModel) ReqCachedSplitData(ticker string) (*servertype.SplitsRecord, error) {
	var record *servertype.SplitsRecord

	record, err := r.findRecentDbSplitsRecord(ticker)
	if err == mongo.ErrNoDocuments {

		splits_data, err := getSplitsDataHttp(ticker)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tmprecord := servertype.SplitsRecord{
			Ticker:     ticker,
			InsertDate: time.Now(),
			Record:     splits_data,
		}
		err = r.insertSplitsDbRecord(&tmprecord)
		if err != nil {
			return nil, err
		}
		return &tmprecord, err
	} else if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return record, err
}
