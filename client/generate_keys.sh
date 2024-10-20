#!/bin/sh

rm -f *.pem

# Generate key for the 3 clients & and a certificate request for each
openssl req -newkey rsa:4096 -nodes -keyout 0-key.pem -out 0-req.pem \
    -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=jacp@itu.dk"
openssl req -newkey rsa:4096 -nodes -keyout 1-key.pem -out 1-req.pem \
    -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=jacp@itu.dk"
openssl req -newkey rsa:4096 -nodes -keyout 2-key.pem -out 2-req.pem \
    -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=jacp@itu.dk"

# Use CA's private key to sign the clients' certificate request
openssl x509 -req -in 0-req.pem -days 60 -CA ../cert/ca-cert.pem -CAkey ../cert/ca-key.pem -CAcreateserial -out 0-cert.pem -extfile ../cert/server-ext.cnf
openssl x509 -req -in 1-req.pem -days 60 -CA ../cert/ca-cert.pem -CAkey ../cert/ca-key.pem -CAcreateserial -out 1-cert.pem -extfile ../cert/server-ext.cnf
openssl x509 -req -in 2-req.pem -days 60 -CA ../cert/ca-cert.pem -CAkey ../cert/ca-key.pem -CAcreateserial -out 2-cert.pem -extfile ../cert/server-ext.cnf

# Remove req-file
rm -f *-req.pem