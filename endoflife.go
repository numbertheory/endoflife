package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Eol               string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	ReleaseDate       string `json:"releaseDate"`
	Lts               bool   `json:"lts"`
}

const usage = `Usage of endoflife:
  -p, --product [PRODUCT]
      specify the product to find the end of life date for
  -c, --cycle [VERSION] 
      specify the product cycle to show
  -h, --help print this help information and exit 
`

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
	flag.StringVar(&product, "p", "", "product name")
	flag.StringVar(&product, "product", "", "product name")
	flag.StringVar(&cycle, "c", "", "product cycle")
	flag.StringVar(&cycle, "cycle", "", "product cycle")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	getProductCycle(product, cycle)
}
