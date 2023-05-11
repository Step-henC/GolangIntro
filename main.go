package main

import (
	"net/http" //built in

	"github.com/gin-gonic/gin" //downloaded go-get
	//"errors"
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
func main() {
	router := gin.Default() //create router to handle routes
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
