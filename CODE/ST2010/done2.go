package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//=============================================================
//=============================================================
//=============================================================
//==========================================================
// step 02 - create slice of pages to be digested.
func doreadpages() []string {
	var pp []string
	// Open the file
	csvfile, err := os.Open("./mr.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("word: %s \n", record[0])
		pp = append(pp, record[0])
	}
	end := len(pp) - 1
	return pp[:end]
}
