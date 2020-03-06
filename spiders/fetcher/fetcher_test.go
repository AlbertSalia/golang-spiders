package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	url := "https://www.lagou.com/jobs/allCity.html"

	bytes, err := Fetch(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
