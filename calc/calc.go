package calc

import (
	"errors"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"time"
)

func readDate(dateString string) (dateTime time.Time, err error) {

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	dateString = reg.ReplaceAllString(dateString, "")

	if len(dateString) != 8 {
		err = errors.New("date must be 8 numbers long")
		return
	}

	dateTime, err = time.Parse("02012006", dateString)

	return
}

// 'calc [value] [startDate] [endDate]'

func ComCalc(packedCommand []string) {

	if len(packedCommand) < 3 {
		fmt.Println("Usage: 'calc [value] [startDate] (endDate)'")
		return
	}

	value, err := strconv.ParseFloat(packedCommand[1], 64)
	if err != nil {
		fmt.Println("[value] must be a valid number!")
		return
	}

	startDate, err := readDate(packedCommand[2])
	if err != nil {
		fmt.Println("[startDate] must be a valid date!")
		return
	}

	endDate := time.Now()

	if len(packedCommand) == 4 {
		endDate, err = readDate(packedCommand[3])
		if err != nil {
			fmt.Println("[endDate] must be a valid date!")
			return
		}
	}

	if IsDateNew(endDate) {
		UpdateRates(endDate)
	}

	totalRate := ReadTable(startDate, endDate)

	value *= totalRate

	value = math.Round(value*100) / 100

	fmt.Printf("%.2f\n", value)

}
