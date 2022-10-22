package dto

type ThreadSummary struct {
	ThreadId           string   `json:"thread_id" example:"TODO!(OBJECTID)"`
	ProblemName        string   `json:"problem_name" example:"1+1=???"`
	ProblemDescription string   `json:"problem_description" example:"What is 'Monad'?"`
	Upvote             int32    `json:"upvote" example:"12"`
	Downvote           int32    `json:"downvote" example:"12"`
	Tags               []string `json:"tags" example:"calculus,2110101,limit,hard"`
}
