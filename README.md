# Go Router

Go Router adalah library http router sederhana yang dibuat menggunakan bahasa Go

___

## Instalasi

`go get -u github.com/bandungrhapsody/rhaprouter`

### Memulai Go Router

Cukup import library dengan mengetik :

`import "github.com/ihsan-aryandi/go-router"`

### Dasar menjalankan Go Router

Contoh cara untuk menjalankan Go Router :

    port := ":8000"

    router := gorouter.NewRouter()
    router.GET("/greet", func(ctx *gorouter.Context) {
        ctx.Write(200, "Hello World")
    })
    router.Listen(":8000")

Kode di atas akan menjalankan server di port **:8000**

### Group Routes

Dengan fungsi ini, anda dapat membuat banyak endpoint dengan method yang bermacam-macam di satu path yang sama.

Cara menggunakannya cukup jalankan fungsi `GroupRoutes`, berikut contoh penggunaannya :

    port := ":8000"

	router := gorouter.NewRouter()
    router.Routes("/user", func(route *gorouter.GroupRoutes) {
		route.GET(func(ctx *gorouter.Context) {
			// Ambil data pengguna
		})

		route.POST(func(ctx *gorouter.Context) {
			// Tambah data pengguna
		})
	})
    router.Listen(":8000")