package module

type DirInfo struct {
	Id int64
	Name    string `json:"pathname"`
	DirName string `json:"name"`
}

type Top struct {
	DirName string `json:"pathname"`
	TuName  string `json:"tuname"`
	Count   string `json:"count"`
	Time   string `json:"time"`
}
