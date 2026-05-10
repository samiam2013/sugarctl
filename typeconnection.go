package main

type connectionsResponse struct {
	Status int `json:"status"`
	Data   []struct {
		ID         string `json:"id"`
		PatientID  string `json:"patientId"`
		Country    string `json:"country"`
		Status     int    `json:"status"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		TargetLow  int    `json:"targetLow"`
		TargetHigh int    `json:"targetHigh"`
		Uom        int    `json:"uom"`
		Sensor     struct {
			DeviceID string `json:"deviceId"`
			Sn       string `json:"sn"`
			A        int    `json:"a"`
			W        int    `json:"w"`
			Pt       int    `json:"pt"`
			S        bool   `json:"s"`
			Lj       bool   `json:"lj"`
		} `json:"sensor"`
		AlarmRules struct {
			C bool `json:"c"`
			H struct {
				On   bool    `json:"on"`
				Th   int     `json:"th"`
				Thmm float64 `json:"thmm"`
				D    int     `json:"d"`
				F    float64 `json:"f"`
			} `json:"h"`
			F struct {
				On   bool    `json:"on"`
				Th   int     `json:"th"`
				Thmm int     `json:"thmm"`
				D    int     `json:"d"`
				Tl   int     `json:"tl"`
				Tlmm float64 `json:"tlmm"`
				Isf  bool    `json:"isf"`
			} `json:"f"`
			L struct {
				On   bool    `json:"on"`
				Th   int     `json:"th"`
				Thmm float64 `json:"thmm"`
				D    int     `json:"d"`
				Tl   int     `json:"tl"`
				Tlmm float64 `json:"tlmm"`
			} `json:"l"`
			Nd struct {
				I int `json:"i"`
				R int `json:"r"`
				L int `json:"l"`
			} `json:"nd"`
			P   int `json:"p"`
			R   int `json:"r"`
			Std struct {
			} `json:"std"`
		} `json:"alarmRules"`
		GlucoseMeasurement struct {
			FactoryTimestamp string `json:"FactoryTimestamp"`
			Timestamp        string `json:"Timestamp"`
			Type             int    `json:"type"`
			ValueInMgPerDl   int    `json:"ValueInMgPerDl"`
			TrendArrow       int    `json:"TrendArrow"`
			TrendMessage     any    `json:"TrendMessage"`
			MeasurementColor int    `json:"MeasurementColor"`
			GlucoseUnits     int    `json:"GlucoseUnits"`
			Value            int    `json:"Value"`
			IsHigh           bool   `json:"isHigh"`
			IsLow            bool   `json:"isLow"`
		} `json:"glucoseMeasurement"`
		GlucoseItem struct {
			FactoryTimestamp string `json:"FactoryTimestamp"`
			Timestamp        string `json:"Timestamp"`
			Type             int    `json:"type"`
			ValueInMgPerDl   int    `json:"ValueInMgPerDl"`
			TrendArrow       int    `json:"TrendArrow"`
			TrendMessage     any    `json:"TrendMessage"`
			MeasurementColor int    `json:"MeasurementColor"`
			GlucoseUnits     int    `json:"GlucoseUnits"`
			Value            int    `json:"Value"`
			IsHigh           bool   `json:"isHigh"`
			IsLow            bool   `json:"isLow"`
		} `json:"glucoseItem"`
		GlucoseAlarm  any `json:"glucoseAlarm"`
		PatientDevice struct {
			Did                 string `json:"did"`
			Dtid                int    `json:"dtid"`
			V                   string `json:"v"`
			Ll                  int    `json:"ll"`
			H                   bool   `json:"h"`
			Hl                  int    `json:"hl"`
			U                   int    `json:"u"`
			FixedLowAlarmValues struct {
				Mgdl  int `json:"mgdl"`
				Mmoll int `json:"mmoll"`
			} `json:"fixedLowAlarmValues"`
			Alarms            bool `json:"alarms"`
			FixedLowThreshold int  `json:"fixedLowThreshold"`
		} `json:"patientDevice"`
		Created int `json:"created"`
	} `json:"data"`
	Ticket struct {
		Token    string `json:"token"`
		Expires  int    `json:"expires"`
		Duration int64  `json:"duration"`
	} `json:"ticket"`
}
