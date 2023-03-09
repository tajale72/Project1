package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	// Load the CA certificate
	caCert, err := ioutil.ReadFile("/path/to/ca.crt")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Load the client certificate and key
	cert, err := tls.LoadX509KeyPair("/path/to/client.crt", "/path/to/client.key")
	if err != nil {
		return nil, err
	}

	// Create a TLS configuration with the CA certificate and client certificate/key
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	// Create the Redis client with the TLS configuration
	client := redis.NewClient(&redis.Options{
		Addr:      "localhost:6379",
		Password:  "mypassword",
		TLSConfig: tlsConfig,
	})

	// Ping the Redis server to test the connection
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	client, err := NewRedisClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
