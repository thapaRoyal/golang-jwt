package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/thaparoyal/golang-jwt/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollecion *mongo.Collection = database.OpenCollection(database.Client, "User")
var validate = validator.New()

func Hashpassword() {}

func VerifyPassword() {}

func Signup() {}

func Login() {}

func GetUsers() {}

func GetUser()
