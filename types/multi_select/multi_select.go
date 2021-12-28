package multi_select

type MultiSelect struct {
	MultiSelect []InnerMultiSelect `json:"multi_select,omitempty"`
}

type InnerMultiSelect struct {
	Name string `json:"name,omitempty"`
}

func ValueOf(s []string) MultiSelect {
	var items []InnerMultiSelect

	for _, i := range s {
		items = append(items, InnerMultiSelect{Name: i})
	}

	return MultiSelect{MultiSelect: items}
}
