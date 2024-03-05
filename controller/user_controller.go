package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}