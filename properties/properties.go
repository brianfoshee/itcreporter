// TODO update this documentation for Reporter
// Package properties defines a type to store an email and password, and methods to
// extract that information from a properly formatted .properties file.
package properties

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Properties represents an email and a password, typically coming from a file
// formatted like autoingestion.properties.example
type Properties struct {
	UserID     string
	Password   string
	Account    string
	Mode       string
	SalesURL   string
	FinanceURL string
}

// NewPropertiesFromFile receives a file name, extracts a username and
// password from it, then returns an instance of Properties containtaing those
// attributes. The argument must be a filesystem path.
// TODO return an error too
func NewFromFile(name string) Properties {
	p := Properties{}
	p.fromFile(name)
	return p
}

func (p *Properties) fromFile(name string) {
	str, err := p.readFile(name)
	if err != nil {
		return
	}
	fileParts := strings.Split(str, "\n")

	for _, part := range fileParts {
		if part == "" {
			continue
		}
		arr := strings.Split(part, "=")
		if len(arr) != 2 {
			fmt.Println("bad part", part)
			continue
		}
		switch arr[0] {
		case "UserId":
			p.UserID = arr[1]
		case "Password":
			p.Password = arr[1]
		case "Mode":
			p.Mode = arr[1]
		case "Account":
			p.Account = arr[1]
		case "SalesUrl":
			p.SalesURL = arr[1]
		case "FinanceUrl":
			p.FinanceURL = arr[1]
		}
	}
}

func (p *Properties) readFile(fn string) (string, error) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("error opening properties file", err)
		return "", err
	}
	defer file.Close()

	// Get the size of the file
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("could not read file stats", err)
		return "", err
	}
	nb := fi.Size()

	// Read the file into memory
	b := make([]byte, nb)
	rb, err := file.Read(b)
	if err != nil {
		fmt.Println("could not read file", err)
		return "", err
	}
	if int64(rb) != nb {
		fmt.Println("did not read in as many bytes as size of file")
		return "", errors.New("did not read in as many bytes as size of file")
	}

	return strings.TrimSpace(string(b)), nil
}
