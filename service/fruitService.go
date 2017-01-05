package service

import (
	. "gomssql/model"

	"net/http"

	"fmt"

	"github.com/labstack/echo"
)

func Query(c echo.Context) error {
	fmt.Println(1111)
	db := Mssql{}
	if err := db.Open(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer db.Close()
	fmt.Println(db)

	rows, err := db.Query("select  Name,Code from dbo.Fruit")
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	} else {
		var fruit Fruit
		var fruits []Fruit
		for rows.Next() {
			rows.Scan(&fruit.Name, &fruit.Code)
			fruits = append(fruits, fruit)
		}
		fmt.Printf("%+v", fruits)
		return c.JSON(http.StatusOK, fruits)
	}

}
