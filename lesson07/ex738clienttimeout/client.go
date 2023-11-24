package main
import (
    "fmt"
    "net/http"
    "io"
    "os"
    "time"
)
func main() { 
    client := http.Client{
        Timeout: 6 * time.Second,
    } 
    resp, err := client.Get("http://httpbin.org/get") 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close() 
    io.Copy(os.Stdout, resp.Body)
}