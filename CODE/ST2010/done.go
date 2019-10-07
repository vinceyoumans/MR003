package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

//=============================================================
//=============================================================
//=============================================================
//---  Step 01
//  these are the words to be tracked.
//============================================================
func createSearchWords01() []string {
	var sW []string
	// Open the file
	csvfile, err := os.Open("./WSearch001.csv")
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
		sW = append(sW, record[0])
	}
	return sW
}

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

//=============================================================
//=============================================================
//=============================================================
//---  Step 03
func step03ListWords() []string {
	// x := []string{}
	var tpageswords01 []string
	var tpageswords02 []string

	for i := 0; i < len(jsonPages01); i++ {
		fn := jsonPages01[i]
		// fmt.Println("=======================================\n")
		fmt.Printf("===i = %d  -  fn = %s\n", i, fn)

		tpageswords01 = readpage(fn)
		fmt.Println("======    returning from readpage  ========")
		// append words to master list
		fmt.Println(tpageswords01)
		for ii := 0; ii < len(tpageswords01); ii++ {
			fmt.Println(ii)
			tpageswords02 = append(tpageswords02, tpageswords01[ii])
		}
		sort.Strings(tpageswords02)
		tpageswords02 = removeDuplicatesUnordered(tpageswords02)

	}
	return tpageswords02
}

//=============================================================
//=============================================================
//=============================================================
//---  Step 04
func readpage(fn string) []string {
	var page PageJSON
	const json02Path = "../../jsonout2/"
	var tempWD01 []string
	var tempWD02 []string

	fn = json02Path + fn
	jsonFile, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &page)

	for i := 0; i < len(page.PARAGRAPHS[0].WORDS); i++ {
		//  Remove the Header and Footer words.
		// 	set all words to lower case
		bbddy := page.PARAGRAPHS[0].WORDS[i].BB.DDY
		bbaay := page.PARAGRAPHS[0].WORDS[i].BB.AAY
		if bbaay < 1990 && bbddy > 140 {
			wd := page.PARAGRAPHS[0].WORDS[i].TEXT
			wd = strings.ToLower(string(wd))
			tempWD01 = append(tempWD01, wd) // unsorted list
		}
	}
	sort.Strings(tempWD01)

	tempWD02 = removeDuplicatesUnordered(tempWD01)
	sort.Strings(tempWD02)
	for i := 0; i < len(tempWD02); i++ {
		fmt.Printf("---  i = %d  temp02 = %s \n", i, tempWD02[i])
	}

	fmt.Println(" ---------  exiting readpage  ------")
	return tempWD02
}

//=========
func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

//=============================================================
//=============================================================
//=============================================================
//---  Step 05
//=========================================
// This does a word dump to csv
// which can be edited.
// Just keep the words that you want to track

func step05CreateBigWordList(w []string) {
	file, err := os.Create("./big_result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range step04Words {
		v := []string{value}
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}
	//os.Exit(3) // Exit here for testing
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//=============================================================
//=============================================================
//=============================================================
//---  Step 06
//=========================================
func step06BuildBigStruct() []word {
	// read the CSV of all words
	var data word
	var datas []word
	csvFile, _ := os.Open("./big_result.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		data.Word = line[0]
		data.Track1 = true //tracking every word
		data.Track2 = false
		data.Count = 0
		datas = append(datas, data)
	}
	return datas
}

//=============================================================
//=============================================================
//=============================================================
//---  Step 07
//=========================================

func step7TagFilesToTrack() {

	// create mape of Search Strings
	// refactor.. SearchWords should have been a MAP from start.
	// wp := make(map[string]struct{}, len(searchWords))
	// wp := make(map[string]bool, len(searchWords))

	searchWords := createSearchWords01()
	//  This is to confirm SearchWords has been populated properly
	for i := 0; i < len(searchWords); i++ {
		fmt.Printf("searchWords -- %d -- %s\n", i, searchWords[i])
	}
	//os.Exit(3) // Exit here for testing

	wp := make(map[string]bool)
	for i := 0; i < len(searchWords); i++ {
		fmt.Printf("---word  %d  = %s \n", i, searchWords[i])
		wp[searchWords[i]] = true
	}

	fmt.Printf("=====  len of searchwords  %d\n", len(searchWords))
	fmt.Printf("=====  len of wp map  %d\n", len(wp))
	//os.Exit(3) // Exit here for testing

	for i := 0; i < len(step6WordsToTrack); i++ {
		w := step6WordsToTrack[i].Word
		_, isPresent := wp[w]

		// rs := sort.SearchStrings(searchWords, w)

		fmt.Printf("======== %t == %s \n", isPresent, w)

		if isPresent {
			// found the word
			step6WordsToTrack[i].Track2 = true
		} else {
			// not found
			step6WordsToTrack[i].Track2 = false
		}

	}

}

//=============================================================
//=============================================================
//=============================================================
//---  Step 10
//=========================================
