package main

import (
	"blog-api-lvmingyin-com/config"
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/router"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"time"
	"blog-api-lvmingyin-com/util"
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

	var cfg config.Config
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		fmt.Println("配置文件读取失败")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}

	listenPort := int64(3333)

	if cfg.ListenPort != 0 {
		listenPort = cfg.ListenPort
	}

	fs, err := os.Create(fmt.Sprintf("logs/error-%d.log", time.Now().Unix()))
	defer fs.Close()

	if err != nil {
		fmt.Println("日志文件打开失败")
		panic(err)
	}
	util.SetLog(fs)

	db.MySqlConn(&cfg.Mysql)
	defer db.Close()
	router.Start(listenPort)
}
