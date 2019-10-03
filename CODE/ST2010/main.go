package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var w map[string]int
var searchWords []string

func main() {
	const json02Path = "../../jsonout2/"
	//const outjpgPath = "../../OUTc/jpg"

	fmt.Println("================================")
	fmt.Println("==  beginning index of json=====")

	w = make(map[string]int)
	searchWords := CreateSearchWords()

	// // Open the file
	// csvfile, err := os.Open("./result2.csv")
	// if err != nil {
	// 	log.Fatalln("Couldn't open the csv file", err)
	// }

	// r := csv.NewReader(csvfile)
	// //r := csv.NewReader(bufio.NewReader(csvfile))

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("word: %s \n", record[0])
	// 	searchWords = append(searchWords, record[0])
	// }
	// //  =======  print out searchwords to confirm its inflated

	for i := 0; i < len(searchWords); i++ {
		fmt.Println(searchWords[i])
	}

	os.Exit(3)
	//========================================

	//doreadpages()

	// fn := json02Path + "mr0176.json"
	// fmt.Println(fn)

	// readpage(json02Path + "mr0175.json")
	// readpage(json02Path + "mr0176.json")
	// // readpage(json02Path + "mr0177.json")
	// fmt.Println("========   mr0178.json===========")
	// readpage(json02Path + "mr0178.json")
	// fmt.Println("========   mr0178.json===========")
	// fmt.Println("========   mr0179.json===========")
	// readpage(json02Path + "mr0179.json")
	// fmt.Println("===========  mr0179.json  ==========")

	// readpage(json02Path + "mr0180.json")
	// fmt.Println("===========  mr0179.json  ==========")

	// // Open our jsonFile
	// jsonFile, err := os.Open(fn)
	// // if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfully Opened users.json")
	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// // we initialize our Users array
	// // var users Users
	// var page PageJSON

	// json.Unmarshal(byteValue, &page)

	// sortedstrings := []string{}
	// for i := 0; i < len(page.PARAGRAPHS[0].WORDS); i++ {

	// 	bbddy := page.PARAGRAPHS[0].WORDS[i].BB.DDY
	// 	if bbddy > 140 {
	// 		// remove the header words from the screen
	// 		wd := page.PARAGRAPHS[0].WORDS[i].TEXT
	// 		fmt.Println(strconv.Itoa(i) + " - word: " + wd)
	// 		sortedstrings = insert(sortedstrings, wd)
	// 	}
	// }

	// fmt.Println(sortedstrings)
	// fmt.Println("===========================================")
	// //  convert this to a map with string count
	// w = make(map[string]int)

	// for i := 0; i < len(sortedstrings); i++ {
	// 	w[sortedstrings[i]]++
	// 	fmt.Println(strconv.Itoa(i) + "  " + sortedstrings[i] + "  -  " + strconv.Itoa(w[sortedstrings[i]]))
	// }
	// fmt.Println("===========================================")
	// for key, value := range w {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	fmt.Println("======================================")
	fmt.Println("======================================")
	fmt.Println("  w len = " + strconv.Itoa(len(w)))
	fmt.Println("======================================")
	fmt.Println("======================================")
	for key, value := range w {
		fmt.Println("Key:", key, "Value:", value)
	}

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

//==================================================
//==================================================
func CreateSearchWords() []string {
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
	//  =======  print out searchwords to confirm its inflated

	// for i := 0; i < len(sW); i++ {
	// 	fmt.Println(sW[i])
	// }
	return sW
}

//==================================================
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
