package main

import (
    "io"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/111", setCookie)
    http.HandleFunc("/222", setCookie2)
    http.ListenAndServe(":8080", nil)
}

// 设置和读取cookie的方法
func setCookie(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name: "mycookie",
        Value: "hello",
        Path: "/",
        Domain: "localhost",
        MaxAge: 120,
    }
    // 使用http.SetCookie方法设置Cookie
    http.SetCookie(w, cookie)

    ck, err := r.Cookie("mycookie")
    if err != nil {
        io.WriteString(w, err.Error())
        return
    }
    io.WriteString(w, ck.Value)
}

/**
响应头方式设置
*/
func setCookie2(w http.ResponseWriter, r *http.Request) {
    ck := &http.Cookie{
        Name:   "myCookie3",
        Value:  "heyguys",
        Path:   "/",
        Domain: "localhost",
        MaxAge: 120,
    }
    // 使用响应头的方式设置
    w.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "%%", -1))

    ck2, err := r.Cookie("myCookie3")
    if err != nil {
        // 错误表示 Cookie 不存在
        io.WriteString(w, err.Error())
        return
    }
    io.WriteString(w, ck2.Value)
}
