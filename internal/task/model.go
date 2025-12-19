package task

type Task struct {
	ID     int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `json:"title"`
	UserID int64  `json:"user_id"`
}

type TaskRequest struct {
	Title  string `json:"title"`
	UserID int64  `json:"user_id"`
}
