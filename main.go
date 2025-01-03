package main

import (
	"embed"

	"github.com/cksidharthan/share-secret/cmd"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	cmd.Start(frontend)
}
