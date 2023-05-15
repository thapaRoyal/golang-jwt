package helpers

// imports
import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CheckUserType checks if the user type stored in the context matches the specified role.
func CheckUserType(c *gin.Context, role string) (err error) {
	// Retrieve the user type from the context
	userType := c.GetString("user_type")
	err = nil // Initialize the error variable as nil

	// Compare the user type with the specified role
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

// MatchUserTypeToUid checks if the user type and user ID stored in the context match the specified values.
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	// Retrieve the user type and user ID from the context
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil // Initialize the error variable as nil

	// If the user type is "USER" and the user ID doesn't match the specified value, return an error
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	// Check if the user type matches the specified value by calling the CheckUserType function
	err = CheckUserType(c, userType)
	return err
}
