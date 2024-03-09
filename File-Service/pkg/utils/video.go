package utils

import (
	"encoding/json"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/File-Service/config"
	"io/ioutil"
	"log"
	"net/http"
)

type Video struct {
	Duration float64 `json:"duration"`
}

func GetDurationVideo(cf *config.Config, urlVideo string) (*Video, error) {
	var video Video
	url := "http://" + cf.OtherServices.DurationVideo + "/duration?url=" + urlVideo
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(bodyBytes, &video)
	return &video, nil
}
