package main

import "fmt"

type Status uint8

const (
	Created Status = iota
	Pending
	Approved
	Rejected
)

func (s Status) String() string {
	switch s {
	case Created:
		return "created"
	case Pending:
		return "pending"
	case Approved:
		return "approved"
	case Rejected:
		return "rejected"
	}
	return "unknown"
}
func main() {
	fmt.Println(Pending)
}
