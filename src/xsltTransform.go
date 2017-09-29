package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "bytes"
    "log"
)

type entity map[string]interface{}

type XsltProc struct {}

func (xsltProc *XsltProc) transform() (result []byte) {
    xmlData, err := processXslt("stylesheet.xslt", "fetch.xml")
    if err != nil {
        fmt.Printf("ProcessXslt: %s\n", err)
        os.Exit(1)
    }
    return xmlData
}

func processXslt(xslFile string, xmlFile string) (xmlData []byte, err error) {
    cmd := exec.Command("xsltproc", "--param", "values", "'trfnumbers=trf10000,trf20000&other=o1'", xslFile, xmlFile)

      cmd.Stdin = strings.NewReader("some input")
    	var out bytes.Buffer
    	cmd.Stdout = &out
    	err = cmd.Run()
    	if err != nil {
    		log.Fatal(err)
    	}
      return out.Bytes(), err
}
