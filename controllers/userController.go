package controllers

import (
	"context"
	"fmt"
	"golang-jwt/auth/helpers"
	helper "golang-jwt/auth/helpers"
	"golang-jwt/auth/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/marif226/golang-jwt-auth/database"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()