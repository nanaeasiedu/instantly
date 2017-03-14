package web

import (
	"github.com/ngenerio/instantly/pkg/models"
)

type Params struct {
	Title   string
	Flashes []interface{}
	User    models.User
	Data    map[string]interface{}
}
