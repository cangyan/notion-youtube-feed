package files

type Files struct {
	Files []InnerFiles `json:"files,omitempty"`
}

type InnerFiles struct {
	Name string `json:"name,omitempty"`
	// Type     string `json:"type,omitempty"`
	External struct {
		Url string `json:"url,omitempty"`
	} `json:"external,omitempty"`
}

func ValueOf(s string) Files {
	item := InnerFiles{}

	item.Name = s
	if len(s) > 100 {
		item.Name = s[0:95] + "..."
	}
	// item.Type = "external"
	item.External.Url = s

	return Files{Files: []InnerFiles{item}}
}
