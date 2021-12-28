package filter_object

import "encoding/json"

type TextEqualsFilterObject struct {
	Property string `json:"property,omitempty"`
	Text     struct {
		Equals string `json:"equals,omitempty"`
	} `json:"text,omitempty"`
}

func GenerateTextOrFilterObject(property string, values []string) string {
	var or []TextEqualsFilterObject
	for _, id := range values {
		tmp := TextEqualsFilterObject{}
		tmp.Property = property
		tmp.Text.Equals = id
		or = append(or, tmp)
	}

	ret := make(map[string]interface{})
	ret["or"] = or

	b, _ := json.Marshal(ret)
	return string(b)
}
