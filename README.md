# Learn Golang

Repo untuk coba-coba web server menggunakan Go Lang

## Dependencies

| Package | Usage | Repo | Docs |
| --- | --- | ---| --- |
| Echo | Web framework | github.com/labstack/echo | https://echo.labstack.com/guide |
| Logrus | Logging activity | github.com/sirupsen/logrus | https://godoc.org/github.com/sirupsen/logrus |

## Command

Start apps
```go run server.go```

## Project trees

**_! Disclaimer_**: _project tree_ dibawah ini saya buat berdasarkan pengalaman pribadi dengan mempertimbangkan fungsi dan efisiensi pemrograman. Silahkan ubah sesuai dengan kebutuhan dan gaya pemrograman masing-masing.

```
learn_golang/
|_ handler/
    |_ v1/
        |_ databases/
        |_ useCases/
        |_ index.go
|_ modules/logger/
    |_ logger.go
|_ routes/
    |_ v1.go
|_ webserver/setupmiddleware/
    |_ setupmiddleware.go
|_ store/
|_ server.go
```

#### Keterangan:
1. File `server.go` adalah main program untuk menjalankan web server. Didalamnya berisi konfigurasi dari mengenai service dari server yang berjalan seperti port, middleware, error handler, dll.
2. Folder `webserver` saya gunakan sebagai directory untuk menyimpan konfigurasi server, sebagai contoh adalah middleware yang mana kita dapat membagi lebih banyak lagi middleware sesuai dengan kebutuhan system kedalam folder ini. Selain itu kita juga dapat menambahkan konfigurasi error handler pada folder ini.
3. Folder `module` saya gunakan untuk menyimpan module / package untuk utilitas system, ex: Logging
4. Folder `routes` saya gunakan untuk menyimpan mapping route. Biasanya saya gunakan version number untuk mempermudah dalam development.
5. Folder `handler` saya gunakan untuk menyimpan program endpoint handler, seluruh endpoint handler terdeklarasi di dalam file `index.go`.
6. Folder `handler/database` saya gunakan untuk menyimpan program database transaction. Seluruh fungsi akan memberikan nilai balikan sesuai dengan balikan dari database.
7. Fodler `handler/useCases` saya gunakan untuk kebutuhan manipulasi data dari database, fungsi di dalam folder ini akan membaca fungsi dari folder `database` untuk mendapatkan nilai balikan database dan melakukan manipulasi.
