package server_type

type ActivityInvolvement struct {
	Activity    string `json:"Activity"`
	Involvement string `json:"Involvement"`
}

type Technicals struct {
	Beta                  float64 `json:"Beta"`
	_52WeekHigh           string  `json:"52WeekHigh"`
	_52WeekLow            string  `json:"52WeekLow"`
	_50DayMA              string  `json:"50DayMA"`
	_200DayMA             string  `json:"200DayMA"`
	SharesShort           int     `json:"SharesShort"`
	SharesShortPriorMonth int     `json:"SharesShortPriorMonth"`
	ShortRatio            float64 `json:"ShortRatio"`
	ShortPercent          float64 `json:"ShortPercent"`
}

type AddressData struct {
	Street  string `json:"Street"`
	City    string `json:"City"`
	State   string `json:"State"`
	Country string `json:"Country"`
}

type AnalystRatings struct {
	Rating      float64 `json:"Rating"`
	TargetPrice float64 `json:"TargetPrice"`
	StrongBuy   int     `json:"StrongBuy"`
	Buy         int     `json:"Buy"`
	Hold        int     `json:"Hold"`
	Sell        int     `json:"Sell"`
	StrongSell  int     `json:"StrongSell"`
}

type Annual struct {
	Date          string  `json:"date"`
	DateFormatted string  `json:"dateFormatted"`
	SharesMln     string  `json:"sharesMln"`
	Shares        float64 `json:"shares"`
}

type AnnualEarnings struct {
	Date      string  `json:"date"`
	EpsActual float64 `json:"epsActual"`
}

type BalanceSheet struct {
	CurrencySymbol string                      `json:"currency_symbol"`
	Quarterly      map[string]BalanceSheetItem `json:"quarterly"`
	Yearly         map[string]BalanceSheetItem `json:"yearly"`
}

