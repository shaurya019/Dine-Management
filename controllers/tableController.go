package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}