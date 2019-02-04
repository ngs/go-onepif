package onepif

type Onepif struct {
	LocationKey    string          `json:"locationKey,omitempty"`
	Title          string          `json:"title,omitempty"`
	Location       string          `json:"location,omitempty"`
	SecureContents []SecureContent `json:"secureContents,omitempty"`
	TypeName       string          `json:"typeName,omitempty"`
}

type SecureContent struct {
	URLs     []URL       `json:"URLs"`
	Fields   []FormField `json:"fields"`
	Sections []Section   `json:"sections"`
}

type URL struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type FormField struct {
	Value       string `json:"value"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Designation string `json:"designation"`
}

type Section struct {
	Title  string  `json:"title"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Kind  string `json:"k"`
	Value string `json:"v"`
	Title string `json:"t"`
}
