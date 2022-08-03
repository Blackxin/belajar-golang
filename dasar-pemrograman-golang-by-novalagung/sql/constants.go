package main

import "os"

var dbUser string = os.Getenv("DB_USER")
var dbPwd string = os.Getenv("DB_PWD")
