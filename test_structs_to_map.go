package main

import "fmt"
import "github.com/fatih/structs"

type UserFeature struct {
	ID                 int    `structs:"id" `
	Name               string `structs:"name" `
	UserFeatureID      int    `structs:"user_feature_id" `
	UserFeatureName    string `structs:"user_feature_name"  `
	UserFeatureContent string `structs:"user_feature_content" `
	Status             int    `structs:"status"              `
	Deleted            int    `structs:"deleted"              `
	CreatedAt          int64  `structs:"created_at"            `
	UpdatedAt          int64  `structs:"updated_at"            `
}

//type UserFeature struct {
//	ID                 int
//	Name               string
//	UserFeatureID      int
//	UserFeatureName    string
//	UserFeatureContent string
//	Status             int
//	Deleted            int
//	CreatedAt          int64
//	UpdatedAt          int64
//}

func main() {
	user := &UserFeature{
		ID:                 1,
		Name:               "abc",
		UserFeatureID:      1000,
		UserFeatureName:    "cdef",
		UserFeatureContent: "oiuy",
		Status:             1,
		Deleted:            1,
		CreatedAt:          10000,
		UpdatedAt:          1000,
	}

	rlt := structs.Map(user)
	fmt.Printf("%+v\n", rlt)
}
