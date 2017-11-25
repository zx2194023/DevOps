package main

import (
   "net/http"
  
   "github.com/labstack/echo"
  "strconv"
  "errors"
)

func main() {
   e := echo.New()
   e.GET("/", func(c echo.Context) error {
       return c.String(http.StatusOK, "Hello, World!")
   })

   e.GET("/add/:num1/:num2", func(c echo.Context) error {
     num1Str := c.Param("num1")
     num2Str := c.Param("num2")
     num1, err := strconv.Atoi(num1Str)
     if err != nil {
        return c.String(http.StatusUnprocessableEntity, err.Error())
     }
     num2, err := strconv.Atoi(num2Str)
     if err != nil {
        return c.String(http.StatusUnprocessableEntity, err.Error())
     }
     result := Add(num1, num2)
     return c.String(http.StatusOK, strconv.Itoa(result))
  })

  e.GET("/div/:num1/:num2", func(c echo.Context) error {
     num1Str := c.Param("num1")
     num2Str := c.Param("num2")
     num1, err := strconv.Atoi(num1Str)
     if err != nil {
        return c.String(http.StatusUnprocessableEntity, err.Error())
     }
     num2, err := strconv.Atoi(num2Str)
     if err != nil {
        return c.String(http.StatusUnprocessableEntity, err.Error())
     }
     result, err := Div(num1, num2)
     if err != nil {
        return c.String(http.StatusUnprocessableEntity, err.Error())
     }
     return c.String(http.StatusOK, strconv.Itoa(result))
  })

  e.Logger.Fatal(e.Start(":1325"))
}

func Add(num1, num2 int) int {
  return num1 + num2
}

func Div(num1, num2 int) (int, error) {
  if num2 == 0{
     return 0, errors.New("number 2 cnannot be 0")
  }
  return num1 / num2, nil
}
