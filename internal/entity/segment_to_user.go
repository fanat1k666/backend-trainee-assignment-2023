package entity

import "time"

type SegmentToUser struct {
	Id           int
	EntryTime    time.Time
	ExitTime     time.Time
	PlanExitTime time.Time
	UserId       int
	SegmentId    int
}
