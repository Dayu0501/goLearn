package demo

/* 复查照片的回复 */
type listAttr struct {
	NewRectificationPeriod    string        `json:"newRectificationPeriod"`
	RecordType                interface{}   `json:"recordType"`
	IsDelete                  string        `json:"isDelete"`
	HiddenDangerID            string        `json:"hiddenDangerId"`
	NewRectificationTime      int64         `json:"newRectificationTime"`
	Autograph                 string        `json:"autograph"`
	UserName                  string        `json:"userName"`
	UserID                    interface{}   `json:"userId"`
	PicURL                    []string      `json:"picUrl"`
	CreatedAtL                string        `json:"createdAtL"`
	PicList                   []string `json:"picList"`
	Claim                     string        `json:"claim"`
	ReviewStatus              string        `json:"reviewStatus"`
	ID                        string        `json:"id"`
	PdfName                   string        `json:"pdfName"`
	UpdatedAtL                string        `json:"updatedAtL"`
	RectificationRequirements interface{}   `json:"rectificationRequirements"`
}

type dataAttr struct {
	Size int        `json:"size"`
	List []listAttr `json:"list"`
}

type ReviewData struct {
	Status  string   `json:"status"`
	Msg     string   `json:"msg"`
	Data    dataAttr `json:"data"`
	Comment string   `json:"comment"`
}
