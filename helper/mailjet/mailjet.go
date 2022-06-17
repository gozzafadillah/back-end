package mailjet

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseMailjet struct {
	Messages []struct {
		Status string `json:"Status"`
		To     []struct {
			Email       string `json:"Email"`
			MessageID   string `json:"MessageID"`
			MessageHref string `json:"MessageHref"`
		} `json:"To"`
	} `json:"Messages"`
}

func Mailjet(data []byte) {

	url := "https://api.mailjet.com/v3.1/send"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))

	basic := base64.StdEncoding.EncodeToString([]byte("177a3a51988d43f5512cf71bff810623" + ":" + "ba69cb7437c1bad179c8af199ba33dd1"))

	req.Header.Add("Authorization", "Basic "+basic)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result = ResponseMailjet{}
	json.Unmarshal(body, &result)
	fmt.Println(result)
	// fmt.Println(string(data))
}
