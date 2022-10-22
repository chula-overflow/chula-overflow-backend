package dto

type ExamSummary struct {
	ExamId      string `json:"exam_id" example:"TODO!(MONGOOBJECTID)"`
	ExamName    string `json:"exam_name" example:"2011 - S1 - Final"`
	ThreadCount int32  `json:"thread_count" example:"12"`
}

type Exam struct {
	ExamName string          `json:"exam_name" example:"2011 - S1 - Final"`
	Threads  []ThreadSummary `json:"threads"`
}
