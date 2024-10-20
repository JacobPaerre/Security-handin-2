package cert

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

func LoadTLSCredentials(certPath, keyPath string) (credentials.TransportCredentials, error) {
    // Load server's certificate and private key
    serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
    if err != nil {
        return nil, err
    }

    // Create the credentials and return it
    config := &tls.Config{
        Certificates: []tls.Certificate{serverCert},
        ClientAuth:   tls.NoClientCert,
    }

    return credentials.NewTLS(config), nil
}

func LoadCAcertificate(caCertPath string) (credentials.TransportCredentials, error) {
    pemCA, err := os.ReadFile(caCertPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read CA certificate: %v", err)
    }

    certPool := x509.NewCertPool()
    if !certPool.AppendCertsFromPEM(pemCA) {
        return nil, fmt.Errorf("failed to append CA certificate")
    }

    config := &tls.Config{
        RootCAs: certPool,
    }

    return credentials.NewTLS(config), nil
}