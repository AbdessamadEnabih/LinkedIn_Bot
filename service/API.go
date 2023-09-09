package API

import (
	"fmt"
	"io"
	"net/http"
)

func GetData() {

	resp, err := http.Get("https://api.publicapis.org/entries")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response status", resp.Status)

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		bodyString := string(bodyBytes)

		fmt.Println(bodyString)

	}

}
