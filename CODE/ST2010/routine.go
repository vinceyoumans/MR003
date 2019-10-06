package main

import (
	"fmt"
	"sort"
)

func step7TagFilesToTrack() {

	// create mape of Search Strings
	// refactor.. SearchWords should have been a MAP from start.
	// wp := make(map[string]struct{}, len(searchWords))
	wp := make(map[string]bool, len(searchWords))
	for i := 0; i < len(searchWords); i++ {
		wp[searchWords[i]] = true
	}

	for i := 0; i < len(step6WordsToTrack); i++ {

		w := step6WordsToTrack[i].Word

		rs := sort.SearchStrings(searchWords, w)

		fmt.Printf("======== %d == %s \n", rs, w)

		if rs < len(searchWords) {
			// found the word
			step6WordsToTrack[i].Track2 = true
		} else {
			// not found
			step6WordsToTrack[i].Track2 = false
		}

	}

}
