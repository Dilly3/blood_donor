package models

import "github.com/lithammer/shortuuid"

type Candidate struct {
	Id         string `json:"id" bson:"id"`
	FullName   string `json:"fullname" bson:"fullname"`
	Mobile     string `json:"mobile" bson:"mobile"`
	Email      string `json:"email" bson:"email" validate:"email,required"`
	Age        int64  `json:"age" bson:"age"`
	BloodGroup string `json:"blood_group" bson:"blood_group"`
	Address    string `json:"address" bson:"address"`
}

func GetshortId() string {
	return shortuuid.New()
}

func NewCandidate() *Candidate {
	return &Candidate{
		Id: GetshortId(),
	}
}
