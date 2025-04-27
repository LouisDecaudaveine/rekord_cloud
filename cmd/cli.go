package main

import (
	// "embed"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	databaserc "github.com/LouisDecaudaveine/rekord_cloud/internal/database"
	"github.com/LouisDecaudaveine/rekord_cloud/internal/parser"
	"github.com/LouisDecaudaveine/rekord_cloud/internal/utils"
)

// var testDataFS embed.FS

func main() {
	fmt.Println("CLI command line interface")

	path := filepath.Clean("./test-data/test-rekordbox-export.xml")
	rawXML, err := os.ReadFile(path)
	utils.Check(err)

	var playlists parser.DJPlaylists
	err = xml.Unmarshal(rawXML, &playlists)
	utils.Check(err)

	parser.PrintAllParsedFile(playlists)

	db, err := databaserc.InitDB()
	utils.Check(err)
	defer db.Close()
}
