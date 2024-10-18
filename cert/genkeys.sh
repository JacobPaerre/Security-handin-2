#!/bin/bash

rm *.pem
rm *.srl

# Generate CA private key and self-signed certificate
for i in 1 2 3; do
    openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key${i}.pem -out ca-cert${i}.pem -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/CN=*.itu.dk/emailAddress=jacp@itu.dk"
done

# Generate web server's private key and CSR
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/CN=*.itu.dk/emailAddress=jacp1@itu.dk"

# Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -CA ca-cert1.pem -CAkey ca-key1.pem -CAcreateserial -out server-cert.pem -days 365