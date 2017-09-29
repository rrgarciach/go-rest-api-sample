package main

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "bytes"
    "log"

)

type document map[string]interface{}

type XsltTransform struct {}

func (trans *XsltTransform) load() {
    jsonData, err := processXslt("stylesheet.xslt", "fetch.xml")
    if err != nil {
        fmt.Printf("ProcessXslt: %s\n", err)
        os.Exit(1)
    }

    documents := struct {
        Deals []document
    }{}

    err = json.Unmarshal(jsonData, &documents)
    if err != nil {
        fmt.Printf("Unmarshal: %s\n", err)
        os.Exit(1)
    }

    fmt.Printf("Deals: %d\n\n", len(documents.Deals))

    for _, deal := range documents.Deals {
        fmt.Printf("DealId: %d\n", int(deal["dealid"].(float64)))
        fmt.Printf("Title: %s\n\n", deal["title"].(string))
    }
}

func processXslt(xslFile string, xmlFile string) (jsonData []byte, err error) {
    // xsltproc -o result.xml --param values 'trfnumbers=trf10000,trf20000&other=o1' stylesheet.xslt fetch.xml
    // cmd := exec.Cmd{
    //     Args: []string{"xsltproc", "--stringparam", "values", "'trfnumbers=trf10000,trf20000&other=o1'", "-o", "result.xml", xslFile, xmlFile},
    //     Env: os.Environ(),
    //     Path: "xsltproc",
    // }
    cmd := exec.Command("xsltproc", "--param", "values", "'trfnumbers=trf10000,trf20000&other=o1'", "-o", "result.xml", xslFile, xmlFile)

      cmd.Stdin = strings.NewReader("some input")
    	var out bytes.Buffer
    	cmd.Stdout = &out
    	err = cmd.Run()
    	if err != nil {
    		log.Fatal(err)
    	}
    	fmt.Printf("in all caps: %q\n", out.String())

      return out.Bytes(), err

    // jsonString, err := cmd.Output()
    // if err != nil {
    //     return jsonData, err
    // }
    //
    // fmt.Printf("%s\n", jsonString)
    //
    // jsonData = []byte(jsonString)
    //
    // return jsonData, err
}
