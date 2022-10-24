package dto

type ThreadSummary struct {
	ThreadId           string   `json:"thread_id" example:"507f1f77bcf86cd799439011"`
	ProblemName        string   `json:"problem_name" example:"1+1=???"`
	ProblemDescription string   `json:"problem_description" example:"What is 'Monad'?"`
	Upvote             int32    `json:"upvote" example:"12"`
	Downvote           int32    `json:"downvote" example:"12"`
	Tags               []string `json:"tags" example:"calculus,2110101,limit,hard"`
}

type Thread struct {
	ThreadId           string   `json:"thread_id" example:"507f1f77bcf86cd799439011"`
	ProblemName        string   `json:"problem_name" example:"1+1=???"`
	ProblemDescription string   `json:"problem_description" example:"What is 'Monad'?"`
	Upvote             int32    `json:"upvote" example:"12"`
	Downvote           int32    `json:"downvote" example:"12"`
	Tags               []string `json:"tags" example:"calculus,2110101,limit,hard"`
	Replies            []Reply  `json:"replies"`
}

type Reply struct {
	ThreadId string `json:"thread_id" example:"507f1f77bcf86cd799439011"`
	Body     string `json:"body" example:"1+1=3"`
	Upvote   int32  `json:"upvote" example:"42"`
	Downvote int32  `json:"downvote" example:"420"`
}

type CreateThread struct {
	CourseName         string   `json:"course_name" example:"2110101 Calculus I" validate:"required"`
	ExamName           string   `json:"exam_name" example:"2011 - S1 - Final" validate:"required"`
	ProblemName        string   `json:"problem_name" example:"1+1=???" validate:"required"`
	ProblemDescription string   `json:"problem_description" example:"What is 'Monad'?" validate:"required"`
	Tags               []string `json:"tags" example:"calculus,2110101,limit,hard" validate:"required"`
}

type CreateReply struct {
	Body string `json:"body" example:"1+1=3" validate:"required"`
}

type CreateThreadResponse struct {
	ThreadId string `json:"thread_id" example:"507f1f77bcf86cd799439011"`
}
