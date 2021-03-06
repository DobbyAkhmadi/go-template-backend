package handlers

import (
	"backend/internal/app/user/models"
	"backend/internal/app/user/repository"
	"backend/pkg/entity"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

var userRepository repository.UserRepository

func init() {
	userRepository = repository.NewUserRepository()
}

// GetAllUsers gets all repository information
func GetAllUsers(c *fiber.Ctx) error {
	users := userRepository.FindAll()

	resp := entity.Response{
		Code:    http.StatusOK,
		Body:    users,
		Title:   "GetAllUsers",
		Message: "All users",
	}

	return c.Status(resp.Code).JSON(resp)
}

// GetSingleUser Gets single user information
func GetSingleUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting user information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	user, err := userRepository.FindByID(id)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting user information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if user == nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("user with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := entity.Response{
		Code:    http.StatusOK,
		Body:    user,
		Title:   "OK",
		Message: "User information",
	}
	return c.Status(resp.Code).JSON(resp)

}

// AddNewUser adds new user
func AddNewUser(c *fiber.Ctx) error {
	user := &models.User{}

	err := c.BodyParser(user)

	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing user body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := userRepository.Save(*user)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	user, err = userRepository.FindByID(id)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if user == nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("user with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := entity.Response{
		Code:    http.StatusOK,
		Body:    user,
		Title:   "OK",
		Message: "new user added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

// UpdateUser updates a user by user id
func UpdateUser(c *fiber.Ctx) error {
	user := &models.User{}

	err := c.BodyParser(user)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing user body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing user ID. (it should be an integer)",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	updatingUser, err := userRepository.FindByID(id)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting user information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if updatingUser == nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("user with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	user.ID = id

	err = userRepository.Update(*user)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in updating user information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	user, err = userRepository.FindByID(id)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly updated user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if user == nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("user with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := entity.Response{
		Code:    http.StatusOK,
		Body:    user,
		Title:   "UpdateUser",
		Message: "user updated successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

// DeleteUser deletes the user from db
func DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in getting user information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	user, err := userRepository.FindByID(id)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if user == nil {
		errorResp := entity.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("user with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding user",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	err = userRepository.Delete(*user)
	if err != nil {
		errorResp := entity.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in deleting user object",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := entity.Response{
		Code:    http.StatusOK,
		Body:    "user deleted successfully",
		Title:   "OK",
		Message: "user deleted successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}
