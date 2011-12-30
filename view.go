package sansmagic

import "http"

type View interface {
	Route(string) bool
	Response(... string) http.Response
} 