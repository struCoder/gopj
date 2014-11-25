package model

//用来保存登陆成功用户的一些信息存到session中
type UserInfo struct {
	Id   string
	Name string
}

type ComplaintMsg struct {
	Id          int
	Name        string
	Phone       string
	BeComplaint string
	Reason      string
	Status      string
	DealPerson  string
}
