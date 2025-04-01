package model

import "time"

type LikecoinEvmMigrateBookSubmissionStatus string

const (
	LikecoinEvmMigrateBookSubmissionStatusSuccess LikecoinEvmMigrateBookSubmissionStatus = "success"
	LikecoinEvmMigrateBookSubmissionStatusFailed  LikecoinEvmMigrateBookSubmissionStatus = "failed"
)

type LikecoinEvmMigrateBookSubmission struct {
	Id           uint64
	CreatedAt    time.Time
	LikeClassID  string
	EvmClassID   string
	Status       LikecoinEvmMigrateBookSubmissionStatus
	FailedReason *string
}
