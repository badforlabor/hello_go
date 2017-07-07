package main

import (
	"net/http"
	"net/url"
)

const cookieName = "MyCookie"

func getSession(w http.ResponseWriter, r *http.Request) string {

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}
	ret, err := url.QueryUnescape(cookie.Value)

	return ret
}
func setSession(w http.ResponseWriter, r *http.Request, session string) {
	// 不设置过期时间，这样，浏览器关闭的时候，cookie就自动失效了。
	cookie := http.Cookie{Name: cookieName, Value: url.QueryEscape(session), Path: "/", HttpOnly: true}
	http.SetCookie(w, &cookie)
}
