package main

type User struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     string  `json:"email"`
	Debt      float64 `json:"debt"`
	Credit    float64 `json:"credit"`
	Password  string  `json:"password"`
}

func (u *User) Add(amount float64) {
	u.Credit += amount
}
