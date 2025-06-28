package data

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{
		ID:       1,
		Name:     "The Great Gatsby",
		Author:   "F. Scott Fitzgerald",
		Quantity: 10,
	},
	{
		ID:       2,
		Name:     "To Kill a Mockingbird",
		Author:   "Harper Lee",
		Quantity: 5,
	},
	{
		ID:       3,
		Name:     "1984",
		Author:   "George Orwell",
		Quantity: 8,
	},
}
