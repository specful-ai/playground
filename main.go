package main

import (
    "database/sql"
    "encoding/csv"
    "flag"
    "fmt"
    "net/http"
    "os"
   "sync/atomic"
   "time"
   _ "github.com/mattn/go-sqlite3"
)

var counter uint64

func setupDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./requests.db")
    if err != nil {
        return nil, err
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, count INTEGER)")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func main() {
    port := flag.String("port", "8080", "port to listen on")
    flag.Parse()

    db, err := setupDB()
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    go func() {
        for {
            <-time.After(time.Minute)
            count := atomic.LoadUint64(&counter)
            fmt.Printf("Total requests received: %d\n", count)
            _, err := db.Exec("INSERT INTO requests (timestamp, count) VALUES (CURRENT_TIMESTAMP, ?)", count)
            if err != nil {
                fmt.Println(err)
            }
        }
    }()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request:")
        writer := csv.NewWriter(os.Stdout)
        defer writer.Flush()
        writer.Write([]string{"Header", "Value"})
        for k, v := range r.Header {
            for _, val := range v {
                writer.Write([]string{k, val})
            }
        }
        count := atomic.AddUint64(&counter, 1)
        fmt.Printf("Request count: %d\n", count)
    })

    addr := fmt.Sprintf(":%s", *port)
    http.ListenAndServe(addr, nil)
}
