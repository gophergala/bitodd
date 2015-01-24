package model

import ()

type Envelope struct {
	Action string `json:"action"`
	Info   *Info  `json:"hello"`
}

// Info action
var INFO = "INFO"

// Info payload
type Info struct {
	UserCount int `json:"user_count"`
}
