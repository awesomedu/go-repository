package model


type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int,name int,gender int,age int,phone int,email int,) Customer{
	return customer{
	Id : id
	Name : name
	Gender : gender
	Age : age
	Phone : phone
	Email : email
	}
}