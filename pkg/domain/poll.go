package domain

type Poll struct {
	//variable       type               notation
	Question string            `json:"question"`
	Answers  map[string]string `json:"answers"`
}

type Answer struct {
	Answer string `json:"answer"`
}
