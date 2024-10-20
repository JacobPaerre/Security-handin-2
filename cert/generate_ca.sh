#!/bin/sh

rm -f *.pem

# Generate CA certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DK/ST=Hovedstaden/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=jacp@itu.dk"