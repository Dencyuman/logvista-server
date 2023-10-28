//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Dencyuman/logvista-server/src/database"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initializeDB() (*gorm.DB, error) {
	wire.Build(database.Connect)
	return &gorm.DB{}, nil
}
