package main

import (
	"fmt"
	"sort"
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
	//os.Exit(3) // Exit here for testing

	//==========================================================
	// step 05 - Dump list to csv file
	//  this file can be edited to include only tracked words.
	step05CreateBigWordList(step04Words)

	//os.Exit(3) // Exit here for testing

}