type BalanceSheetItem struct {
	Date                                             string `json:"date"`
	FilingDate                                       string `json:"filing_date"`
	CurrencySymbol                                   string `json:"currency_symbol"`
	TotalAssets                                      string `json:"totalAssets"`
	IntangibleAssets                                 string `json:"intangibleAssets"`
	EarningAssets                                    string `json:"earningAssets"`
	OtherCurrentAssets                               string `json:"otherCurrentAssets"`
	TotalLiab                                        string `json:"totalLiab"`
	TotalStockholderEquity                           string `json:"totalStockholderEquity"`
	DeferredLongTermLiab                             string `json:"deferredLongTermLiab"`
	OtherCurrentLiab                                 string `json:"otherCurrentLiab"`
	CommonStock                                      string `json:"commonStock"`
	RetainedEarnings                                 string `json:"retainedEarnings"`
	OtherLiab                                        string `json:"otherLiab"`
	GoodWill                                         string `json:"goodWill"`
	OtherAssets                                      string `json:"otherAssets"`
	Cash                                             string `json:"cash"`
	TotalCurrentLiabilities                          string `json:"totalCurrentLiabilities"`
	NetDebt                                          string `json:"netDebt"`
	ShortTermDebt                                    string `json:"shortTermDebt"`
	ShortLongTermDebt                                string `json:"shortLongTermDebt"`
	ShortLongTermDebtTotal                           string `json:"shortLongTermDebtTotal"`
	OtherStockholderEquity                           string `json:"otherStockholderEquity"`
	PropertyPlantEquipment                           string `json:"propertyPlantEquipment"`
	TotalCurrentAssets                               string `json:"totalCurrentAssets"`
	LongTermInvestments                              string `json:"longTermInvestments"`
	NetTangibleAssets                                string `json:"netTangibleAssets"`
	ShortTermInvestments                             string `json:"shortTermInvestments"`
	NetReceivables                                   string `json:"netReceivables"`
	LongTermDebt                                     string `json:"longTermDebt"`
	Inventory                                        string `json:"inventory"`
	AccountsPayable                                  string `json:"accountsPayable"`
	TotalPermanentEquity                             string `json:"totalPermanentEquity"`
	NoncontrollingInterestInConsolidatedEntity       string `json:"noncontrollingInterestInConsolidatedEntity"`
	TemporaryEquityRedeemableNoncontrollingInterests string `json:"temporaryEquityRedeemableNoncontrollingInterests"`
	AccumulatedOtherComprehensiveIncome              string `json:"accumulatedOtherComprehensiveIncome"`
	AdditionalPaidInCapital                          string `json:"additionalPaidInCapital"`
	CommonStockTotalEquity                           string `json:"commonStockTotalEquity"`
	PreferredStockTotalEquity                        string `json:"preferredStockTotalEquity"`
	RetainedEarningsTotalEquity                      string `json:"retainedEarningsTotalEquity"`
	TreasuryStock                                    string `json:"treasuryStock"`
	AccumulatedAmortization                          string `json:"accumulatedAmortization"`
	NonCurrrentAssetsOther                           string `json:"nonCurrrentAssetsOther"`
	DeferredLongTermAssetCharges                     string `json:"deferredLongTermAssetCharges"`
	NonCurrentAssetsTotal                            string `json:"nonCurrentAssetsTotal"`
	CapitalLeaseObligations                          string `json:"capitalLeaseObligations"`
	LongTermDebtTotal                                string `json:"longTermDebtTotal"`
	NonCurrentLiabilitiesOther                       string `json:"nonCurrentLiabilitiesOther"`
	NonCurrentLiabilitiesTotal                       string `json:"nonCurrentLiabilitiesTotal"`
	NegativeGoodwill                                 string `json:"negativeGoodwill"`
	Warrants                                         string `json:"warrants"`
	PreferredStockRedeemable                         string `json:"preferredStockRedeemable"`
	CapitalSurpluse                                  string `json:"capitalSurpluse"`
	LiabilitiesAndStockholdersEquity                 string `json:"liabilitiesAndStockholdersEquity"`
	CashAndShortTermInvestments                      string `json:"cashAndShortTermInvestments"`
	PropertyPlantAndEquipmentGross                   string `json:"propertyPlantAndEquipmentGross"`
	PropertyPlantAndEquipmentNet                     string `json:"propertyPlantAndEquipmentNet"`
	AccumulatedDepreciation                          string `json:"accumulatedDepreciation"`
	NetWorkingCapital                                string `json:"netWorkingCapital"`
	NetInvestedCapital                               string `json:"netInvestedCapital"`
	CommonStockSharesOutstanding                     string `json:"commonStockSharesOutstanding"`
}

type CashFlow struct {
	CurrencySymbol string                  `json:"currency_symbol"`
	Quarterly      map[string]CashFlowItem `json:"quarterly"`
	Yearly         map[string]CashFlowItem `json:"yearly"`
}

