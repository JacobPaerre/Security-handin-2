# Initial setup
Before running the program all certificates has to be made. There will be 1 certificate for a **CA** which will be used to sign the **servers** and **clients**' certificate. To create all certificates, do:
```
make cert
```

# How to run
There are 2 ways to run the application:

1. Using 2 terminals: 1 for the server and 1 for all the clients.
2. Using 4 terminals: 1 for the server and 3 for the clients.

To start the server (from room-directory):
```
make startServer
```

To have the all the clients running in the same terminal:
```
make startClient
```

To have all the clients run in a terminal of its own:
```
go run client/client.go -id x
```
Where _x_ has to be 0, 1 and 2. The clients must be run within 10 seconds of the first starting or the program will fail.