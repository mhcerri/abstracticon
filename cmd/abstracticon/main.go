package main

import (
	"crypto"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "crypto/md5"

	"github.com/mhcerri/abstracticon"
)

func main() {

	hashes := map[string]crypto.Hash{
		"md5": crypto.MD5,
	}

	var hashNames []string
	for n, _ := range hashes {
		hashNames = append(hashNames, n)
	}

	var (
		hashName    = flag.String("hash", "md5", "Hash function. Options: "+strings.Join(hashNames, ", "))
		multiplier  = flag.Int("multiplier", 8, "Number of pixels that represents a point.")
		points      = flag.Int("points", 8, "Number of points in the height and width of the image.")
		transparent = flag.Bool("transparent", false, "Use transparent background.")
		mirrored    = flag.Bool("mirrored", true, "Mirror the image left to right.")
		output      = flag.String("output", "output.png", "Output file.")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Reads data from stdin and generates an icon. Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *multiplier < 0 {
		log.Fatal("Invalid multiplier.")
	}
	if *points < 0 {
		log.Fatal("Invalid points.")
	}

	attrs := abstracticon.Attrs{
		Hash:        hashes[*hashName],
		Multiplier:  *multiplier,
		Points:      *points,
		Transparent: *transparent,
		NotMirrored: !*mirrored,
	}
	data, err := ioutil.ReadAll(os.Stdin)
	img := abstracticon.RenderFromBytes(data, attrs)
	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
