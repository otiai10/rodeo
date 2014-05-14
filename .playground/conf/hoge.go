package main

import "github.com/robfig/config"
import "fmt"

func main() {
	c, _ := config.ReadDefault("my.conf")
	b, _ := c.Bool("hoge", "fuga")
	fmt.Println(b)
}
