package models

type Event struct {
	ID         int `json:"id"`
	Attributes struct {
		Content     string      `json:"content"`
		Meta        interface{} `json:"meta"`
		Name        string      `json:"name"`
		ShortLiner  string      `json:"shortLiner"`
		MinTeamSize int         `json:"minTeamSize"`
		MaxTeamSize int         `json:"maxTeamSize"`
		Open        bool        `json:"open"`
		CreatedAt   string      `json:"createdAt"`
		UpdatedAt   string      `json:"updatedAt"`
		PublishedAt string      `json:"publishedAt"`
	} `json:"attributes"`
}
type Team struct {
	MemberIDs []int  `json:"members"`
	Name      string `json:"name"`
	EventID   int    `json:"event"`
}
