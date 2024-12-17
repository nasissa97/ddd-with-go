package tdd

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -package mocks -destination mocks/cookies.go tdd CookieStockChecker,CardCharger,EmailSender
