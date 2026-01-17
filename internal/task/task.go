package task

import (
	"rpcprac/internal/user"
	"sync"
)

type Task struct {
	Id             int64
	Title          string
	Difficulty     string
	XpReward       int32
	AssignedUserId int64
	Completed      bool
}

type TaskStorage struct {
	mu     sync.Mutex
	tasks  map[int64]*Task
	nextID int64
	us     *user.UserStorage
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		tasks:  make(map[int64]*Task),
		nextID: 1,
	}
}

func (s *TaskStorage) CreateTask(title, difficulty string, xp_reward int32) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &Task{
		Id:         s.nextID,
		Title:      title,
		Difficulty: difficulty,
		XpReward:   xp_reward,
	}

	s.tasks[s.nextID] = task
	s.nextID++

	return task
}
func (s *TaskStorage) GetTask(id int64) *Task {
	task, ok := s.tasks[id]
	if !ok {
		return &Task{}
	}

	return task
}
func (s *TaskStorage) ListTasks() []Task {
	taskslist := make([]Task, 0, len(s.tasks))

	for _, user := range s.tasks {
		taskslist = append(taskslist, *user)
	}

	return taskslist
}
func (s *TaskStorage) AssignTask(taskid, userid int64) *Task {
	task := s.tasks[taskid]
	task.AssignedUserId = userid
	return task
}

func (s *TaskStorage) CompleteTask(id int64) *Task {
	task := s.tasks[id]
	task.Completed = true
	user := s.us.GetUser(task.Id)
	user.Xp += task.XpReward
	return task
}
