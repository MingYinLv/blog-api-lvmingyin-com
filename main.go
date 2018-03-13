package main

import (
	"blog-api-lvmingyin-com/config"
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/router"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	fileName := "config.development.toml"
	if len(args) > 0 {
		fileName = args[0]
	}
	fmt.Printf("读取配置文件：%s\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("配置文件读取失败")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}

	var config config.Config
	if _, err := toml.Decode(string(data), &config); err != nil {
		fmt.Println("配置文件读取失败")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}

	listenPort := int64(3333)

	if config.ListenPort != 0 {
		listenPort = config.ListenPort
	}

	fmt.Println(config)

	db.MySqlConn(&config.Mysql)
	router.Start(listenPort)
}
