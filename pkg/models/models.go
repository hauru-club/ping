// Package models contains data containers for
// server and cli part of ping project.
package models

// Packet contains some data that is contained
// by ICMP and IPv4 packets.
type Packet struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
	Seq int    `json:"seq"`
	Len int    `json:"len"`
}
