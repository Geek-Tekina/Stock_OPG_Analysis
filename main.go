package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

type Stock struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

// this is Load function which accepts the string for csv path, return slice of stock objects and error. It opens the files and performs the read operation.
func Load(path string) ([]Stock, error) {
	// this method is used to open the file before doing any I/O operation on it
	f, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	// fmt.Print(f)
	fmt.Println("File opened Successfully !!")
	r := csv.NewReader(f)
	// fmt.Print(r)
	rows, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("File read successfully by csv reader !!")
	// Now we don't need the header row i.e. the first row
	rows = slices.Delete(rows, 0, 1) // delete from rows starting from 0 index and delete 1 row
	fmt.Println("Rows after deleting the header row :")

	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println("Now the rows will go into Stock struct as single values")

	var stocks []Stock // a slice of Stock structure

	for _, row := range rows {
		ticker := row[0]
		// gap := row[1] this gives error
		gap, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue // means if one row has some error in gap field we continue to other row
		}
		// openingPrice := row[2]
		openingPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		stocks = append(stocks, Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}

	return stocks, nil
}

func main() {
	stocks, err := Load("./opg.csv")
	if err != nil {
		fmt.Println("Erros in reading the file", err)
		return
	}
	fmt.Println("The CSV  file is now a slice of Stock Objects.")
	for index, stockElement := range stocks {
		fmt.Printf("Stock %v : %v \n", index+1, stockElement)
	}

	fmt.Println("Now the program will filter out the unworthy stocks as per the condition that Gap >= 0.1")

	stocks = slices.DeleteFunc(stocks, func(s Stock) bool {
		if math.Abs(s.Gap) < .1 {
			return true
		} else {
			return false
		}
	})

	fmt.Println("The list of worthy stocks is :", stocks)
}
