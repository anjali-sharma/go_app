//Redis Handler

package  main

import ( 
"encoding/json"
"fmt"
"net"
"os"
"bytes"
"github.com/xuyu/goredis"
)

const ( 
    CONN_HOST = ""
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type Response struct {
   ReqId int `json:id`
   Payload string `json:payload`
}

func main() {

    listen, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
    
    if err != nil {
        fmt.Println("Error in Listening")
        os.Exit(1)
    }
    
    // close listner when server down
    defer listen.Close()

    fmt.Println("Listening on " + CONN_PORT)

    for  {
        conn, error := listen.Accept()
        if error  != nil {
            fmt.Println("Error Acception failed")
            os.Exit(1)
        }

         // create new routine to handle connection
        go handleRequest(conn)
    }
}


func handleRequest(conn net.Conn) {
    buf := make([]byte, 1024)
    reqLen, err := conn.Read(buf)

    n := bytes.Index(buf, []byte{0})
    if err != nil {
        fmt.Println("Error reading:", err)
    }

    fmt.Println(reqLen)

    GetFromJson(buf, reqLen)
    
    message := string(buf[:n-1])

    conn.Write([]byte(message))
    conn.Close()
}



func storeinRedis(res Response){
    client, _ := goredis.Dial(&goredis.DialConfig{Address: "127.0.0.1:6379"})

    res2B, _ := json.Marshal(res)
    fmt.Println("res2B", string(res2B))

    reply, _ := client.LPush("Rqueue", string(res2B))

    fmt.Println("Reply: ", reply)

    /*
    fmt.Println("Stored to Redis Hash...")
    fmt.Println(err)
    fmt.Println(err2)
    fmt.Println(reply) */
}


func GetFromJson(buf[] byte,n int) {
    var res Response
    err := json.Unmarshal(buf[:n-1], &res)

    fmt.Println(err)
    fmt.Println("ReqId: ",res.ReqId)
    fmt.Println("Payload: ", res.Payload)

    storeinRedis(res)

}
