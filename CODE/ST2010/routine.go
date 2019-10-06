package main

import (
	"fmt"
)

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
