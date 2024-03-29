package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "hello world",
	// })

	// name := c.Params.ByName("name")
	// c.JSON(200, gin.H{
	// 	"message": "hello " + name,
	// })

	c.XML(200, Person{FirstName: "Abdul", LastName: "Somad"})

}

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "helloq-workd",
	// 	})
	// })

	router := gin.Default()
	// router.GET("/", IndexHandler)
	// router.GET("/:name", IndexHandler)
	router.GET("/", IndexHandler)

	router.Run(":5000")
}

curl -X GET http://172.16.2.23:5000