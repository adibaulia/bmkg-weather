package bmkg

import "encoding/xml"

type (
	BMKGForecast struct {
		XMLName          xml.Name `xml:"data"`
		Text             string   `xml:",chardata"`
		Source           string   `xml:"source,attr"`
		Productioncenter string   `xml:"productioncenter,attr"`
		Forecast         struct {
			Text   string `xml:",chardata"`
			Domain string `xml:"domain,attr"`
			Issue  struct {
				Text      string `xml:",chardata"`
				Timestamp string `xml:"timestamp"`
				Year      string `xml:"year"`
				Month     string `xml:"month"`
				Day       string `xml:"day"`
				Hour      string `xml:"hour"`
				Minute    string `xml:"minute"`
				Second    string `xml:"second"`
			} `xml:"issue"`
			Area []struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"id,attr"`
				Latitude    string `xml:"latitude,attr"`
				Longitude   string `xml:"longitude,attr"`
				Coordinate  string `xml:"coordinate,attr"`
				Type        string `xml:"type,attr"`
				Region      string `xml:"region,attr"`
				Level       string `xml:"level,attr"`
				Description string `xml:"description,attr"`
				Domain      string `xml:"domain,attr"`
				Tags        string `xml:"tags,attr"`
				Name        []Name `xml:"name"`
				Parameter   []struct {
					Text        string `xml:",chardata"`
					ID          string `xml:"id,attr"`
					Description string `xml:"description,attr"`
					Type        string `xml:"type,attr"`
					Timerange   []struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						H        string `xml:"h,attr"`
						Datetime string `xml:"datetime,attr"`
						Day      string `xml:"day,attr"`
						Value    []struct {
							Text string `xml:",chardata"`
							Unit string `xml:"unit,attr"`
						} `xml:"value"`
					} `xml:"timerange"`
				} `xml:"parameter"`
			} `xml:"area"`
		} `xml:"forecast"`
	}

	Name struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	}
)
