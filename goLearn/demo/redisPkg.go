package demo

/* request body struct */
type PhotoReqAttr struct {
	MainPhoto     string `json:"mainPhoto"`
	MainPhotoPath string `json:"mainPhotoPath"`
	SubPhoto      string `json:"subPhoto"`
	SubPhotoPath  string `json:"subPhotoPath"`
}

type OtherAlgTaskReqAttr struct {
	InSystemTime  string `json:"inSystemTime"`
	ReponseChan   string `json:"reponseChan"`
	TimeoutSecond string `json:"timeoutSecond"`
}

type AlgTaskListReqAttr struct {
	AlgIDs         []int               `json:"algIDs"`
	AlgTaskKeyCode string              `json:"algTaskKeyCode"`
	EncodeKey      string              `json:"encodeKey"`
	Other          OtherAlgTaskReqAttr `json:"other"`
	Param          interface{}         `json:"param"`
	Photo          PhotoReqAttr        `json:"photo"`
}

type AlgTaskInfoReqPkg struct {
	AlgTasks       []*AlgTaskListReqAttr `json:"algTaskList"`
	EncodeKey      string                `json:"encodeKey"`
	InSystemTime   string                `json:"inSystemTime"`
	PicProcessTime string                `json:"picProcessTime"`
	ResponseChan    string                `json:"reponseChan"`
	SessionID      string                `json:"sessionId"`
	TimeoutSecond  string                `json:"timeoutSecond"`
}

/* response body struct */
type AlgReturnRespAttr struct {
	Num7002 string `json:"7002"`
	Status string `json:"status"`
}

type Num7002RespAttr struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type ReturnResultMemberListRespAttr struct {
	Num7002 []Num7002RespAttr `json:"7002"`
}

type PhotoSavePathRespAttr struct {
	MainPhoto string `json:"mainPhoto"`
	SubPhoto string `json:"subPhoto"`
	OtherPhoto []string `json:"otherPhoto"`
}

type OtherRespAttr struct {
	InSystemTime string `json:"inSystemTime"`
	ReponseChan string `json:"reponseChan"`
	TimeoutSecond string `json:"timeoutSecond"`
	AlgObjTime float64 `json:"AlgObjTime"`
	AlgProcTime float64 `json:"AlgProcTime"`
	AlgPicParseTime float64 `json:"AlgPicParseTime"`
	OutSystemTime string `json:"OutSystemTime"`
	AlgIP string `json:"AlgIp"`
	AlgVersion string `json:"AlgVersion"`
	PicDownloadTimeOut string `json:"PicDownloadTimeOut"`
}

type AlgTaskInfoRespPkg struct {
	AlgTaskKeyCode string `json:"algTaskKeyCode"`
	AlgReturn AlgReturnRespAttr `json:"algReturn"`
	ReturnResultMemberList []ReturnResultMemberListRespAttr `json:"returnResultMemberList"`
	PhotoSavePath PhotoSavePathRespAttr `json:"photoSavePath"`
	Other OtherRespAttr `json:"other"`
}