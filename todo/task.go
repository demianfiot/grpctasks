package todo

type Task struct {
	Id             int64
	Title          string
	Difficulty     string
	XpReward       int32
	AssignedUserId int64
	Completed      bool
}
