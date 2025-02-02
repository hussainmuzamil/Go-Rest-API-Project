package types

type Student struct {
	Id    int
	Name  string `json:"name" required:"true"`
	Age   int    `json:"age" required:"true"`
	Email string `json:"email" required:"true"`
}
