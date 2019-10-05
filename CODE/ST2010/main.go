package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var w map[string]int
var searchWords []string
var jsonPages01 []string
var step03Words []string
var step04Words []string

func main() {
	const json02Path = "../../jsonout2/"
	//const outjpgPath = "../../OUTc/jpg"

	fmt.Println("================================")
	fmt.Println("==  beginning index of json=====")

	w = make(map[string]int)

	//==========================================================
	//  step 1 - populate search words
	//  all of the words that should be tagged for searching.
	//  This model can be used to create several index strategies.
	searchWords := createSearchWords01()
	//  This is to confirm SearchWords has been populated properly
	for i := 0; i < len(searchWords); i++ {
		fmt.Println(searchWords[i])
	}
	// os.Exit(3)  // Exit here for testing
	//========================================

	//==========================================================
	// step 02 - create slice of pages to be digested.
	// not important what page they are listed.
	jsonPages01 = doreadpages()
	for i := 0; i < len(jsonPages01); i++ {
		fmt.Println("----- jsonPages01   ---" + jsonPages01[i])
	}
	//Expected result.  Slice of strings with filenames to be inspected.
	//os.Exit(3) // Exit here for testing
	//========================================

	//==========================================================
	// step 03 - create slice of pages to be digested.
	// not important what page they are listed.
	step03Words = step03ListWords()
	for i := 0; i < len(step03Words); i++ {
		fmt.Printf("--- step03Words - %d - %s \n", i, step03Words[i])
	}
	//Expected result.  Slice of strings with filenames to be inspected.
	// os.Exit(3) // Exit here for testing
	//========================================

	//==========================================================
	// step 04 - Sort then remove duplicates
	step04Words = removeDuplicatesUnordered(step03Words)
	sort.Strings(step04Words)
	for i := 0; i < len(step04Words); i++ {
		fmt.Printf("--- step04 - %d - %s \n", i, step04Words[i])
	}
	os.Exit(3) // Exit here for testing

	//==========================================================
	// step 05 - Dump list to csv file
	//  this file can be edited to include only tracked words.

	fmt.Println("======================================")
	fmt.Println("=====  Sorted          ===============")
	fmt.Println("  w len = " + strconv.Itoa(len(w)))
	fmt.Println("======================================")
	fmt.Println("======================================")

	keys := make([]string, 0, len(w))
	for k := range w {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var fw word
	var fww []word

	for _, k := range keys {
		fmt.Println(k, w[k])
		fw.Track1 = false
		fw.Track2 = false
		fw.Count = w[k]
		fw.Word = k
		fww = append(fww, fw)
	}

	//==============================
	// confirm that fww is saved properly
	for i := 0; i < len(fww); i++ {
		fmt.Println(fww[i])
	}

	//==============================
	// tag words to be tracked

	for i := 0; i < len(fww); i++ {
		// fww[i].Track1 = lookupword1(fww[i].Word)
		// fww[i].Track2 = lookupword2(fww[i].Word)
	}

	//==============================
	// save json to a file

	jsonData, err := json.Marshal(fww)
	jsonFile, err := os.Create("./Result01.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())

	//==============================
	// save to CDF of words, to be added to search list

	var data = []string{}
	for i := 0; i < len(fww); i++ {
		data = append(data, fww[i].Word)
		fmt.Println(data[i])
	}

	file, err := os.Create("./result.csv")
	checkError("Cannot create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {

		fmt.Println("=====================")
		fmt.Println(value)
		v := []string{value}
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}

}
