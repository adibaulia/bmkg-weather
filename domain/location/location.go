package location

type Province string

const (
	DKIJakarta Province = "DKIJakarta"
)

type (
	Location struct {
		ID          string     `json:"id"`
		Latitude    string     `json:"latitude"`
		Longitude   string     `json:"longitude"`
		Coordinate  string     `json:"coordinate"`
		Type        string     `json:"type"`
		Region      string     `json:"region"`
		Level       string     `json:"level"`
		Description string     `json:"description"`
		Domain      string     `json:"domain"`
		Tags        string     `json:"tags"`
		NameLang    []NameLang `json:"name_lang"`
		Province    Province   `json:"province"`
	}

	NameLang struct {
		Name string `json:"name"`
		Lang string `json:"lang"`
	}
)
