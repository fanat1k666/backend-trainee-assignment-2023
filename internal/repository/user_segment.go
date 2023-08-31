package repository

type UserSegment interface {
	CreateSegment(name string) error
	DropSegment(name string) error
	UserSegment(userId int, segmentName []string) error
	ShowSegment(userId int) ([]ShowUsersSegment, error)
	DropUserFromSegment(userId int, name string) error
}

type ShowUsersSegment struct {
	UserId int
	Name   string
}
