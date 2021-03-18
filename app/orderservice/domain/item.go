package domain


type Item struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
}
