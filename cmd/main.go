package main

import (
	"fmt"
	"github.com/ihsan-aryandi/go-router"
	"strconv"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var users []*User
var seq int64 = 1

func main() {
	port := ":8000"

	router := gorouter.NewRouter()
	router.GET("/text/greet", func(ctx *gorouter.Context) {
		ctx.Write(200, "Hello World")
	})

	router.GET("/json/greet", func(ctx *gorouter.Context) {
		ctx.JSON(200, gorouter.Map{
			"message": "Hello World",
		})
	})

	/*
		Group Routes
	*/
	router.Routes("/user", func(route *gorouter.GroupRoutes) {
		route.GET(func(ctx *gorouter.Context) {
			ctx.JSON(200, gorouter.Map{
				"users": users,
			})
		})

		route.POST(func(ctx *gorouter.Context) {
			user := &User{}

			if err := ctx.Body(user); err != nil {
				ctx.JSON(400, gorouter.Map{
					"error": "Invalid request",
				})

				return
			}

			user.Id = seq
			seq++

			users = append(users, user)

			ctx.JSON(200, gorouter.Map{
				"users": users,
			})
		})
	})

	router.Routes("/user/{id}", func(route *gorouter.GroupRoutes) {
		route.GET(func(ctx *gorouter.Context) {
			strId := ctx.Param("id")
			id, _ := strconv.Atoi(strId)

			for _, user := range users {
				if user.Id == int64(id) {
					ctx.JSON(200, user)
					return
				}
			}

			ctx.JSON(400, gorouter.Map{
				"error": "user not found",
			})
		})
	})

	router.Use(func(handler gorouter.Handler) gorouter.Handler {
		return func(ctx *gorouter.Context) {
			fmt.Println("This is middleware")

			handler(ctx)
		}
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := router.Listen(":8000"); err != nil {
		fmt.Printf("err : %s", err.Error())
	}
}
