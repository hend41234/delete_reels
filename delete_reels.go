package main

import (
	del "deletereels/delete"
	gettoken "deletereels/get_token"
	"deletereels/utils"
	"fmt"
	"log"
)

func main() {
	getTokenRequire := gettoken.GetAccessTokenModels{
		EndPoint:    fmt.Sprintf("%v/%v/accounts", utils.Utils.BaseUrlGraphApi, utils.Utils.UserID),
		AccessToken: utils.Utils.SystemUserAccess,
	}
	get, err := getTokenRequire.GetAccessToken()
	if err != nil {
		log.Println(err)
	}
	for _, data := range get.Data {
		if data.ID == "229092887266424" {
			fmt.Println("get token " + data.ID)
			postID := del.GetPostID(data.ID, data.AccessToken)

			for _, id := range postID.Data {
				del.Delete(id.ID, data.AccessToken)
			}
		}
		continue
	}
}
