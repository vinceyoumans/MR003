package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

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

//==============================================
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

//==================================================

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
