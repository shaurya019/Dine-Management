package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	helper "golang-restaurant-management/helpers"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}