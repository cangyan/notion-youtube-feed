package date

type Date struct {
	Date struct {
		Start string `json:"start,omitempty"`
		// TimeZone string `json:"time_zone,omitempty"`
	} `json:"date,omitempty"`
}

func ValueOf(s string) Date {
	var d Date
	d.Date.Start = s
	// d.Date.TimeZone = "Asia/Shanghai"

	return d
}
