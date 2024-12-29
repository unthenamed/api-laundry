package model

type Details struct {
	Id           int `json:"id"`
	BillId       int `json:"billId"`
	ProductId    int `json:"productId"`
	ProductPrice int `json:"productPrice"`
	Qty          int `json:"qty"`
}

type bills struct {
	Id          int       `json:"id"`
	BillDate    string    `json:"billDate"`
	EntryDate   string    `json:"entryDate"`
	FinishDate  string    `json:"finishDate"`
	EmployeeId  int       `json:"employeeId"`
	CustomerId  int       `json:"customerId"`
	BillDetails []Details `json:"billDetails"`
}

type RDetails struct {
	Id           int     `json:"id"`
	BillId       int     `json:"billId"`
	Product      Product `json:"products"`
	ProductPrice int     `json:"productsPrice"`
	Qty          int     `json:"qty"`
}

type response struct {
	Id          int        `json:"id"`
	BillDate    string     `json:"billDate"`
	EntryDate   string     `json:"entryDate"`
	FinishDate  string     `json:"finishDate"`
	Employee    Employees  `json:"employees"`
	Customer    Customers  `json:"customers"`
	BillDetails []RDetails `json:"billDetails"`
}

type query struct {
	StartDate   string
	EndDate     string
	ProductName string
}

type Transaction struct {
	Bills    bills    `json:"bills"`
	Response response `json:"response"`
	Query    query    `json:"query"`
}
