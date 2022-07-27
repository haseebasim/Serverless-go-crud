package models

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	CreatedTime  string `json:"createdtime"`
	ModifiedTime string `json:"modifiedtime"`
}
