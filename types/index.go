package types

type ImportantDate struct {
	Label       string `json:"label"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type Event struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ImportantDates []ImportantDate `json:"important_dates"`
}

type LoadedCustomId struct {
	CustomID string `json:"custom_id"`
	Payload string `json:"payload"`
}

