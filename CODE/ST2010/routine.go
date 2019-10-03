package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//==============================================
func readpage(fn string) {

	const json02Path = "../../jsonout2/"
	//const outjpgPath = "../../OUTc/jpg"
	fmt.Println("================================")
	fmt.Println("=====  " + fn + "    =====")

	fn = json02Path + fn

	// Open our jsonFile
	jsonFile, err := os.Open(fn)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	// var users Users
	var page PageJSON

	json.Unmarshal(byteValue, &page)

	// sortedstrings := []string{}
	for i := 0; i < len(page.PARAGRAPHS[0].WORDS); i++ {

		bbddy := page.PARAGRAPHS[0].WORDS[i].BB.DDY
		bbaay := page.PARAGRAPHS[0].WORDS[i].BB.AAY
		if bbaay < 1990 && bbddy > 140 {
			// remove the header words from the screen
			wd := page.PARAGRAPHS[0].WORDS[i].TEXT
			// strings.ToLower(string(wd))
			strings.ToLower(wd)
			wd = strings.ToLower(string(wd))

			fmt.Println(strconv.Itoa(i) + " - word: " + wd)
			sortedstrings = insert(sortedstrings, wd)
		}
	}

	fmt.Println(sortedstrings)
	fmt.Println("===========================================")
	//  convert this to a map with string count

	for i := 0; i < len(sortedstrings); i++ {
		w[sortedstrings[i]]++
		fmt.Println(strconv.Itoa(i) + "  " + sortedstrings[i] + "  -  " + strconv.Itoa(w[sortedstrings[i]]))
	}
	fmt.Println("===========================================")
	for key, value := range w {
		fmt.Println("Key:", key, "Value:", value)
	}

}

//=============================================================
//=============================================================
//=============================================================
func doreadpages() {

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

	for i := 0; i < len(pp); i++ {
		fmt.Println(strconv.Itoa(i) + " ---- " + pp[i])
		readpage(pp[i])
	}
}

//==============================================
func lookupword1(s string) bool {

	i := sort.Search(len(searchWords), func(i int) bool { return s <= searchWords[i] })

	if i < len(searchWords) && searchWords[i] == s {
		//fmt.Printf("Found %s at index %d in %v.\n", s, i, searchWords)
		return true
	} else {
		//fmt.Printf("Did not find %s in %v.\n", s, searchWords)
		return false
	}

}

//==============================================
func lookupword2(s string) bool {
	// search list of words to track

	i := sort.Search(len(searchWords), func(i int) bool { return s <= searchWords[i] })

	if i < len(searchWords) && searchWords[i] == s {
		// fmt.Printf("Found %s at index %d in %v.\n", s, i, searchWords)
		return true
	} else {
		// fmt.Printf("Did not find %s in %v.\n", s, searchWords)
		return false
	}

}

//==============================================================
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//==============================================================
func insert(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}
