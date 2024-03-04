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
	return `epub-index-creator (list-books|find-book)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` epub-index-creator list-books --body '{
      "limit": 62,
      "offset": 7864368622572586298
   }'` + "\n" +
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

		epubIndexCreatorListBooksFlags    = flag.NewFlagSet("list-books", flag.ExitOnError)
		epubIndexCreatorListBooksBodyFlag = epubIndexCreatorListBooksFlags.String("body", "REQUIRED", "")

		epubIndexCreatorFindBookFlags    = flag.NewFlagSet("find-book", flag.ExitOnError)
		epubIndexCreatorFindBookIsbnFlag = epubIndexCreatorFindBookFlags.String("isbn", "REQUIRED", "ISBN of the book")
	)
	epubIndexCreatorFlags.Usage = epubIndexCreatorUsage
	epubIndexCreatorListBooksFlags.Usage = epubIndexCreatorListBooksUsage
	epubIndexCreatorFindBookFlags.Usage = epubIndexCreatorFindBookUsage

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
				data, err = epubindexcreatorc.BuildListBooksPayload(*epubIndexCreatorListBooksBodyFlag)
			case "find-book":
				endpoint = c.FindBook()
				data, err = epubindexcreatorc.BuildFindBookPayload(*epubIndexCreatorFindBookIsbnFlag)
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

Additional help:
    %[1]s epub-index-creator COMMAND --help
`, os.Args[0])
}
func epubIndexCreatorListBooksUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator list-books -body JSON

ListBooks implements ListBooks.
    -body JSON: 

Example:
    %[1]s epub-index-creator list-books --body '{
      "limit": 62,
      "offset": 7864368622572586298
   }'
`, os.Args[0])
}

func epubIndexCreatorFindBookUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] epub-index-creator find-book -isbn STRING

FindBook implements FindBook.
    -isbn STRING: ISBN of the book

Example:
    %[1]s epub-index-creator find-book --isbn "Reiciendis est est ut nihil."
`, os.Args[0])
}
