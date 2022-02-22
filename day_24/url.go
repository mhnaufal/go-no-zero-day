package day_24

import (
	"fmt"
	"net/url"
)

func Url() {
	/// Parsing URL string into url.URL type will gave us more information
	/// regarding to the url for such as protocol, path, query, etc

	var urlString = "https://github.com/mhnaufal/go-no-zero-day?file=var"
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("URL : %s\n", urlString)

	fmt.Printf("Protocol : %s\n", u.Scheme)
	fmt.Printf("Host : %s\n", u.Host)
	fmt.Printf("Path : %s\n", u.Path)

	var file = u.Query()["file"][0]
	fmt.Printf("File : %s\n", file)	
}