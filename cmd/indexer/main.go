package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marutaku/epub-index-creator/batch/database"
	"github.com/marutaku/epub-index-creator/batch/indexer"
	"github.com/urfave/cli/v2"
)

func createBookFromEpubFile(epubFilePath string) (*indexer.HTMLBook, error) {
	fmt.Printf("epubFilePath: %s\n", epubFilePath)
	tempDir, err := indexer.UnzipEPub(epubFilePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("tempDir: %s\n", tempDir)
	opfFilePath, err := indexer.FindOPFFilePath(tempDir)
	if err != nil {
		return nil, err
	}
	fmt.Printf("opfFilePath: %s\n", opfFilePath)
	book, err := indexer.NewHTMLBookFromOPF(opfFilePath)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func main() {
	app := cli.App{
		Name: "EPub Indexer",
		Action: func(cCtx *cli.Context) error {
			epubFilePath := cCtx.Args().First()
			bookDB := database.NewBookDatabase("ent.db")
			pageDB := database.NewPageDatabase("ent.db")
			keywordDB := database.NewKeywordDatabase("ent.db")
			book, err := createBookFromEpubFile(epubFilePath)
			if err != nil {
				return err
			}
			bookDomain, err := book.ToDomain()
			if err != nil {
				return err
			}
			bookDB.Save(cCtx.Context, *bookDomain)
			for index, page := range bookDomain.Pages {
				err := pageDB.Save(cCtx.Context, bookDomain, page.Title, page.Path)
				if err != nil {
					return err
				}
				keywords, err := book.Pages[index].ExtractKeywords()
				if err != nil {
					return err
				}
				for _, keyword := range keywords {
					err := keywordDB.Save(cCtx.Context, page, keyword)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
