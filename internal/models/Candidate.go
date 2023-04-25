package models

type Candidate struct {
	Id         int64  `json:"id"`
	FullName   string `json:"fullname"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Age        int64  `json:"age"`
	BloodGroup string `json:"blood_group"`
	Address    string `json:"address"`
}
