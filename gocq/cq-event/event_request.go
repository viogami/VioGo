package cqEvent

import (
	"encoding/json"
	"log"
)

type AddFriendRequest struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	UserID      int64  `json:"user_id"`
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
}

type AddGroupRequest struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	SubType     string `json:"sub_type"`
	GroupID     int64  `json:"group_id"`
	UserID      int64  `json:"user_id"`
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
}

func Get_request_info(p []byte, RequestType string) any {
	var add_friend_request AddFriendRequest
	var add_group_request AddGroupRequest

	switch RequestType {
	case "friend":
		err := json.Unmarshal(p, &add_friend_request)
		if err != nil {
			log.Println("Error parsing JSON to add_friend_request:", err)
		}
		return add_friend_request
	case "group":
		err := json.Unmarshal(p, &add_group_request)
		if err != nil {
			log.Println("Error parsing JSON to add_group_request:", err)
		}
		return add_group_request
	}
	return nil
}
