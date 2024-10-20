#!/bin/bash

# Run clients (exits after shares have been sent)
go run client.go -id 0 &
go run client.go -id 1 &
go run client.go -id 2 &