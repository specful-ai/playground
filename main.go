package main

import (
    "fmt"
    "net"
    "bufio"
)

type Client struct {
    conn net.Conn
    name string
    ch   chan string
}

var (
    clients = make(map[net.Addr]Client)
    entering = make(chan Client)
    leaving = make(chan Client)
    messages = make(chan string)
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()

    fmt.Println("Chat server is running on port 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    ch := make(chan string)
    client := Client{conn: conn, name: "", ch: ch}
    addr := conn.RemoteAddr()

    clients[addr] = client
    entering <- client

    defer func() {
        delete(clients, addr)
        leaving <- client
        conn.Close()
    }()

    go clientWriter(conn, ch)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        messages <- fmt.Sprintf("%s: %s", client.name, scanner.Text())
    }
}

func clientWriter(conn net.Conn, ch chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg)
    }
}

func broadcast(msg string, sender Client) {
    for _, client := range clients {
        if client != sender {
            client.ch <- msg
        }
    }
}

func handleMessages() {
    for {
        select {
        case msg := <-messages:
            for _, client := range clients {
                client.ch <- msg
            }
        case client := <-entering:
            for _, c := range clients {
                client.ch <- fmt.Sprintf("%s has joined the chat", c.name)
            }
            clients[client.conn.RemoteAddr()] = client
            client.ch <- "Welcome to the chat room!"
        case client := <-leaving:
            for _, c := range clients {
                client.ch <- fmt.Sprintf("%s has left the chat", c.name)
            }
            close(client.ch)
        }
    }
}
