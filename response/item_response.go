package response

import "github.com/labstack/echo"

// item struct to represent a webshop listing e.g. bol.com
type ItemResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
