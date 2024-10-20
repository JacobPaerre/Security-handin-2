#!/bin/bash



# Run clients (runs in the background and ends when shares are send to the server)
go run client.go -id 0 &
go run client.go -id 1 &
go run client.go -id 2 &
