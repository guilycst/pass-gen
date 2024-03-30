package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"

	"github.com/guilycst/pass-gen/pkg"
)

type Opts struct {
	passwordLength *int
	useSymbols     bool
	useDigits      bool
	useLetters     bool
	batchSize      *int
	omitLf         *bool
	hash           *bool
}

var options = Opts{}

func init() {
	options.omitLf = flag.Bool("n", false, "do not output the trailing newline, ignored if -b bigger than one")
	options.hash = flag.Bool("h", false, "output sha256sum hash of the password")
	options.passwordLength = flag.Int("l", 20, "password length")
	options.batchSize = flag.Int("b", 1, "number of passwords generated, ignores option -n if batch size bigger than one")
	use := flag.String("u", "lsd", "[l][s][d] use letters and/or symbols and/or digits")
	flag.Parse()

	if len(*use) > 3 || len(*use) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if (*options.batchSize) > 1 {
		lf := false
		options.omitLf = &lf
	}

	uses := []rune(*use)
	for _, u := range uses {
		switch u {
		case 'l':
			options.useLetters = true
		case 's':

			options.useSymbols = true
		case 'd':

			options.useDigits = true
		default:
			fmt.Printf("%q not accepted on option -u\nvalid values are: l, s and d\n", u)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
}

func main() {
	for range *options.batchSize {
		pswd := pkg.Generate(*options.passwordLength, options.useLetters, options.useSymbols, options.useDigits)
		if *options.hash {
			pswd = fmt.Sprintf("%s - %x", pswd, sha256.Sum256([]byte(pswd)))
		}
		print(pswd)
	}
}

func print(s string) {
	if *options.omitLf {
		fmt.Printf(s)
	} else {
		fmt.Println(s)
	}
}
