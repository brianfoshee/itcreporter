package itcreporter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brianfoshee/itcreporter/properties"
)

var apiVersion = "1.0"

type Reporter struct {
	properties properties.Properties
	c          *http.Client
}

func New() Reporter {
	r := Reporter{
		c: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	r.properties = properties.NewFromFile("Reporter.properties")
	return r
}

type request struct {
	Account    string `json:"account,omitempty"`
	UserID     string `json:"userid"`
	QueryInput string `json:"queryInput"`
	Version    string `json:"version"`
	Mode       string `json:"mode"`
	Password   string `json:"password"`
}

// Command parses cmd and executes the associated function.
// TODO: figure out what this should return
func (r *Reporter) Command(cmd string) {
	// example:
	// Sales.getStatus
	// Finance.getStatus

	// cmd has to be a certain format
	// Sales.getReport 86088768, Sales, Summary, Daily, 20150201
	// turns into
	// Sales.getReport, 86088768,Sales,Summary,Daily,20150201
	cmd = strings.Replace(cmd, ",", "", -1)
	args := strings.Split(cmd, " ")
	if len(args) > 1 {
		args[0] = args[0] + " "

		cmd = strings.Join(args, ",")
		cmd = strings.Replace(cmd, " ,", ", ", 1)
	}

	var req request
	req.UserID = r.properties.UserID
	req.Password = r.properties.Password
	req.Version = apiVersion
	req.Mode = r.properties.Mode
	req.Account = r.properties.Account
	req.QueryInput = fmt.Sprintf("[p=Reporter.properties, %s]", cmd)

	// Anything starting with Sales should use the sales URL
	// Anything starting with Finance should use the finance URL
	var url string
	if strings.HasPrefix(cmd, "Sales") {
		url = r.properties.SalesURL
	} else if strings.HasPrefix(cmd, "Finance") {
		url = r.properties.FinanceURL
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatal("error marshaling req", err)
	}

	q := strings.NewReader("jsonRequest=" + string(b))
	//log.Print(q)

	resp, err := r.c.Post(url, "application/x-www-form-urlencoded", q)
	if err != nil {
		log.Fatal("error making request", err)
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") == "application/a-gzip" {
		name := resp.Header.Get("Filename")
		f, err := os.Create(name)
		if err != nil {
			log.Fatal("error creating file", err)
		}
		defer f.Close()

		if _, err := io.Copy(f, resp.Body); err != nil {
			log.Fatal("error copying body", err)
		}

		msg := resp.Header.Get("downloadmsg")
		log.Printf(msg)
	} else {
		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("couldn't read body", err)
		}

		log.Print(string(out))
	}
}

// Both 'Sales and Trends' and 'Payments and Financial' commands
func (r *Reporter) getStatus() {
}

func getAccounts() {
}

// Sales and Trends Commands
// start with Sales.[command]

// This command returns a list of vendor numbers for which you can download reports.
func getVendors() {
}

/*
This command downloads a report.
If a report is delayed, getReport will return a delay message, including an estimated time of availability, if known.
In robot mode, delays will return an error code of “117”. If an estimated time of availability is known, a “retry” value is also returned. The retry value is expressed in milliseconds, and indicates how long to wait before trying again. See below for an example.
*/
// Sales.getReport [vendor number], [report type], [report subtype], [date type], [date]
func getReport() {
}
