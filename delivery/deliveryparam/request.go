package deliveryparam

type Request struct {
	Command           string
	CreateTaskRequest CreateTaskRequest
}

type CreateTaskRequest struct {
	Title      string
	DueDate    string
	CategoryID int
}
