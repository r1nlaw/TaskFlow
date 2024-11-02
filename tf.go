package taskflow

type TaskFlowList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TaskFlowItem struct {
	NameCard      string `json:"name"`
	Content       string `json:"content"`
	Status        string `json:"status"`
	ChangeContent string `json:"changecontent"`
	ChangeStatus  string `json:"сhangestatus"`
	Id            int    `json:"id"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
