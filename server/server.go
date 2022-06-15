package main


import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "strconv"
    "math/rand"
    "os"
    "time"
)


func handleConnection(c net.Conn) {
    fmt.Printf("Serving %s\n", c.RemoteAddr().String())
    for {
        netData, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            return
        }

        temp := strings.TrimSpace(string(netData))
        if temp == "STOP" {
            break
        }

        //result := strconv.Itoa(random()) + "\n"
        result := strconv.Itoa(4) + "\n"
        c.Write([]byte(string(result)))
    }
    c.Close()
}


func main() {
    fmt.Printf("Starting sirkel server on port 8009")
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }

    PORT := ":" + arguments[1]
    l, err := net.Listen("tcp4", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer l.Close()
    rand.Seed(time.Now().Unix())

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        go handleConnection(c)
    }
}
