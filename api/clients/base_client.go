package clients

import (
	"io"
	"log"
	"net/http"
)

func GetRequest(requestUrl string) ([]byte, error) {
	res, err := http.Get(requestUrl)
	if err != nil || res.StatusCode != 200 {
		log.Printf("Unsuccessful GET request. Error %s", err.Error())
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}
