package handler

import (
	"fmt"
	"net/http"
	model "netl/pkg/models/student"

	"github.com/labstack/echo/v4"
)

// RegisterUser is a handler function that registers a user by persisting its data into the database with
// service methods
func (h *Handler) RegisterUser(c echo.Context) error {
	// inits
	register := model.RegisterModel{}
	err := c.Bind(&register)
	// checking if binding fails
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		h.logger.Print("handler failed")
		return err
	}
	// creating context
	derivedContext := c.Request().Context()
	// accessign user service layer
	serviceErr := h.userSvc.AuthorizeUser(derivedContext, register)
	// failure check on user service
	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, serviceErr)
		fmt.Println(serviceErr)
		h.logger.Print("user service failed")
	} else {
		c.JSON(http.StatusOK, "student was registered successfully")
	}
	return nil
}

// LoginUser is a handler method that logs in a student by validating the user data from the database
func (h *Handler) LoginUser(c echo.Context) error {
	// inits
	login := model.LoginPayload{}
	err := c.Bind(&login)
	// checking if errors occured during binding the login payload
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		h.logger.Err(err)
		return err
	}
	// creating context
	derivedContext := c.Request().Context()
	// accessing the user service layer
	serviceErr := h.userSvc.AuthenticateUser(derivedContext, login)
	// failurue check on user service
	if serviceErr != nil {
		c.JSON(http.StatusBadRequest, "failed to login")
		h.logger.Err(err)
		return serviceErr
	} else {
		// returning a successfull login response
		c.JSON(http.StatusOK, "successfull login")
	}
	return nil
}
