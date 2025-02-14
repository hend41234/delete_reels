package del

import (
	"deletereels/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type OK struct {
	Success bool `json:"success"`
}

func Delete(PostID string, PageToken string) bool {
	urls := fmt.Sprintf("%v/%v?access_token=%v", utils.Utils.BaseUrlGraphApi, PostID, PageToken)

	req, _ := http.NewRequest("DELETE", urls, nil)
	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	model := OK{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&model)
	if decodeErr != nil {
		log.Println("error decoe : " + decodeErr.Error())
	}
	fmt.Println(model)
	return model.Success
}
