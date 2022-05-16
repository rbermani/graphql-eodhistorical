package model

type Technicals struct {
	Beta                  *float64 `json:"Beta"`
	_52WeekHigh           *float64 `json:"52WeekHigh"`
	_52WeekLow            *float64 `json:"52WeekLow"`
	_50DayMA              *float64 `json:"50DayMA"`
	_200DayMA             *float64 `json:"200DayMA"`
	SharesShort           *int     `json:"SharesShort"`
	SharesShortPriorMonth *int     `json:"SharesShortPriorMonth"`
	ShortRatio            *float64 `json:"ShortRatio"`
	ShortPercent          *float64 `json:"ShortPercent"`
}
