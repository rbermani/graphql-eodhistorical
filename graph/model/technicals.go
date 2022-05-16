package model

type Technicals struct {
	Beta                  float64 `json:"Beta"`
	FiftyTwoWeekHigh      float64 `json:"52WeekHigh"`
	FiftyTwoWeekLow       float64 `json:"52WeekLow"`
	FiftyDayMA            float64 `json:"50DayMA"`
	TwoHundredDayMA       float64 `json:"200DayMA"`
	SharesShort           int     `json:"SharesShort"`
	SharesShortPriorMonth int     `json:"SharesShortPriorMonth"`
	ShortRatio            float64 `json:"ShortRatio"`
	ShortPercent          float64 `json:"ShortPercent"`
}
