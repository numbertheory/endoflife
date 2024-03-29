package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"github.com/fatih/color"
	"time"
	"strconv"
)

type ProductCycle struct {
	Eol               string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	ReleaseDate       string `json:"releaseDate"`
	Lts               bool   `json:"lts"`
}

type Product struct {
	Eol               []ProductCycle `json:"eol"`
	Latest            []ProductCycle `json:"latest"`
	LatestReleaseDate []ProductCycle `json:"latestReleaseDate"`
	ReleaseDate       []ProductCycle `json:"releaseDate"`
	Lts               []ProductCycle `json:"lts"`
}


const usage = `Usage of endoflife:
  -p, --product [PRODUCT]
      specify the product to find information about
  -c, --cycle [VERSION] 
      specify the product cycle to show
  -h, --help print this help information and exit 
`

func isPast(eolDate string) (string) {
	currentTime := time.Date(
		time.Now().Year(),
		time.Now().Month(), 
		time.Now().Day(),
		0, 
		0, 
		0,
		100,
		time.Local)
	
	eolYear, _ := strconv.Atoi(strings.Split(eolDate, "-")[0])
	eolMonth, _ := strconv.Atoi(strings.Split(eolDate, "-")[1])
	eolDay, _ := strconv.Atoi(strings.Split(eolDate, "-")[2])

	eolTime := time.Date(
		eolYear,
		time.Month(eolMonth),
		eolDay,
		0, 
		0, 
		0,
		0,
		time.Local)
	
	diffTime := eolTime.Sub(currentTime)
	if (diffTime < 0) {
		red := color.New(color.FgRed).SprintFunc()
		return red(eolDate)
	} else {
		return eolDate
	}
}


func getProductCycle(productName string, cycle string) {
	var url = "https://endoflife.date/api/" + productName + "/" + cycle + ".json"
	if (cycle == "") {
		url = "https://endoflife.date/api/" + productName + ".json"
	}
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if (resp.StatusCode == 404) {
		fmt.Printf("Product or cycle does not exist\n")
		return 
	}

	if (cycle == "") { 
		product := make([]ProductCycle, 0)
		yellow := color.New(color.FgYellow).SprintFunc()
		json.Unmarshal([]byte(bodyBytes), &product)
		fmt.Printf("|    EOL     |   Latest Version  |\n")
		fmt.Printf("+============+===================+\n")
		for i := 0; i < len(product); i++ {
			buffer := strings.Repeat(" ", 18 - len(product[i].Latest))
			fmt.Printf("| " + isPast(product[i].Eol) + " | " + yellow(product[i].Latest) + buffer + "|\n")
		}
		fmt.Printf("+============+===================+\n")
	} else {
		var productCycle ProductCycle
		json.Unmarshal(bodyBytes, &productCycle)
		fmt.Printf("|    EOL     |   Latest Version  |\n")
		fmt.Printf("+============+===================+\n")
		buffer := strings.Repeat(" ", 18 - len(productCycle.Latest))
		fmt.Printf("| " + isPast(productCycle.Eol) + " | " + productCycle.Latest + buffer + "|\n")
		fmt.Printf("+============+===================+\n")
	}
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
	if product != "" && cycle != "" {
		getProductCycle(product, cycle)	
	} else if (product != "") {
		getProductCycle(product, "")	
	} else {
		fmt.Print(usage)
	}
}
