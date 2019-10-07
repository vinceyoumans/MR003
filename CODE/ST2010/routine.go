package main

import (
	"strconv"
	"strings"
)

//=============================================================
//=============================================================
//=============================================================
//==========================================================
// step 10 - create slice of pages to be digested.
// creates json of pages
func Step10_GenBigScreen01() []MRPage {

	var tp MRPage
	var tps []MRPage

	pp := 1
	ppp := 1
	// jsonPages01
	// for i := 0; i < len(jsonPages01); i++ {
	for index, MRfilename := range jsonPages01 {

		tp.FileName = jsonPages01[index]
		tp.PNum = IDNum(jsonPages01[index])
		tp.Index = index
		i := IDNum(MRfilename)

		if i == 1 {
			tp.PageNum = "title - Volume 1 of 2"
		} else if i == 2 {
			tp.PageNum = "Blank"
		} else if i == 3 {
			tp.PageNum = "V1 - i"
		} else if i == 4 {
			tp.PageNum = "V1 - ii"
		} else if i == 5 {
			tp.PageNum = "V1 - iii"
		} else if i == 6 {
			tp.PageNum = "V1 - iv"
		} else if i == 7 {
			tp.PageNum = "V1 - v"
		} else if i == 8 {
			tp.PageNum = "Blank"
		} else if i >= 9 && i <= 207 {
			tp.PageNum = "V1 - " + strconv.Itoa(pp)
			pp++
		} else if i == 208 {
			tp.PageNum = "Title"
		} else if i == 209 {
			tp.PageNum = "V2 - i"
		} else if i == 210 {
			tp.PageNum = "V2 - ii"
		} else if i == 211 {
			tp.PageNum = "V2 - iii"
		} else if i == 212 {
			tp.PageNum = "V2 - iv"
		} else if i >= 213 && i <= 394 {
			tp.PageNum = "V2 - " + strconv.Itoa(ppp)
			ppp++
			// 394 = p182
		} else if i == 395 {
			tp.PageNum = "Appendix A"
		} else if i == 396 {
			tp.PageNum = "Blank"
		} else if i == 397 {
			tp.PageNum = "A-1"
		} else if i == 398 {
			tp.PageNum = "Blank"
		} else if i == 399 {
			tp.PageNum = "Appendix B"
		} else if i == 400 {
			tp.PageNum = "Blank"
		} else if i == 401 {
			tp.PageNum = "B-1"
		} else if i == 402 {
			tp.PageNum = "B-2"
		} else if i == 403 {
			tp.PageNum = "B-3"
		} else if i == 404 {
			tp.PageNum = "B-4"
		} else if i == 405 {
			tp.PageNum = "B-5"
		} else if i == 406 {
			tp.PageNum = "B-6"
		} else if i == 407 {
			tp.PageNum = "B-7"
		} else if i == 408 {
			tp.PageNum = "B-8"
		} else if i == 409 {
			tp.PageNum = "B-9"
		} else if i == 410 {
			tp.PageNum = "B-10"
		} else if i == 411 {
			tp.PageNum = "B-11"
		} else if i == 412 {
			tp.PageNum = "B-12"
		} else if i == 413 {
			tp.PageNum = "B-13"
		} else if i == 414 {
			tp.PageNum = "B-14"
		} else if i == 415 {
			tp.PageNum = "Blank"
		} else if i == 416 {
			tp.PageNum = "Appendix C"
		} else if i == 417 {
			tp.PageNum = "C-1"
		} else if i == 418 {
			tp.PageNum = "C-2"
		} else if i == 419 {
			tp.PageNum = "C-3"
		} else if i == 420 {
			tp.PageNum = "C-4"
		} else if i == 421 {
			tp.PageNum = "C-5"
		} else if i == 422 {
			tp.PageNum = "C-6"
		} else if i == 423 {
			tp.PageNum = "C-7"
		} else if i == 424 {
			tp.PageNum = "C-8"
		} else if i == 425 {
			tp.PageNum = "C-9"
		} else if i == 426 {
			tp.PageNum = "C-10"
		} else if i == 427 {
			tp.PageNum = "C-11"
		} else if i == 428 {
			tp.PageNum = "C-12"
		} else if i == 429 {
			tp.PageNum = "C-13"
		} else if i == 430 {
			tp.PageNum = "C-14"
		} else if i == 431 {
			tp.PageNum = "C-15"
		} else if i == 432 {
			tp.PageNum = "C-16"
		} else if i == 433 {
			tp.PageNum = "C-17"
		} else if i == 434 {
			tp.PageNum = "C-18"
		} else if i == 435 {
			tp.PageNum = "C-19"
		} else if i == 436 {
			tp.PageNum = "C-20"
		} else if i == 437 {
			tp.PageNum = "C-21"
		} else if i == 438 {
			tp.PageNum = "C-22"
		} else if i == 439 {
			tp.PageNum = "C-23"
		} else if i == 440 {
			tp.PageNum = "Blank"
		} else if i == 441 {
			tp.PageNum = "Appendix D"
		} else if i == 442 {
			tp.PageNum = "Blank"
		} else if i == 443 {
			tp.PageNum = "D-1"
		} else if i == 444 {
			tp.PageNum = "D-2"
		} else if i == 445 {
			tp.PageNum = "D-3"
		} else if i == 446 {
			tp.PageNum = "D-4"
		} else if i == 447 {
			tp.PageNum = "D-5"
		} else if i == 448 {
			tp.PageNum = "D-6"
		}
		tps = append(tps, tp)
	}
	return tps
}

// IDNum returns int value of file name
func IDNum(fn string) int {
	// fnn = strings.split(fn,".")
	fnn := strings.Replace(fn, ".jpg", "", -1)
	fnn = strings.Replace(fnn, ".json", "", -1)
	fnn = strings.Replace(fnn, "mr0", "", -1)
	ifn, _ := strconv.Atoi(fnn)
	return ifn

}
