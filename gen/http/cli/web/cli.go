// Code generated by goa v3.15.0, DO NOT EDIT.
//
// web HTTP client CLI support package
//
// Command:
// $ goa gen github.com/marutaku/epub-index-creator/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	epubindexcreatorc "github.com/marutaku/epub-index-creator/gen/http/epub_index_creator/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `epub-index-creator (list-books|find-book|create-book|update-book|delete-book|list-pages|find-page|create-page|update-page|delete-page|list-keywords-in-page)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` epub-index-creator list-books --limit 32 --offset 907143555073089594` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		epubIndexCreatorFlags = flag.NewFlagSet("epub-index-creator", flag.ContinueOnError)

		epubIndexCreatorListBooksFlags      = flag.NewFlagSet("list-books", flag.ExitOnError)
		epubIndexCreatorListBooksLimitFlag  = epubIndexCreatorListBooksFlags.String("limit", "100", "")
		epubIndexCreatorListBooksOffsetFlag = epubIndexCreatorListBooksFlags.String("offset", "", "")

		epubIndexCreatorFindBookFlags    = flag.NewFlagSet("find-book", flag.ExitOnError)
		epubIndexCreatorFindBookIsbnFlag = epubIndexCreatorFindBookFlags.String("isbn", "REQUIRED", "")

		epubIndexCreatorCreateBookFlags    = flag.NewFlagSet("create-book", flag.ExitOnError)
		epubIndexCreatorCreateBookBodyFlag = epubIndexCreatorCreateBookFlags.String("body", "REQUIRED", "")

		epubIndexCreatorUpdateBookFlags    = flag.NewFlagSet("update-book", flag.ExitOnError)
		epubIndexCreatorUpdateBookBodyFlag = epubIndexCreatorUpdateBookFlags.String("body", "REQUIRED", "")
		epubIndexCreatorUpdateBookIsbnFlag = epubIndexCreatorUpdateBookFlags.String("isbn", "REQUIRED", "ISBN of the book")

		epubIndexCreatorDeleteBookFlags    = flag.NewFlagSet("delete-book", flag.ExitOnError)
		epubIndexCreatorDeleteBookIsbnFlag = epubIndexCreatorDeleteBookFlags.String("isbn", "REQUIRED", "")

		epubIndexCreatorListPagesFlags      = flag.NewFlagSet("list-pages", flag.ExitOnError)
		epubIndexCreatorListPagesIsbnFlag   = epubIndexCreatorListPagesFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorListPagesLimitFlag  = epubIndexCreatorListPagesFlags.String("limit", "100", "")
		epubIndexCreatorListPagesOffsetFlag = epubIndexCreatorListPagesFlags.String("offset", "", "")

		epubIndexCreatorFindPageFlags      = flag.NewFlagSet("find-page", flag.ExitOnError)
		epubIndexCreatorFindPageIsbnFlag   = epubIndexCreatorFindPageFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorFindPagePageIDFlag = epubIndexCreatorFindPageFlags.String("page-id", "REQUIRED", "")

		epubIndexCreatorCreatePageFlags      = flag.NewFlagSet("create-page", flag.ExitOnError)
		epubIndexCreatorCreatePageBodyFlag   = epubIndexCreatorCreatePageFlags.String("body", "REQUIRED", "")
		epubIndexCreatorCreatePageIsbnFlag   = epubIndexCreatorCreatePageFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorCreatePagePageIDFlag = epubIndexCreatorCreatePageFlags.String("page-id", "REQUIRED", "")

		epubIndexCreatorUpdatePageFlags      = flag.NewFlagSet("update-page", flag.ExitOnError)
		epubIndexCreatorUpdatePageBodyFlag   = epubIndexCreatorUpdatePageFlags.String("body", "REQUIRED", "")
		epubIndexCreatorUpdatePageIsbnFlag   = epubIndexCreatorUpdatePageFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorUpdatePagePageIDFlag = epubIndexCreatorUpdatePageFlags.String("page-id", "REQUIRED", "")

		epubIndexCreatorDeletePageFlags      = flag.NewFlagSet("delete-page", flag.ExitOnError)
		epubIndexCreatorDeletePageIsbnFlag   = epubIndexCreatorDeletePageFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorDeletePagePageIDFlag = epubIndexCreatorDeletePageFlags.String("page-id", "REQUIRED", "")

		epubIndexCreatorListKeywordsInPageFlags      = flag.NewFlagSet("list-keywords-in-page", flag.ExitOnError)
		epubIndexCreatorListKeywordsInPageIsbnFlag   = epubIndexCreatorListKeywordsInPageFlags.String("isbn", "REQUIRED", "")
		epubIndexCreatorListKeywordsInPagePageIDFlag = epubIndexCreatorListKeywordsInPageFlags.String("page-id", "REQUIRED", "")
	)
	epubIndexCreatorFlags.Usage = epubIndexCreatorUsage
	epubIndexCreatorListBooksFlags.Usage = epubIndexCreatorListBooksUsage
	epubIndexCreatorFindBookFlags.Usage = epubIndexCreatorFindBookUsage
	epubIndexCreatorCreateBookFlags.Usage = epubIndexCreatorCreateBookUsage
	epubIndexCreatorUpdateBookFlags.Usage = epubIndexCreatorUpdateBookUsage
	epubIndexCreatorDeleteBookFlags.Usage = epubIndexCreatorDeleteBookUsage
	epubIndexCreatorListPagesFlags.Usage = epubIndexCreatorListPagesUsage
	epubIndexCreatorFindPageFlags.Usage = epubIndexCreatorFindPageUsage
	epubIndexCreatorCreatePageFlags.Usage = epubIndexCreatorCreatePageUsage
	epubIndexCreatorUpdatePageFlags.Usage = epubIndexCreatorUpdatePageUsage
	epubIndexCreatorDeletePageFlags.Usage = epubIndexCreatorDeletePageUsage
	epubIndexCreatorListKeywordsInPageFlags.Usage = epubIndexCreatorListKeywordsInPageUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "epub-index-creator":
			svcf = epubIndexCreatorFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "epub-index-creator":
			switch epn {
			case "list-books":
				epf = epubIndexCreatorListBooksFlags

			case "find-book":
				epf = epubIndexCreatorFindBookFlags

			case "create-book":
				epf = epubIndexCreatorCreateBookFlags

			case "update-book":
				epf = epubIndexCreatorUpdateBookFlags

			case "delete-book":
				epf = epubIndexCreatorDeleteBookFlags

			case "list-pages":
				epf = epubIndexCreatorListPagesFlags

			case "find-page":
				epf = epubIndexCreatorFindPageFlags

			case "create-page":
				epf = epubIndexCreatorCreatePageFlags

			case "update-page":
				epf = epubIndexCreatorUpdatePageFlags

			case "delete-page":
				epf = epubIndexCreatorDeletePageFlags

			case "list-keywords-in-page":
				epf = epubIndexCreatorListKeywordsInPageFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "epub-index-creator":
			c := epubindexcreatorc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list-books":
				endpoint = c.ListBooks()
				data, err = epubindexcreatorc.BuildListBooksPayload(*epubIndexCreatorListBooksLimitFlag, *epubIndexCreatorListBooksOffsetFlag)
			case "find-book":
				endpoint = c.FindBook()
				data, err = epubindexcreatorc.BuildFindBookPayload(*epubIndexCreatorFindBookIsbnFlag)
			case "create-book":
				endpoint = c.CreateBook()
				data, err = epubindexcreatorc.BuildCreateBookPayload(*epubIndexCreatorCreateBookBodyFlag)
			case "update-book":
				endpoint = c.UpdateBook()
				data, err = epubindexcreatorc.BuildUpdateBookPayload(*epubIndexCreatorUpdateBookBodyFlag, *epubIndexCreatorUpdateBookIsbnFlag)
			case "delete-book":
				endpoint = c.DeleteBook()
				data, err = epubindexcreatorc.BuildDeleteBookPayload(*epubIndexCreatorDeleteBookIsbnFlag)
			case "list-pages":
				endpoint = c.ListPages()
				data, err = epubindexcreatorc.BuildListPagesPayload(*epubIndexCreatorListPagesIsbnFlag, *epubIndexCreatorListPagesLimitFlag, *epubIndexCreatorListPagesOffsetFlag)
			case "find-page":
				endpoint = c.FindPage()
				data, err = epubindexcreatorc.BuildFindPagePayload(*epubIndexCreatorFindPageIsbnFlag, *epubIndexCreatorFindPagePageIDFlag)
			case "create-page":
				endpoint = c.CreatePage()
				data, err = epubindexcreatorc.BuildCreatePagePayload(*epubIndexCreatorCreatePageBodyFlag, *epubIndexCreatorCreatePageIsbnFlag, *epubIndexCreatorCreatePagePageIDFlag)
			case "update-page":
				endpoint = c.UpdatePage()
				data, err = epubindexcreatorc.BuildUpdatePagePayload(*epubIndexCreatorUpdatePageBodyFlag, *epubIndexCreatorUpdatePageIsbnFlag, *epubIndexCreatorUpdatePagePageIDFlag)
			case "delete-page":
				endpoint = c.DeletePage()
				data, err = epubindexcreatorc.BuildDeletePagePayload(*epubIndexCreatorDeletePageIsbnFlag, *epubIndexCreatorDeletePagePageIDFlag)
			case "list-keywords-in-page":
				endpoint = c.ListKeywordsInPage()
				data, err = epubindexcreatorc.BuildListKeywordsInPagePayload(*epubIndexCreatorListKeywordsInPageIsbnFlag, *epubIndexCreatorListKeywordsInPagePageIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// epub-index-creatorUsage displays the usage of the epub-index-creator command
// and its subcommands.
func epubIndexCreatorUsage() {
	fmt.Fprintf(os.Stderr, `Service is the epub_index_creator service interface.
Usage:
    %[1]s [globalflags] epub-index-creator COMMAND [flags]

COMMAND:
    list-books: ListBooks implements ListBooks.
    find-book: FindBook implements FindBook.
    create-book: CreateBook implements CreateBook.
    update-book: UpdateBook implements UpdateBook.
    delete-book: DeleteBook implements DeleteBook.
    list-pages: ListPages implements ListPages.
    find-page: FindPage implements FindPage.
    create-page: CreatePage implements CreatePage.
    update-page: UpdatePage implements UpdatePage.
    delete-page: DeletePage implements DeletePage.
    list-keywords-in-page: ListKeywordsInPage implements ListKeywordsInPage.

Additional help:
    %[1]s epub-index-creator COMMAND --help
`, os.Args[0])
}
func epubIndexCreatorListBooksUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator list-books -limit INT -offset INT

ListBooks implements ListBooks.
    -limit INT: 
    -offset INT: 

Example:
    %[1]s epub-index-creator list-books --limit 32 --offset 907143555073089594
`, os.Args[0])
}

func epubIndexCreatorFindBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator find-book -isbn STRING

FindBook implements FindBook.
    -isbn STRING: 

Example:
    %[1]s epub-index-creator find-book --isbn "1135973318059"
`, os.Args[0])
}

func epubIndexCreatorCreateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator create-book -body JSON

CreateBook implements CreateBook.
    -body JSON: 

Example:
    %[1]s epub-index-creator create-book --body '{
      "author": "Nam quibusdam vero rem aliquam voluptatibus.",
      "isbn": "1554585076269",
      "language": "Eligendi ipsa porro.",
      "publisher": "Qui in omnis quaerat odit.",
      "title": "Eligendi placeat."
   }'
`, os.Args[0])
}

func epubIndexCreatorUpdateBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator update-book -body JSON -isbn STRING

UpdateBook implements UpdateBook.
    -body JSON: 
    -isbn STRING: ISBN of the book

Example:
    %[1]s epub-index-creator update-book --body '{
      "author": "Aut incidunt adipisci consectetur.",
      "language": "Repudiandae dolorem est eligendi sit velit.",
      "publisher": "Reiciendis est est ut nihil.",
      "title": "Repudiandae beatae et non consequatur dolore ad."
   }' --isbn "5645897679908"
`, os.Args[0])
}

func epubIndexCreatorDeleteBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator delete-book -isbn STRING

DeleteBook implements DeleteBook.
    -isbn STRING: 

Example:
    %[1]s epub-index-creator delete-book --isbn "1816204039546"
`, os.Args[0])
}

func epubIndexCreatorListPagesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator list-pages -isbn STRING -limit INT -offset INT

ListPages implements ListPages.
    -isbn STRING: 
    -limit INT: 
    -offset INT: 

Example:
    %[1]s epub-index-creator list-pages --isbn "7825822385138" --limit 72 --offset 296639033918209475
`, os.Args[0])
}

func epubIndexCreatorFindPageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator find-page -isbn STRING -page-id INT

FindPage implements FindPage.
    -isbn STRING: 
    -page-id INT: 

Example:
    %[1]s epub-index-creator find-page --isbn "7612721239821" --page-id 6033189840785181948
`, os.Args[0])
}

func epubIndexCreatorCreatePageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator create-page -body JSON -isbn STRING -page-id INT

CreatePage implements CreatePage.
    -body JSON: 
    -isbn STRING: 
    -page-id INT: 

Example:
    %[1]s epub-index-creator create-page --body '{
      "page": {
         "keywords": [
            "Introduction",
            "Chapter 1",
            "Chapter 2"
         ],
         "title": "Introduction"
      }
   }' --isbn "3961327538617" --page-id 6074730612648124555
`, os.Args[0])
}

func epubIndexCreatorUpdatePageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator update-page -body JSON -isbn STRING -page-id INT

UpdatePage implements UpdatePage.
    -body JSON: 
    -isbn STRING: 
    -page-id INT: 

Example:
    %[1]s epub-index-creator update-page --body '{
      "page": {
         "keywords": [
            "Introduction",
            "Chapter 1",
            "Chapter 2"
         ],
         "title": "Introduction"
      }
   }' --isbn "7181898111974" --page-id 1462980330098305913
`, os.Args[0])
}

func epubIndexCreatorDeletePageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator delete-page -isbn STRING -page-id INT

DeletePage implements DeletePage.
    -isbn STRING: 
    -page-id INT: 

Example:
    %[1]s epub-index-creator delete-page --isbn "Nesciunt amet unde recusandae saepe velit." --page-id 8044889701868089044
`, os.Args[0])
}

func epubIndexCreatorListKeywordsInPageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator list-keywords-in-page -isbn STRING -page-id INT

ListKeywordsInPage implements ListKeywordsInPage.
    -isbn STRING: 
    -page-id INT: 

Example:
    %[1]s epub-index-creator list-keywords-in-page --isbn "2563681766707" --page-id 1462347514823464230
`, os.Args[0])
}
