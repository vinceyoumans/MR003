package main

import (
	"sort"
)

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
// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

//==============================================================
// func insert(ss []string, s string) []string {
// 	i := sort.SearchStrings(ss, s)
// 	ss = append(ss, "")
// 	copy(ss[i+1:], ss[i:])
// 	ss[i] = s
// 	return ss
// }
