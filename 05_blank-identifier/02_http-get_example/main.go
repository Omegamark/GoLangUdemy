package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	res, err := http.Get("http://www.mcleods.com/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	page, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", page)
// }

// The function below is the same as the one above with 'blank identifiers' and no error checking.

func main() {
	res, _ := http.Get("https://www.giantbomb.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
