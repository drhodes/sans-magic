package sansmagic

type View interface {
	Route(... string) Response
} 