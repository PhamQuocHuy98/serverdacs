package main

import(
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)
func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Golang =))")
}
func main()  {
	t:=10
	fmt.Print(t)
	e := echo.New()

	e.GET("/", yallo)

	e.Start(":8000")
}
