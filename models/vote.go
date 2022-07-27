package models

import (
	"encoding/json"
	"errors"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"direction"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string  `json:"post_id"`
		Direction float64 `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}
