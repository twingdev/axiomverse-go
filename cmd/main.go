package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alecthomas/kong"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// CLI structure for commands
var CLI struct {
	Config string `help:"Path to configuration file" type:"path" short:"c" default:"config.yaml"`

	Get struct {
		Key string `arg:"" help:"Key to retrieve from etcd."`
	} `cmd:"" help:"Get a key from etcd"`

	Set struct {
		Key   string `arg:"" help:"Key to set in etcd."`
		Value string `arg:"" help:"Value to set for the key."`
	} `cmd:"" help:"Set a key-value pair in etcd"`
}

func main() {
	// Parse CLI input
	ctx := kong.Parse(&CLI)

	// Load configuration
	if err := initConfig(CLI.Config); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize etcd client
	etcdClient, err := initEtcdClient()
	if err != nil {
		log.Fatalf("Error initializing etcd client: %v", err)
	}
	defer etcdClient.Close()

	// Execute commands based on CLI input
	switch ctx.Command() {
	case "get <key>":
		getKey(etcdClient, CLI.Get.Key)
	case "set <key> <value>":
		setKey(etcdClient, CLI.Set.Key, CLI.Set.Value)
	default:
		fmt.Println("Command not recognized.")
	}
}

// Initialize Viper with config file and environment variables
func initConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv() // Automatically read environment variables

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Initialize etcd client using configurations from Viper
func initEtcdClient() (*clientv3.Client, error) {
	endpoint := viper.GetString("etcd.endpoint")
	dialTimeout := viper.GetDuration("etcd.dial_timeout")
	if dialTimeout == 0 {
		dialTimeout = 5 * time.Second
	}

	return clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: dialTimeout,
	})
}

// Get key from etcd
func getKey(client *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Fatalf("Error getting key %s: %v", key, err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
}

// Set key-value in etcd
func setKey(client *clientv3.Client, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := client.Put(ctx, key, value)
	if err != nil {
		log.Fatalf("Error setting key %s: %v", key, err)
	}
	fmt.Printf("Successfully set key %s to value %s\n", key, value)
}
