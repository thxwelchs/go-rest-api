package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/person", func(c echo.Context) error {
		// Person 구조체
		var Person struct {
			Name    string   `json:"name"`
			Age     int      `json:"age"`
			Hobbies []string `json:"hobbies"`
		}

		Person.Name = "이태훈"
		Person.Age = 29
		// hobbies := make([]string, 3)
		// hobbies[0] = "영화보기"
		// hobbies[1] = "음악듣기"
		// hobbies[2] = "개발하기"
		Person.Hobbies = []string{"영화보기", "음악듣기", "개발하기"}

		return c.JSON(http.StatusOK, &Person)

	})

	e.Logger.Fatal(e.Start(":8080")) // http server listen 8080
}
