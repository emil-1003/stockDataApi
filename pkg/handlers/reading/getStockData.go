package reading

import (
	"encoding/json"
	"github.com/emilstorgaardandersen/stockDataApi/pkg/models"
	"github.com/gocolly/colly"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		symbol := params["id"]

		stock := stockInfo(symbol)
		json.NewEncoder(w).Encode(stock)
	}
}

func GetMyPortfolio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var stockdata []models.StockData

		for _, item := range []string{"QFUEL.OL","LUN.CO","SPVIGAKL.CO","SPIDJWKL.CO","DKIDKIX.CO","SPIC25KL.CO","SPIEMIKL.CO"} {
			stockdata = append(stockdata, stockInfo(item))
		}

		json.NewEncoder(w).Encode(stockdata)
	}
}

func GetMultiData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var stockdata []models.StockData

		for _, item := range r.URL.Query()["stocks"] {
			stockdata = append(stockdata, stockInfo(item))
		}

		json.NewEncoder(w).Encode(stockdata)
	}
}

func findAllTagText(element *colly.HTMLElement, tag string) []string {
	var data []string

	element.ForEach(tag, func(_ int, el *colly.HTMLElement) {
		data = append(data, el.Text)
	})

	return data
}

func stockInfo(symbol string) models.StockData {
	var data []string
	var test models.StockData
	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s",symbol)

	fmt.Println("Getting: " + symbol)

	collector := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com", "www.finance.yahoo.com"),
	)

	collector.OnHTML("div[class='D(ib) Mend(20px)']", func(element *colly.HTMLElement){

		data = findAllTagText(element, "span")
		test = models.StockData {
			Symbol: symbol,
			Price:  data[0],
			Change: data[1],
		}

	})

	collector.Visit(url)

	return test
}
