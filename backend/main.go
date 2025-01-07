package main

import (
	"github.com/cksidharthan/ghost-send/cmd"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Start()
}
