package dto

type Course struct {
	CourseName string        `json:"course_name" example:"2110101 Calculus I"`
	Exams      []ExamSummary `json:"exams"`
}

type CourseSummary struct {
	CourseId    string `json:"course_id" example:"TODO!(MONGOOBJECTID)"`
	CourseName  string `json:"course_name" example:"2011 - S1 - Final"`
	ThreadCount int32  `json:"thread_count" example:"12"`
}
