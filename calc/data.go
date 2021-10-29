package calc

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func webScrape() (rateStrings []string) {
	res, err := http.Get("http://www.yahii.com.br/incc.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("Status code error: %d %s\n", res.StatusCode, res.Status)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("body > div > center > center > center:nth-child(15) > center > center > center > font:nth-child(5) > center > center > div > center > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		entry := s.Find("font").Text()
		rateStrings = append(rateStrings, entry)
	})
	return
}

func formatStrings(rateStrings []string) (fmtString string) {
	rateStrings = rateStrings[14:]
	rateStrings = rateStrings[:len(rateStrings)-14]

	for i, entry := range rateStrings {
		entry = strings.TrimSpace(entry)
		rateStrings[i] = entry
		if (i%14 == 0) || (i%14 == 13) || (len(entry) == 0) {
			continue
		} else {
			entry = entry[:len(entry)-1]
			entry = strings.ReplaceAll(entry, ",", ".")
			if entry[0] == '(' {
				entry = "-" + entry[3:]
			}
			fmtString = fmtString + entry + " "
		}
	}
	return
}

func saveToFile(s string, date time.Time) {

	filePath := filepath.Join("calc", "data", "rates.txt")
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	dateString := date.Format("02012006")

	s = dateString + " " + s

	_, err = file.WriteString(s)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated rates!")

}

func IsDateNew(newDate time.Time) bool {

	filePath := filepath.Join("calc", "data", "rates.txt")

	file, err := os.Open(filePath)

	if os.IsNotExist(err) {
		return true
	} else if err != nil {
		log.Fatal(err)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileString := string(fileContent)

	oldDate, err := time.Parse("02012006", fileString[:8])

	if err != nil {
		log.Fatal(err)
	}

	return newDate.After(oldDate)
}

func UpdateRates(date time.Time) {

	if date.After(time.Now()) {
		date = time.Now()
	}

	rateStrings := webScrape()
	stringToSave := formatStrings(rateStrings)
	saveToFile(stringToSave, date)

}

func ReadTable(startDate time.Time, endDate time.Time) (totalRate float64) {

	filePath := filepath.Join("calc", "data", "rates.txt")

	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileString := string(fileContent)

	packedRates := strings.Fields(fileString)

	epochYear := 1990

	startDay := startDate.Day()
	startMonth := int(startDate.Month())
	startYear := startDate.Year()

	endMonth := int(endDate.Month())
	endYear := endDate.Year()

	epochToStart := (startYear - epochYear) * 12

	if startDay <= 15 {
		epochToStart += startMonth
	} else {
		epochToStart += startMonth + 1
	}

	epochToEnd := (endYear - epochYear) * 12

	epochToEnd += endMonth

	calcRates := packedRates[epochToStart:epochToEnd]

	totalRate = 1

	for _, rateString := range calcRates {

		rateFloat, err := strconv.ParseFloat(rateString, 64)

		if err != nil {
			log.Fatal(err)
		}

		rateFloat = 1 + rateFloat/100

		totalRate *= rateFloat
	}

	return

}
