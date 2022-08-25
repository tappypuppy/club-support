package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "strconv"
    _"net/http"
    _ "github.com/mattn/go-sqlite3"
)

// ユーザDB
type User struct {
    gorm.Model
    Name   string
    Email string
    Password string
}

type Event struct {
    gorm.Model
    Title   string
    Type string
    StartTime string
    EndTime string
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&Event{})
    db.AutoMigrate(&User{})
    defer db.Close()
}

//DB追加
func dbInsertUser(name string, email string, password string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&User{Name: name, Email: email, Password:password})
    defer db.Close()
}

//DB更新
func dbUpdateUser(id int, name string, email string, password string) {
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
func dbDeleteUser(id int) {
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
func dbGetAllUser() []User {
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
func dbGetOneUser(id int) User {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetOne())")
    }
    var user User
    db.First(&user, id)
    db.Close()
    return user
}


//DB追加
func dbInsertEvent(title string, types string, starttime string, endtime string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&Event{Title: title, Type: types, StartTime: starttime, EndTime: endtime})
    defer db.Close()
}

//DB更新
func dbUpdateEvent(id int, title string, types string, starttime string, endtime string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbUpdate)")
    }
    var event Event
    db.First(&event, id)
    event.Title = title
    event.Type = types
    event.StartTime = starttime
    event.EndTime = endtime
    db.Save(&event)
    db.Close()
}

//DB削除
func dbDeleteEvent(id int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbDelete)")
    }
    var event Event
    db.First(&event, id)
    db.Delete(&event)
    db.Close()
}

//DB全取得
func dbGetAllEvent() []Event {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetAll())")
    }
    var events []Event
    db.Order("created_at desc").Find(&events)
    db.Close()
    return events
}

//DB一つ取得
func dbGetOneEvent(id int) Event {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetOne())")
    }
    var event Event
    db.First(&event, id)
    db.Close()
    return event
}



func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    dbInit()

    router.GET("/event", func(ctx *gin.Context) {
        events := dbGetAllEvent()
        ctx.HTML(200, "index_event.html", gin.H{
            "events": events,
        })
    })

    //Index
    router.GET("/user", func(ctx *gin.Context) {
        users := dbGetAllUser()
        ctx.HTML(200, "index_user.html", gin.H{
            "users": users,
        })
    })

    //Create
    router.POST("/new_user", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        email := ctx.PostForm("email")
        password := ctx.PostForm("password")
        dbInsertUser(name, email, password)
        ctx.Redirect(302, "/user")
    })

    //Create
    router.POST("/new_event", func(ctx *gin.Context) {
        title := ctx.PostForm("title")
        types := ctx.PostForm("types")
        starttime := ctx.PostForm("starttime")
        endtime := ctx.PostForm("starttime")
        dbInsertEvent(title, types, starttime, endtime)
        ctx.Redirect(302, "/event")
    })

    //Detail
    router.GET("/detail_user/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        user := dbGetOneUser(id)
        ctx.HTML(200, "detail_user.html", gin.H{"user": user})
    })

    //Update
    router.POST("/update_user/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        name := ctx.PostForm("name")
        email := ctx.PostForm("email")
        password := ctx.PostForm("password")
        dbUpdateUser(id, name, email, password)
        ctx.Redirect(302, "/user")
    })

    //削除確認
    router.GET("/delete_check_user/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        user := dbGetOneUser(id)
        ctx.HTML(200, "delete_user.html", gin.H{"user": user})
    })

    //Delete
    router.POST("/delete_user/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDeleteUser(id)
        ctx.Redirect(302, "/user")

    })

    router.Run(":8080")
}
