package vo

type Book struct {
	Id        int64  `json:"Id"`
	Name      string `json:"Name"`
	Remark    string `json:"Remark"`
	CreatedBy string `json:"Created_by"`
}
