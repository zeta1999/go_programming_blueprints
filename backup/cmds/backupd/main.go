package main

import (
	"flag"
	"github.com/daikimiura/go_programming_blueprints/backup"
	"github.com/matryer/filedb"
	"log"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Int("interval", 10, "チェックの間隔(秒)")
		archive  = flag.String("archive", "archive", "アーカイブの保存先")
		dbpath   = flag.String("db", "./db", "filedbのデータベースパス")
	)
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}

	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}
}