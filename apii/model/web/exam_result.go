package web

import "time"

type ExamRequest struct {
	UserId   int
	CourseId int
	Anwers   []int `json:"anwers"`
}

type QuestionWithOptions struct {
	Content string `json:"content"`
}

type ExamResultResponse struct {
	CourseId           int       `json:"course_id"`
	UserId             int       `json:"user_id"`
	Score              int       `json:"score"`
	TotalQuestions     int       `json:"total_questions"`
	Attempt            int       `json:"attempt"`
	SubmittedAt        time.Time `json:"submitted_at"`
	RewardAddress      string    `json:"reward_address"`
	CertificateAddress string    `json:"certificate_address"`
}