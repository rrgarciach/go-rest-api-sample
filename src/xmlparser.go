package main

//
// import (
// 	"encoding/xml"
// 	"flag"
// 	"fmt"
// 	"os"
// )
//
// type Redirect struct {
// 	Title string `xml:"title,attr"`
// }
//
// type Page struct {
// 	Title string   `xml:"title"`
// 	Redir Redirect `xml:"redirect"`
// 	Text  string   `xml:"revision>text"`
// }
//
// func readXml(xmlStr string) {
// 	xml.Marshal
// 	flag.Parse()
//
// 	xmlFile, err := os.Open(*inputFile)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer xmlFile.Close()
//
// 	decoder := xmlStr.NewDecoder(xmlFile)
// 	total := 0
// 	var inElement string
// 	for {
// 		// Read tokens from the XML
// 		t, _ := decoder.Token()
// 		if t == nil {
// 			break
// 		}
// 		// Inspect the type of the token just read
// 		switch se := t.(type) {
// 		case xmlStr.StartElement:
// 			// If have just read a StartElement token
// 			inElement = se.Name.Local
//
// 			if inElement == "fetch" {
// 				var p Page
// 				// decode a whole chunk of following XML into the variable
// 				decoder.DecodeElement(&p, &se)
//
// 				// Process
// 				p.Title = CanonicalizeTitle(p.Title)
// 				m := filter.MatchString(p.Title)
// 				if !m && p.Redir.Title == "" {
// 					WritePage(p.Title, p.Text)
// 					total++
// 				}
// 			}
// 		default:
// 		}
// 	}
// }
