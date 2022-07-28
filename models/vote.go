package models

type VoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"`
}

//func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
//	required := struct {
//		PostID    json.Number `json:"post_id"`
//		Direction float64     `json:"direction,string"`
//	}{}
//	err = json.Unmarshal(data, &required)
//	fmt.Println(required.PostID)
//	fmt.Println(required.Direction)
//	if err != nil {
//		return
//	} else if len(required.PostID) == 0 {
//		err = errors.New("缺少必填字段post_id")
//	} else if required.Direction == 0 {
//		err = errors.New("缺少必填字段direction")
//	} else {
//		v.PostID = required.PostID.String()
//		v.Direction = required.Direction
//	}
//	return
//}
