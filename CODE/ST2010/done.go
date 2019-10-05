package main

import (
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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//=============================================================
//=============================================================
//=============================================================
// func doreadpages() string {
func doreadpages() []string {

	//  create an array of files to scan
	pp := [...]string{
		"mr0001.json",
		"mr0002.json",
		"mr0003.json",
		"mr0004.json",
		"mr0005.json",
		"mr0006.json",
		"mr0007.json",
		"mr0008.json",
		"mr0009.json",
		"mr0010.json",
		"mr0011.json",
		"mr0012.json",
		"mr0013.json",
		"mr0014.json",
		"mr0015.json",
		"mr0016.json",
		"mr0017.json",
		"mr0018.json",
		"mr0019.json",
		"mr0020.json",
		"mr0021.json",
		"mr0022.json",
		"mr0023.json",
		"mr0024.json",
		"mr0025.json",
		"mr0026.json",
		"mr0027.json",
		"mr0028.json",
		"mr0029.json",
		"mr0030.json",
		"mr0031.json",
		"mr0032.json",
		"mr0033.json",
		"mr0034.json",
		"mr0035.json",
		"mr0036.json",
		"mr0037.json",
		"mr0038.json",
		"mr0039.json",
		"mr0040.json",
		"mr0050.json",
		"mr0060.json",
		"mr0070.json",
		"mr0080.json",
		"mr0090.json",
		"mr0100.json",
		"mr0110.json",
		"mr0120.json",
		"mr0130.json",
		"mr0140.json",
		"mr0150.json",
		"mr0160.json",
		"mr0170.json",
		"mr0171.json",
		"mr0172.json",
		"mr0173.json",
		"mr0174.json",
		"mr0175.json",
		"mr0176.json",
		// "mr0177.json",
		"mr0178.json",
		"mr0179.json",
		"mr0180.json",
		"mr0181.json",
		"mr0182.json",
		"mr0183.json",
	}

	// for i := 0; i < len(pp); i++ {
	// 	fmt.Println(strconv.Itoa(i) + " ---- " + pp[i])
	// 	readpage(pp[i])
	// }

	end := len(pp) - 1
	return pp[:end]
}

//==================================================
//==================================================
func createSearchWords01() []string {
	var sW []string
	// Open the file
	csvfile, err := os.Open("./result2.csv")
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
