package del

import (
	"deletereels/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DataModels struct {
	ID string `json:"id"`
}
type ResModels struct {
	Data []DataModels `json:"data"`
}

func GetPostID(ID string, Token string) ResModels {
	urls := fmt.Sprintf("%v/%v/video_reels?access_token=%v", utils.Utils.BaseUrlGraphApi, ID, Token)
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
