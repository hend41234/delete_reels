package main

import (
	del "deletereels/delete"
	gettoken "deletereels/get_token"
	"deletereels/utils"
	"fmt"
	"log"
	"strconv"
)

func choosePage(tokens gettoken.ResGetAccessTokenModels) (ID, Name, PageToken string) {
	Data := tokens.Data
	for i, data := range Data {
		fmt.Println(i, ". ID : "+data.ID+" | Name : "+data.Name)
	}
	fmt.Println("Choose Pages 🔥 : ")
	var id string
	fmt.Scanln(&id)
	fixID, _ := strconv.Atoi(id)
	if fixID > len(Data) {
		fmt.Println("ID not found")
		return "", "", ""
	}
	return Data[fixID].ID, Data[fixID].Name, Data[fixID].AccessToken
}

type OK struct {
	Code int `json:"code"`
}

func x(o_p ...interface{}) {
	x := o_p[0]
	b := o_p[1]
	log.Println(x, b)
}

func main() {
	getTokenRequire := gettoken.GetAccessTokenModels{
		EndPoint:    fmt.Sprintf("%v/%v/accounts", utils.Utils.BaseUrlGraphApi, utils.Utils.UserID),
		AccessToken: utils.Utils.SystemUserAccess,
	}
	get, err := getTokenRequire.GetAccessToken()
	if err != nil {
		log.Println(err)
	}
	ok := 1
	pageID := ""
	for ok == 1 {
		id, name, TOKEN := choosePage(get)
		log.Println("ID : " + pageID + " | Name : " + name + " will be deleted🕐🕐🕐")
		if id == "" {
			continue
		}
		delete := []del.OK{}

		// check methode delete by inputed users
		method := MethodeDelete()
		if method != "default 25" {
			method = IsNotDefault(method)
			checkMethod, err := strconv.Atoi(method)
			if err != nil {
				since := method
				postID := del.GetPostID(id, TOKEN, since)
				delete = postID.Delete(TOKEN)
			} else {
				limit := checkMethod
				postID := del.GetPostID(id, TOKEN, limit)
				delete = postID.Delete(TOKEN)
			}
		} else {
			postID := del.GetPostID(id, TOKEN)
			delete = postID.Delete(TOKEN)
		}

		check := 0
		for i := range delete {
			if delete[i].Code == 200 {
				check++
			}
		}
		if check == len(delete) {
			log.Println("Delete success✅✅✅!!!")
		} else if check != len(delete) {
			log.Println("Delete success, But have any error✅⛔⛔!!!")
		} else {
			log.Println("Delete failed❌❌❌!!!")
		}
		ok = 0
	}

}
