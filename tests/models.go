package tests

type APIResponse struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Pet struct {
	Id        int64      `json:"id"`
	Categorys []Category `json:"category"`
	Name      string     `json:"name"`
	Photo_url []string   `json:"photo_url"`
	Tags      []Tag      `json:"tags"`
	Status    string     `json:"status"`
}

type Order struct {
	Id        int64    `json:"id"`
	Pet_id    int64    `json:"pet_id"`
	Quantity  int32    `json:"quantity"`
	Ship_date string   `json:"ship_date"`
	Status    []string `json:"status"`
	Complete  bool     `json:"complete"`
}
type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
}
