package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "strconv"
    _"net/http"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    gorm.Model
    Name   string
    Email string
    Password string
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&User{})
    defer db.Close()
}

//DB追加
func dbInsert(name string, email string, password string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&User{Name: name, Email: email, Password:password})
    defer db.Close()
}

//DB更新
func dbUpdate(id int, name string, email string, password string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbUpdate)")
    }
    var user User
    db.First(&user, id)
    user.Name = name
    user.Password = password
    db.Save(&user)
    db.Close()
}

//DB削除
func dbDelete(id int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbDelete)")
    }
    var user User
    db.First(&user, id)
    db.Delete(&user)
    db.Close()
}

//DB全取得
func dbGetAll() []User {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetAll())")
    }
    var users []User
    db.Order("created_at desc").Find(&users)
    db.Close()
    return users
}

//DB一つ取得
func dbGetOne(id int) User {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetOne())")
    }
    var user User
    db.First(&user, id)
    db.Close()
    return user
}


func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    dbInit()

    //Index
    router.GET("/", func(ctx *gin.Context) {
        users := dbGetAll()
        ctx.HTML(200, "index.html", gin.H{
            "users": users,
        })
    })

    //Create
    router.POST("/new", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        email := ctx.PostForm("email")
        password := ctx.PostForm("password")
        dbInsert(name, email, password)
        ctx.Redirect(302, "/")
    })

    //Detail
    router.GET("/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        user := dbGetOne(id)
        ctx.HTML(200, "detail.html", gin.H{"user": user})
    })

    //Update
    router.POST("/update/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        name := ctx.PostForm("name")
        email := ctx.PostForm("email")
        password := ctx.PostForm("password")
        dbUpdate(id, name, email, password)
        ctx.Redirect(302, "/")
    })

    //削除確認
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        user := dbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"user": user})
    })

    //Delete
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        ctx.Redirect(302, "/")

    })

    router.Run(":8080")
}
