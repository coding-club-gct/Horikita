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
type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Provider   string `json:"provider"`
	Confirmed  bool   `json:"confirmed"`
	Blocked    bool   `json:"blocked"`
	DiscordUID string `json:"discordUID"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	UserDetail struct {
		ID           int    `json:"id"`
		RollNo       string `json:"rollNo"`
		Name         string `json:"name"`
		Dept         string `json:"dept"`
		PhnNumber    string `json:"phnNumber"`
		WhtspNumber  string `json:"whtspNumber"`
		GctMailId    string `json:"gctMailId"`
		LinkedinLink string `json:"linkedinLink"`
		Address      string `json:"address"`
		Gender       string `json:"gender"`
		Dob          string `json:"dob"`
		MailId       string `json:"mailId"`
		Batch        int    `json:"batch"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
		PublishedAt  string `json:"publishedAt"`
	} `json:"userDetail"`
}
