//If you change something in this file, it's only value of defaultPort.
//このファイル内で変更を行うとしたら，defaultPortの値のみです
package config

import (
	"fmt"
	"os"
	"strconv"
)

//set Const value
const (
	portKey     = "PORT"
	defaultPort = 8080
)

//get Port from os.
func Port() int {
	num, err := getInt(portKey)
	if err != nil {
		return defaultPort
	}
	return num
}

//get environment value (int) using key string.
func getInt(key string) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return 0, fmt.Errorf("config:[%s] not found", key)
	}
	num, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("config:[%s] should number", key)
	}
	return num, nil
}

//get environment value (string) using key string.
func getString(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("config:[%s] not found", key)
	}
	return v, nil
}