type CashFlowItem struct {
	Date                                  string `json:"date"`
	FilingDate                            string `json:"filing_date"`
	CurrencySymbol                        string `json:"currency_symbol"`
	Investments                           string `json:"investments"`
	ChangeToLiabilities                   string `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities string `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         string `json:"netBorrowings"`
	TotalCashFromFinancingActivities      string `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           string `json:"changeToOperatingActivities"`
	NetIncome                             string `json:"netIncome"`
	ChangeInCash                          string `json:"changeInCash"`
	BeginPeriodCashFlow                   string `json:"beginPeriodCashFlow"`
	EndPeriodCashFlow                     string `json:"endPeriodCashFlow"`
	TotalCashFromOperatingActivities      string `json:"totalCashFromOperatingActivities"`
	Depreciation                          string `json:"depreciation"`
	OtherCashflowsFromInvestingActivities string `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         string `json:"dividendsPaid"`
	ChangeToInventory                     string `json:"changeToInventory"`
	ChangeToAccountReceivables            string `json:"changeToAccountReceivables"`
	SalePurchaseOfStock                   string `json:"salePurchaseOfStock"`
	OtherCashflowsFromFinancingActivities string `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     string `json:"changeToNetincome"`
	CapitalExpenditures                   string `json:"capitalExpenditures"`
	ChangeReceivables                     string `json:"changeReceivables"`
	CashFlowsOtherOperating               string `json:"cashFlowsOtherOperating"`
	ExchangeRateChanges                   string `json:"exchangeRateChanges"`
	CashAndCashEquivalentsChanges         string `json:"cashAndCashEquivalentsChanges"`
	ChangeInWorkingCapital                string `json:"changeInWorkingCapital"`
	OtherNonCashItems                     string `json:"otherNonCashItems"`
	FreeCashFlow                          string `json:"freeCashFlow"`
}

type ESGScores struct {
	Disclaimer                 string                         `json:"Disclaimer"`
	RatingDate                 string                         `json:"RatingDate"`
	TotalEsg                   float64                        `json:"TotalEsg"`
	TotalEsgPercentile         float64                        `json:"TotalEsgPercentile"`
	EnvironmentScore           float64                        `json:"EnvironmentScore"`
	EnvironmentScorePercentile int                            `json:"EnvironmentScorePercentile"`
	SocialScore                float64                        `json:"SocialScore"`
	SocialScorePercentile      int                            `json:"SocialScorePercentile"`
	GovernanceScore            float64                        `json:"GovernanceScore"`
	GovernanceScorePercentile  int                            `json:"GovernanceScorePercentile"`
	ControversyLevel           int                            `json:"ControversyLevel"`
	ActivitiesInvolvement      map[string]ActivityInvolvement `json:"ActivitiesInvolvement"`
}

type Earnings struct {
	History map[string]History        `json:"History"`
	Trend   map[string]Trend          `json:"Trend"`
	Annual  map[string]AnnualEarnings `json:"Annual"`
}

type EquityFundamentals struct {
	General             General                       `json:"General"`
	Highlights          Highlights                    `json:"Highlights"`
	Valuation           Valuation                     `json:"Valuation"`
	SharesStats         SharesStats                   `json:"SharesStats"`
	Technicals          Technicals                    `json:"Technicals"`
	SplitsDividends     SplitsDividends               `json:"SplitsDividends"`
	AnalystRatings      AnalystRatings                `json:"AnalystRatings"`
	Holders             Holders                       `json:"Holders"`
	InsiderTransactions map[string]InsiderTransaction `json:"InsiderTransactions"`
	ESGScores           ESGScores                     `json:"ESGScores"`
	OutstandingShares   OutstandingShares             `json:"OutstandingShares"`
	Earnings            Earnings                      `json:"Earnings"`
	Financials          Financials                    `json:"Financials"`
}

type Financials struct {
	BalanceSheet    BalanceSheet    `json:"Balance_Sheet"`
	CashFlow        CashFlow        `json:"Cash_Flow"`
	IncomeStatement IncomeStatement `json:"Income_Statement"`
}

type General struct {
	Code                  string             `json:"Code"`
	Type                  string             `json:"Type"`
	Name                  string             `json:"Name"`
	Exchange              string             `json:"Exchange"`
	CurrencyCode          string             `json:"CurrencyCode"`
	CurrencyName          string             `json:"CurrencyName"`
	CurrencySymbol        string             `json:"CurrencySymbol"`
	CountryName           string             `json:"CountryName"`
	EmployerIDNumber      string             `json:"EmployerIdNumber"`
	FiscalYearEnd         string             `json:"FiscalYearEnd"`
	InternationalDomestic string             `json:"InternationalDomestic"`
	Sector                string             `json:"Sector"`
	Industry              string             `json:"Industry"`
	GicSector             string             `json:"GicSector"`
	GicGroup              string             `json:"GicGroup"`
	GicIndustry           string             `json:"GicIndustry"`
	GicSubIndustry        string             `json:"GicSubIndustry"`
	HomeCategory          string             `json:"HomeCategory"`
	IsDelisted            bool               `json:"IsDelisted"`
	Description           string             `json:"Description"`
	Address               string             `json:"Address"`
	AddressData           AddressData        `json:"AddressData"`
	Listings              map[string]Listing `json:"Listings"`
	Officers              map[string]Officer `json:"Officers"`
	Phone                 string             `json:"Phone"`
	FullTimeEmployees     int                `json:"FullTimeEmployees"`
	UpdatedAt             string             `json:"UpdatedAt"`
}

type Highlights struct {
	MarketCapitalization    int     `json:"MarketCapitalization"`
	MarketCapitalizationMln float64 `json:"MarketCapitalizationMln"`
	WallStreetTargetPrice   float64 `json:"WallStreetTargetPrice"`
	BookValue               float64 `json:"BookValue"`
	DividendShare           float64 `json:"DividendShare"`
	DividendYield           float64 `json:"DividendYield"`
	EarningsShare           float64 `json:"EarningsShare"`
	MostRecentQuarter       string  `json:"MostRecentQuarter"`
	ProfitMargin            float64 `json:"ProfitMargin"`
}

type History struct {
	ReportDate        string  `json:"reportDate"`
	Date              string  `json:"date"`
	BeforeAfterMarket string  `json:"beforeAfterMarket"`
	Currency          string  `json:"currency"`
	EpsActual         float64 `json:"epsActual"`
	EpsEstimate       float64 `json:"epsEstimate"`
	EpsDifference     float64 `json:"epsDifference"`
	SurprisePercent   float64 `json:"surprisePercent"`
}

type Holders struct {
	Institutions map[string]Institution `json:"Institutions"`
	Funds        map[string]Institution `json:"Funds"`
}

type IncomeStatement struct {
	CurrencySymbol string                     `json:"currency_symbol"`
	Quarterly      map[string]IncomeStatement `json:"quarterly"`
	Yearly         map[string]IncomeStatement `json:"yearly"`
}

type IncomeStatementItem struct {
	Date                              string `json:"date"`
	FilingDate                        string `json:"filing_date"`
	CurrencySymbol                    string `json:"currency_symbol"`
	ResearchDevelopment               string `json:"researchDevelopment"`
	EffectOfAccountingCharges         string `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	MinorityInterest                  string `json:"minorityInterest"`
	NetIncome                         string `json:"netIncome"`
	SellingGeneralAdministrative      string `json:"sellingGeneralAdministrative"`
	SellingAndMarketingExpenses       string `json:"sellingAndMarketingExpenses"`
	GrossProfit                       string `json:"grossProfit"`
	ReconciledDepreciation            string `json:"reconciledDepreciation"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	NonOperatingIncomeNetOther        string `json:"nonOperatingIncomeNetOther"`
	OperatingIncome                   string `json:"operatingIncome"`
	OtherOperatingExpenses            string `json:"otherOperatingExpenses"`
	InterestExpense                   string `json:"interestExpense"`
	TaxProvision                      string `json:"taxProvision"`
	InterestIncome                    string `json:"interestIncome"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	ExtraordinaryItems                string `json:"extraordinaryItems"`
	NonRecurring                      string `json:"nonRecurring"`
	OtherItems                        string `json:"otherItems"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	TotalRevenue                      string `json:"totalRevenue"`
	TotalOperatingExpenses            string `json:"totalOperatingExpenses"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        string `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            string `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        string `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares string `json:"netIncomeApplicableToCommonShares"`
	PreferredStockAndOtherAdjustments string `json:"preferredStockAndOtherAdjustments"`
}

type InsiderTransaction struct {
	Date                        string  `json:"date"`
	OwnerCik                    string  `json:"ownerCik"`
	OwnerName                   string  `json:"ownerName"`
	TransactionDate             string  `json:"transactionDate"`
	TransactionCode             string  `json:"transactionCode"`
	TransactionAmount           int     `json:"transactionAmount"`
	TransactionPrice            float64 `json:"transactionPrice"`
	TransactionAcquiredDisposed string  `json:"transactionAcquiredDisposed"`
	PostTransactionAmount       int     `json:"postTransactionAmount"`
	SecLink                     string  `json:"secLink"`
}

type Institution struct {
	Name          string  `json:"name"`
	Date          string  `json:"date"`
	TotalShares   float64 `json:"totalShares"`
	TotalAssets   float64 `json:"totalAssets"`
	CurrentShares int     `json:"currentShares"`
	Change        int     `json:"change"`
	ChangeP       float64 `json:"change_p"`
}

type Listing struct {
	Code     string `json:"Code"`
	Exchange string `json:"Exchange"`
	Name     string `json:"Name"`
}

type NumberDividendsByYear struct {
	Year  int `json:"Year"`
	Count int `json:"Count"`
}

type Officer struct {
	Name     string `json:"Name"`
	Title    string `json:"Title"`
	YearBorn string `json:"YearBorn"`
}

type OutstandingShares struct {
	Annual    map[string]Annual    `json:"annual"`
	Quarterly map[string]Quarterly `json:"quarterly"`
}

type Quarterly struct {
	Date          string  `json:"date"`
	DateFormatted string  `json:"dateFormatted"`
	SharesMln     string  `json:"sharesMln"`
	Shares        float64 `json:"shares"`
}

type SharesStats struct {
	SharesOutstanding       int     `json:"SharesOutstanding"`
	SharesFloat             int     `json:"SharesFloat"`
	PercentInsiders         float64 `json:"PercentInsiders"`
	PercentInstitutions     float64 `json:"PercentInstitutions"`
	SharesShort             string  `json:"SharesShort"`
	SharesShortPriorMonth   string  `json:"SharesShortPriorMonth"`
	ShortRatio              string  `json:"ShortRatio"`
	ShortPercentOutstanding string  `json:"ShortPercentOutstanding"`
	ShortPercentFloat       string  `json:"ShortPercentFloat"`
}

type SplitsDividends struct {
	ForwardAnnualDividendRate  float64                          `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield float64                          `json:"ForwardAnnualDividendYield"`
	PayoutRatio                float64                          `json:"PayoutRatio"`
	DividendDate               string                           `json:"DividendDate"`
	ExDividendDate             string                           `json:"ExDividendDate"`
	LastSplitFactor            string                           `json:"LastSplitFactor"`
	LastSplitDate              string                           `json:"LastSplitDate"`
	NumberDividendsByYear      map[string]NumberDividendsByYear `json:"NumberDividendsByYear"`
}

