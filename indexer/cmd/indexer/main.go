package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marutaku/epub-index-creator/indexer/internal/domain"
	"github.com/marutaku/epub-index-creator/indexer/internal/expand"
	"github.com/urfave/cli/v2"
)

func createIndexFromEPub(epubFilePath string) error {
	fmt.Printf("epubFilePath: %s\n", epubFilePath)
	tempDir, err := expand.UnzipEPub(epubFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("tempDir: %s\n", tempDir)
	opfFilePath, err := expand.FindOPFFilePath(tempDir)
	if err != nil {
		return err
	}
	fmt.Printf("opfFilePath: %s\n", opfFilePath)
	book, err := domain.NewBookFromOPF(opfFilePath)
	if err != nil {
		return err
	}
	for _, page := range book.Pages {
		keywords, err := page.ExtractKeywords()
		if err != nil {
			return err
		}
		title, err := page.Title()
		if err != nil {
			return err
		}
		fmt.Printf("title: %s, keywords: %v\n", title, keywords)
	}
	return nil
}

func main() {
	app := cli.App{
		Name: "EPub Indexer",
		Action: func(cCtx *cli.Context) error {
			epubFilePath := cCtx.Args().First()
			return createIndexFromEPub(epubFilePath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
