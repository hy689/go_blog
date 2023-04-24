package model

type Article struct {
	Category Category
	Title    string `json:"title"`
	Cid      string `json:"cid"`
	Content  string `json:"content"`
	Img      string `json:"img"`
	Desc     string `json:"desc"`
}
