package del

import (
	"deletereels/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type DataModels struct {
	ID string `json:"id"`
}
type ResModels struct {
	Data []DataModels `json:"data"`
}

func GetPostID(ID string, Token string, o_p ...interface{}) ResModels {
	var urls string
	if o_p[0] == "all" {
		urls = fmt.Sprintf("%v/%v/video_reels?access_token=%v", utils.Utils.BaseUrlGraphApi, ID, Token)
	} else {
		log.Println("methode is not default")
		check_methode := reflect.TypeOf(o_p[0])
		if check_methode.String() == "string" {
			since := o_p[0].(string)
			urls = fmt.Sprintf("%v/%v/video_reels?since=%v&access_token=%v", utils.Utils.BaseUrlGraphApi, ID, since, Token)
		} else if check_methode.String() == "int" {
			limit := o_p[0].(int)
			urls = fmt.Sprintf("%v/%v/video_reels?limit=%v&access_token=%v", utils.Utils.BaseUrlGraphApi, ID, limit, Token)
		}
	}

	req, _ := http.NewRequest("GET", urls, nil)
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	model := ResModels{}
	// read, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(read))
	decodeErr := json.NewDecoder(resp.Body).Decode(&model)
	if decodeErr != nil {
		log.Println("erorr decode : " + decodeErr.Error())
	}
	return model
}
