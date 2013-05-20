package main

import (
	"flag"
)

func main() {
	flag.Parse()
	addr := flag.String("addr", "127.0.0.1:8080", "http service address")

	config := ConfigServer{
		oauthKey:      "M9MHfTfKDyF5yZM6xueTxg",
		oauthSecret:   "1lClcicoUNEKA1pycLLO0Jruo0NA2AgK3KhLFY4jo",
		oauthCallback: "http://127.0.0.1:8080/auth/twitter/callback",

		storeSecret: "foobarquxsecret",
	}

	server := NewServer(&config)
	server.Run(*addr)
}
