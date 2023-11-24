//go:generate stringer -type=Role
package main

type Role int

const (
    Unknown Role = iota
    Admin
    Moderator
    User
)

func main() {
    // ...
}