type Trend struct {
	Date                             string `json:"date"`
	Period                           string `json:"period"`
	Growth                           string `json:"growth"`
	EarningsEstimateAvg              string `json:"earningsEstimateAvg"`
	EarningsEstimateLow              string `json:"earningsEstimateLow"`
	EarningsEstimateHigh             string `json:"earningsEstimateHigh"`
	EarningsEstimateYearAgoEps       string `json:"earningsEstimateYearAgoEps"`
	EarningsEstimateNumberOfAnalysts string `json:"earningsEstimateNumberOfAnalysts"`
	EarningsEstimateGrowth           string `json:"earningsEstimateGrowth"`
	RevenueEstimateAvg               string `json:"revenueEstimateAvg"`
	RevenueEstimateLow               string `json:"revenueEstimateLow"`
	RevenueEstimateHigh              string `json:"revenueEstimateHigh"`
	RevenueEstimateYearAgoEps        string `json:"revenueEstimateYearAgoEps"`
	RevenueEstimateNumberOfAnalysts  string `json:"revenueEstimateNumberOfAnalysts"`
	RevenueEstimateGrowth            string `json:"revenueEstimateGrowth"`
	EpsTrendCurrent                  string `json:"epsTrendCurrent"`
	EpsTrend7daysAgo                 string `json:"epsTrend7daysAgo"`
	EpsTrend30daysAgo                string `json:"epsTrend30daysAgo"`
	EpsTrend60daysAgo                string `json:"epsTrend60daysAgo"`
	EpsTrend90daysAgo                string `json:"epsTrend90daysAgo"`
	EpsRevisionsUpLast7days          string `json:"epsRevisionsUpLast7days"`
	EpsRevisionsUpLast30days         string `json:"epsRevisionsUpLast30days"`
	EpsRevisionsDownLast7days        string `json:"epsRevisionsDownLast7days"`
	EpsRevisionsDownLast30days       string `json:"epsRevisionsDownLast30days"`
}

type Valuation struct {
	EnterpriseValue        int     `json:"EnterpriseValue"`
	EnterpriseValueRevenue float64 `json:"EnterpriseValueRevenue"`
	EnterpriseValueEbitda  float64 `json:"EnterpriseValueEbitda"`
}
