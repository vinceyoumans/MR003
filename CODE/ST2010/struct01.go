package main

type index001 struct {
	Fname string `json:"fname"`
}

type word struct {
	Track1 bool   `json:"track1"`
	Track2 bool   `json:"track2"`
	Word   string `json:"word"`
	Count  int    `json:"count"`
}

//============================================================
type PageJSON struct {
	PAGENUMBER        int    `json:"PAGE_NUMBER"`
	FileNameLarge     string `json:"FileName_Large"`
	FileNameThumb     string `json:"FileName_Thumb"`
	PageWidth         int    `json:"PageWidth"`
	PageHeight        int    `json:"PageHeight"`
	DocumentText      string `json:"Document_Text"`
	DocumentTextFixed string `json:"Document_Text_Fixed"`
	PARAGRAPHS        []struct {
		TEXT      string `json:"TEXT"`
		TEXTFixed string `json:"TEXT_fixed"`
		WORDS     []struct {
			TEXT string `json:"TEXT"`
			BB   struct {
				AAX int `json:"AAX"`
				AAY int `json:"AAY"`
				BBX int `json:"BBX"`
				BBY int `json:"BBY"`
				CCX int `json:"CCX"`
				CCY int `json:"CCY"`
				DDX int `json:"DDX"`
				DDY int `json:"DDY"`
			} `json:"BB"`
		} `json:"WORDS"`
		BB struct {
			AAX int `json:"AAX"`
			AAY int `json:"AAY"`
			BBX int `json:"BBX"`
			BBY int `json:"BBY"`
			CCX int `json:"CCX"`
			CCY int `json:"CCY"`
			DDX int `json:"DDX"`
			DDY int `json:"DDY"`
		} `json:"BB"`
	} `json:"PARAGRAPHS"`
	WC interface{} `json:"WC"`
}

type UUID [16]byte

type MRPages struct {
	MRPage []MRPage `json:"mrpage"` // slice of MRPage
}
type MRPage struct {
	Index    int    `json:"index"`     // 1 to x page number incremented
	PNum     int    `json:"pnum"`      // integer file name
	PageNum  string `json:"page_num"`  // pdf labeled page number
	FileName string `json:"file_name"` // file name of pdf
	Text     string `json:"text"`      // the raw text of the page
}
