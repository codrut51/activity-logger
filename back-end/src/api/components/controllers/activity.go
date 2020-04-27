package controllers

import (
	"time"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Activity struct {
	ActivityId string `json:"id"`
	ActivityTitle string `json:"title"`
	ActivityDate time.Time `json:"date"`
}