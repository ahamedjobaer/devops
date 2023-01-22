package main

import (
	"log"
	"net/http"
	"strconv"

	"devops/my_math"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/shyfur-t", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Shyfur, Sumon")
	})

	e.GET("/math/sum", func(c echo.Context) error {
		number1, err := strconv.Atoi(c.QueryParam("number1"))

		if err != nil {
			log.Println("Error during conversion")
			return c.String(http.StatusInternalServerError, "Error during number 1 conversion")
		}

		number2, err2 := strconv.Atoi(c.QueryParam("number2"))

		if err2 != nil {
			log.Println("Error during conversion")
			return c.String(http.StatusInternalServerError, "Error during number 1 conversion")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"result": my_math.Sum(number1, number2),
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
