package del

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type OK struct {
	Code int `json:"code"`
}
type Res []OK
type BatchModels struct {
	Method      string `json:"method"`
	RelativeURL string `json:"relative_url"`
}

func (ids ResModels) Delete(PageToken string) (ok Res) {
	urls := "https://graph.facebook.com/v22.0/"

	var BRequest []BatchModels
	for _, data := range ids.Data {
		BRequest = append(BRequest, BatchModels{
			Method:      "DELETE",
			RelativeURL: data.ID,
		})
	}

	payload, _ := json.Marshal(BRequest)
	formURL := url.Values{}
	formURL.Add("batch", (string(payload)))
	formURL.Add("access_token", PageToken)

	request, _ := http.NewRequest("POST", urls, bytes.NewBuffer([]byte(formURL.Encode())))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	res, _ := client.Do(request)
	defer res.Body.Close()
	readRes, _ := io.ReadAll(res.Body)
	log.Println(string(readRes))

	decodeErr := json.NewDecoder(res.Body).Decode(&ok)
	if decodeErr != nil {
		log.Println("erorr decode : " + decodeErr.Error())
	}

	return
}
