#!/bin/bash

# Run clients (exits after shares have been sent)
go run client.go -id 0 -val 100 &
go run client.go -id 1 -val 200 &
go run client.go -id 2 -val 300 &