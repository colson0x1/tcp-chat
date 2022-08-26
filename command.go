package main 

type commandID int

const (
	CMD_USERNAME commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_EXIT
)

type command struct {
	id commandID
	client *client
	args []string
}