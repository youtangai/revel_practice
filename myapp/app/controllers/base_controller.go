package controllers

import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    "strings"
    "fmt"
    "myapp/app/models"
    _ "github.com/go-sql-driver/mysql"
)

type BaseController struct {
    *revel.Controller
}

func convertJsonFormat(success string, data interface{}, err models.JsonError) models.JsonResponse {
    return models.JsonResponse{success, data, err}
}

func connectDB() *gorm.DB {
  db, err := gorm.Open("mysql", getConnectionString())

  if err != nil {
    revel.ERROR.Println("FATAL", err)
    panic(err)
  }
  return db
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")
    timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}
