package title

type Title struct {
	Title []InnerTitle `json:"title,omitempty"`
}

type InnerTitle struct {
	Text struct {
		Content string `json:"content,omitempty"`
	} `json:"text,omitempty"`
}

func ValueOf(s string) Title {
	item := InnerTitle{}
	item.Text.Content = s
	return Title{
		Title: []InnerTitle{item},
	}
}
