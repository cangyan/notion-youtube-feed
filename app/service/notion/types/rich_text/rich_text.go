package rich_text

type RichText struct {
	RichText []InnerRichText `json:"rich_text,omitempty"`
}

type InnerRichText struct {
	Text struct {
		Content string `json:"content,omitempty"`
	} `json:"text,omitempty"`
}

func ValueOf(s string) RichText {
	item := InnerRichText{}
	item.Text.Content = s
	return RichText{
		RichText: []InnerRichText{item},
	}
}
