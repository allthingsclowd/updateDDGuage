package main

import (
	"flag"
	datadog "github.com/allthingsclowd/datadoghelper"
)

func main() {
	filePtr := flag.String("file", "metric.json", "Default's to using configuration file metric.json. Use -file=<filename> to use an alternative configuration file")
	flag.Parse()
	datadog.UpdateDataDogGuagefromFile("webpageapp", "backendTotal", *filePtr )
}