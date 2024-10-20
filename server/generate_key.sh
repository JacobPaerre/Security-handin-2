#!/bin/sh

rm -f *.pem

# Generate key for the server & and a certificate request
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem \
    -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=jacp@itu.dk"

# Use CA's private key to sign the server's certificate request
openssl x509 -req -in server-req.pem -days 60 -CA ../cert/ca-cert.pem -CAkey ../cert/ca-key.pem -CAcreateserial -out server-cert.pem -extfile ../cert/server-ext.cnf

# Remove req-file
rm -f server-req.pem