package main

import (
	"net/http" //built in

	"errors"

	"github.com/gin-gonic/gin" //downloaded go-get
)

//set up dependency tracking
//command line: go mod init example/name-of-folder

//command line: go get ...whatever dependency needed
//using this to install gin framework

type book struct {
	ID       string `json:"id"` //capitalized names make it an exported/public field
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{ // in memory api as a start to golang
	{ID: "1", Title: "1984", Author: "George Orwell", Quantity: 92},
	{ID: "2", Title: "The Catcher and the Rye", Author: "I can't remember", Quantity: 12},
	{ID: "3", Title: "Kite Runner", Author: "Also cannot remember", Quantity: 1000},
	{ID: "4", Title: "A Tale of Two Cities", Author: "Charles Dickens", Quantity: 4},
	{ID: "5", Title: "The Grapes of Wrath", Author: "John Steinbeck", Quantity: 55},
}

func getBooks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) //H allows customized json
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func addBook(c *gin.Context) { // context stores all info related to specific req
	var newBook book
	// use c to bind json data payload to this new book object
	if err := c.BindJSON(&newBook); err != nil {
		return //the BindJson returns the error response auto if err is not null
	}

	books = append(books, newBook)              //add book to list
	c.IndentedJSON(http.StatusCreated, newBook) //return the book we just created

}

func checkOutBook(c *gin.Context) {

	id, ok := c.GetQuery("id") // query looks like ?id=2
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) //H allows customized json
		return
	}

	if book.Quantity <= 0 {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Available"}) //H allows customized json
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // query looks like ?id=2
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) //H allows customized json
		return
	}

	if book.Quantity <= 0 {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Available"}) //H allows customized json
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
func main() {
	router := gin.Default() //create router to handle routes
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkOutBook)
	router.Run("localhost:8080")
}
