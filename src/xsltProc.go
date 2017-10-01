package main

import (
    "fmt"
    "os"
    "os/exec"
    "bytes"
    "log"
    "strings"
)

type entity map[string]interface{}

type XsltProc struct {}

func (xsltProc *XsltProc) transform(params map[string][]string) (result []byte) {
    stringParams := parseParams(params)
    xmlData, err := processXslt("./assets/stylesheet.xslt", "./assets/fetch.xml", stringParams)
    if err != nil {
        fmt.Printf("ProcessXslt: %s\n", err)
        os.Exit(1)
    }
    return xmlData
}

func parseParams(params map[string][]string) (stringParams string) {
  result := "'"
  for key, param := range params {
    for _, value := range param {
      result = result + string(key) + "=" + value + "&"
    }
  }
  return strings.TrimSuffix(result, "&") + "'"
}

func processXslt(xslFile string, xmlFile string, stringParams string) (xmlData []byte, err error) {
    cmd := exec.Command("xsltproc", "--param", "values", stringParams, xslFile, xmlFile)

  	var out bytes.Buffer
  	cmd.Stdout = &out
  	err = cmd.Run()
  	if err != nil {
  		log.Fatal(err)
  	}
    // fmt.Println(out.String()) // Uncomment to see transformed string output
    return out.Bytes(), err
}
