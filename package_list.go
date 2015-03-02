package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath"
	"errors"
	"container/list"
	"math/rand"
	"sort"
	"hash/crc32"
	"crypto/sha1"
	"encoding/gob"
	"net"
	"net/http"
	"net/rpc"
	"io"
	"flag"
	"sync"
	"time"
)
//http
func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    io.WriteString(
        res,
        `<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
    )
}
func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":9000", nil)
}

func server() {
    // listen on a port
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        // accept a connection
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        // handle the connection
        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn) {
    // receive the message
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received", msg)
    }

    c.Close()
}

func client() {
    // connect to the server
    c, err := net.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    // send the message
    msg := "Hello World"
    fmt.Println("Sending", msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }

    c.Close()
}

func main() {
    fmt.Println(
        // true
        strings.Contains("test", "es"),

        // 2
        strings.Count("test", "t"),

        // true
        strings.HasPrefix("test", "te"),

        // true
        strings.HasSuffix("test", "st"),

        // 1
        strings.Index("test", "e"),

        // "a-b"
        strings.Join([]string{"a","b"}, "-"),

        // == "aaaaa"
        strings.Repeat("a", 5),

        // "bbaa"
        strings.Replace("aaaa", "a", "b", 2),

        // []string{"a","b","c","d","e"}
        strings.Split("a-b-c-d-e", "-"),

        // "test"
        strings.ToLower("TEST"),

        // "TEST"
        strings.ToUpper("test"),

    )


		//OS

		file, err := os.Open("test.txt")
	if err != nil {
			// handle the error here
			return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
			return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
			return
	}

	str := string(bs)
	fmt.Println(str)


	//io/util

	bs, err := ioutil.ReadFile("test.txt")
if err != nil {
		return
}
str := string(bs)
fmt.Println(str)


// filepath

filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
    })

		// erros

		err := errors.New("error message")


		//sort

		type Person struct {
    Name string
    Age int
}

type ByName []Person

func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func main() {
    kids := []Person{
        {"Jill",9},
        {"Jack",10},
    }
    sort.Sort(ByName(kids))
    fmt.Println(kids)
}

//hash

h := crc32.NewIEEE()
h.Write([]byte("test"))
v := h.Sum32()
fmt.Println(v)

// crypt

h := sha1.New()
h.Write([]byte("test"))
bs := h.Sum([]byte{})
fmt.Println(bs)
}

//encoding gob

go server()
go client()

var input string
fmt.Scanln(&input)

//math rand

// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
