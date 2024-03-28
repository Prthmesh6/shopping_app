package models

type Brand string
type Category string
type Cart map[int64]Item

type Item struct {
	ID          int64
	Discription string // optional
	Price       int64  //
	Brand       Brand
	Category    Category
}

// wallet
type User struct {
	Id   int64
	Name string
	Cart Cart
}
