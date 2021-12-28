package number

type Number struct {
	Number float64 `json:"number,omitempty"`
}

func ValueOf(n float64) Number {
	return Number{Number: n}
}
