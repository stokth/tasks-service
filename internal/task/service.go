package task

type Service struct {
	repo TasksRepository
}

func NewService(repo TasksRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListTasks() ([]Task, error) {
	return s.repo.ListTasks()
}

func (s *Service) GetTask(id int64) (*Task, error) {
	return s.repo.GetTask(id)
}

func (s *Service) CreateTask(task *Task) (*Task, error) {
	tsk := Task{
		Title:  task.Title,
		UserID: task.UserID,
	}

	if err := s.repo.CreateTask(&tsk); err != nil {
		return &Task{}, err
	}

	return &tsk, nil
}

func (s *Service) UpdateTask(id int64, task *Task) (*Task, error) {
	tsk, err := s.repo.GetTask(id)
	if err != nil {
		return &Task{}, err
	}

	tsk.Title = task.Title

	if err := s.repo.UpdateTask(tsk); err != nil {
		return &Task{}, err
	}

	return tsk, nil
}

func (s *Service) DeleteTask(id int64) error {
	return s.repo.DeleteTask(id)
}
