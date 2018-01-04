package main

import "net/http"
import "time"
import "fmt"
import "bytes"
import "encoding/binary"
import "crypto/rand"

func main() {
    http.HandleFunc("/", httpHandle)
    http.ListenAndServe("localhost:8081", nil)
}

func httpHandle(w http.ResponseWriter, r *http.Request) {
    randId, err := randUint32()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
        return
    }
    s1 := fmt.Sprintf("你好<%d>: %s\n", randId, time.Now().Format("15:04:05"))
    w.Write([]byte(s1))
    if f, ok := w.(http.Flusher); ok {
        f.Flush()
    }
    time.Sleep(2 * time.Second)
    s2 := fmt.Sprintf("世界<%d>: %s\n", randId, time.Now().Format("15:04:05"))
    w.Write([]byte(s2))
    if f, ok := w.(http.Flusher); ok {
        f.Flush()
    }
}

func randUint32() (uint32, error) {
    b1 := make([]byte, 4)
    _, err := rand.Read(b1)
    if err !=nil {
        return 0, err
    }
    var res uint32
    b2 := bytes.NewBuffer(b1)
    err = binary.Read(b2, binary.LittleEndian, &res)
    if err != nil {
        return 0, err
    }
    return res, nil
}
