# Go Router

Go Router adalah library http router sederhana yang dibuat menggunakan bahasa Go

---

## Instalasi

Buka terminal Anda lalu ketik :

`go get -u github.com/ihsan-aryandi/go-router`

---

### Memulai Go Router

Cukup import library dengan mengetik :

```go
import "github.com/ihsan-aryandi/go-router"
```

### Dasar menjalankan Go Router

Contoh cara untuk menjalankan Go Router :

``` go
package main

import "github.com/ihsan-aryandi/go-router"

func main() {
    port := ":8000" // Lebih baik jika disimpan di file config atau env.

    router := gorouter.NewRouter()
    router.GET("/greet", func(ctx *gorouter.Context) {
        ctx.Write(200, "Hello World")
        // Output: Hello World
    })
    router.Listen(port)
}
```

Kode di atas akan menjalankan server di port **:8000**

### JSON Response
Untuk mengembalikan response dengan tipe JSON, cukup panggil fungsi `JSON` yang ada dalam `gorouter.Context`.

Berikut cara penggunaannya :
``` go
package main

import "github.com/ihsan-aryandi/go-router"

type Response struct {
    Message string `json:"message"`
}

func main() {
    port := ":8000" // Lebih baik jika disimpan di file config atau env.

    router := gorouter.NewRouter()
    router.GET("/greet", func(ctx *gorouter.Context) {
        res := &Response{
            Message: "Hello World",
        }
        
        ctx.JSON(200, res)
        // Output : {"message": "Hello World"}
    })
    router.Listen(port)
}
```

Atau jika anda tidak ingin menggunakan **struct** sebagai response. Go Router menyediakan tipe **map** agar anda dapat menambahkan custom field sendiri.

Cara menggunakan map cukup ketik `gorouter.Map{}`, berikut contoh penggunaannya :

``` go
package main

import "github.com/ihsan-aryandi/go-router"

func main() {
    port := ":8000" // Lebih baik jika disimpan di file config atau env.

    router := gorouter.NewRouter()
    router.GET("/greet", func(ctx *gorouter.Context) {
        ctx.JSON(200, gorouter.Map{
            "message": "Hello World"
        })
        // Output : {"message": "Hello World"}
    })
    router.Listen(port)
}
```

### Param
Param adalah nilai yang diambil dari path, Anda dapat melakukan itu dengan cara memanggil fungsi `Param` yang ada di dalam `gorouter.Context`. 

Fungsi `Param` ini membutuhkan satu parameter yaitu key (string). Nilai yang dikembalikan dari fungsi ini adalah string.

Ketika menulis path, Anda harus menambahkan placeholder dalam path yang anda buat. Format placeholdernya adalah `{nama_param}`, contoh path `/user/{id}` 

Contoh penggunaan :
``` go
package main

import "github.com/ihsan-aryandi/go-router"

func main() {
    port := ":8000" // Lebih baik jika disimpan di file config atau env.

    router := gorouter.NewRouter()
    router.GET("/user/{id}", func(ctx *gorouter.Context) {
        id := ctx.Param("id") // Key harus sama dengan placeholder pada path
        // Contoh Output : 10
        
        // Lakukan sesuatu...
    })
    router.Listen(port)
}
```

### Group Routes

Dengan fungsi ini, anda dapat membuat banyak endpoint dengan method yang bermacam-macam di satu path yang sama.

Cara menggunakannya cukup jalankan fungsi `Routes`, berikut contoh penggunaannya :

``` go
package main

import "github.com/ihsan-aryandi/go-router"

func main() {
    port := ":8000" // Lebih baik jika disimpan di file config atau env.

    router := gorouter.NewRouter()
    router.Routes("/user", func(route *gorouter.GroupRoutes) {
        route.GET(func(ctx *gorouter.Context) {
            // Ambil data pengguna
        })
        
        route.POST(func(ctx *gorouter.Context) {
            // Tambah data pengguna
        })
    })
    router.Listen(port)
}
```
