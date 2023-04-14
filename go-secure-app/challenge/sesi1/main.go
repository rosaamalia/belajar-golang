package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func checkStatus(data Data) (string, string){
	var water string
	var wind string

	if (data.Water <= 5) {
		water = "aman"
	} else if (data.Water <= 8 && data.Water >= 6) {
		water = "siaga"
	} else if (data.Water > 8) {
		water = "bahaya"
	}

	if (data.Wind <= 15 && data.Wind >= 7) {
		wind = "siaga"
	} else if (data.Wind > 15) {
		wind = "bahaya"
	} else if (data.Wind < 7) {
		wind = "aman"
	}
	
	return water, wind
}

func main() {
	for {
		// Membuat data dengan nilai random untuk water dan wind
		data := Data{Water: rand.Intn(100) + 1, Wind: rand.Intn(100) + 1}

		// Mengirim data dalam format JSON ke URL yang diberikan
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewReader(jsonData))
		if err != nil {
			fmt.Println(err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		// Mengirim HTTP Request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer resp.Body.Close()

		// Membaca response dari server
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Unmarshal data JSON ke struct Data
		var result Data
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
			return
		}

		waterStatus, windStatus := checkStatus(result)
		fmt.Println(string(body))
		fmt.Println("Status water:", waterStatus)
		fmt.Println("Status wind:", windStatus)

		// Menunggu selama 15 detik sebelum mengirim data kembali
		time.Sleep(15 * time.Second)
	}
}
