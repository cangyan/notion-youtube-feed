package notion

type NotionDatabaseQueryResp struct {
	Object     string      `json:"object,omitempty"`
	Results    []Block     `json:"results,omitempty"`
	NextCursor interface{} `json:"next_cursor,omitempty"`
	HasMore    bool        `json:"has_more,omitempty"`
}

type Block struct {
	Object     string `json:"object,omitempty"`
	Id         string `json:"id,omitempty"`
	Properties struct {
		ID struct {
			RichText []struct {
				Text struct {
					Content string `json:"content,omitempty"`
				} `json:"text,omitempty"`
			} `json:"rich_text,omitempty"`
		} `json:"id,omitempty"`
	} `json:"properties,omitempty"`
}

func (r *NotionDatabaseQueryResp) GetArticleIds() []string {
	data := make([]string, 0)
	if len(r.Results) > 0 {
		for _, item := range r.Results {
			data = append(data, item.Properties.ID.RichText[0].Text.Content)
		}
	}

	return data
}
