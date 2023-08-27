package repository

type UserSegment interface {
	CreateSegment(name string) error
	DropSegment(name string) error
	AddSegmentToUser(userId int, segmentId []string) error
}
