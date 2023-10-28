package controller

import (
	gorm "gorm.io/gorm"
)

type AppController struct {
	DB *gorm.DB
}
