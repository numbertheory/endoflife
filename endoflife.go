package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"flag"
)

type Product struct {
	Eol               string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	ReleaseDate       string `json:"releaseDate"`
	Lts               bool   `json:"lts"`
}

func getProductCycle(productName string, cycle string) {
	var url = "https://endoflife.date/api/" + productName + "/" + cycle + ".json"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    var product Product
    json.Unmarshal(bodyBytes, &product)
    fmt.Printf(product.Latest + "\n")
}

func main() {
	var product string    
    var cycle string 
	flag.StringVar(&product, "p", "python", "Specify product name. Default is python.")
    flag.StringVar(&cycle, "c", "3.9", "Specify cycle. Default is 3.9")
    flag.Parse()
 
	getProductCycle(product, cycle)
}