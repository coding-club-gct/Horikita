package models

type Event struct {
    Data struct {
        ID int `json:"id"`
        Attributes struct {
            Content string `json:"content"`
            Meta    interface{} `json:"meta"`
            Name    string `json:"name"`
            MinTeamSize int `json:"minTeamSize"`
            MaxTeamSize int `json:"maxTeamSize"`
            CreatedAt string `json:"createdAt"`
            UpdatedAt string `json:"updatedAt"`
            PublishedAt string `json:"publishedAt"`
        } `json:"attributes"`
    } `json:"data"`
    Meta interface{} `json:"meta"`
}

type Team struct {
    StudentIDs []int
    EventID int
}