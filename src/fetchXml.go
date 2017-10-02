package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FetchXmlService struct{}

func (fetchXmlService *FetchXmlService) getFetchXml() (err error) {
	var GET_FETCH_XML_URL string
	if GET_FETCH_XML_URL = os.Getenv("GET_FETCH_XML_URL"); GET_FETCH_XML_URL == "" {
		return errors.New("No GET_FETCH_XML_URL was defined.")
	}

	url := GET_FETCH_XML_URL
	filepath := "./assets/fetch.xml"
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Repare request to get data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	defer out.Close()

	xml, err := (*FetchXmlService)(nil).extractFromJSON(resp.Body)
	if err != nil {
		return err
	}

	// Writer the body to file
	xmlReader := bytes.NewReader([]byte(xml))
	_, err = io.Copy(out, xmlReader)
	if err != nil {
		return err
	}

	return nil
}

func (fetchXmlService *FetchXmlService) extractFromXml(body io.ReadCloser) (fetchXml string, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	type Condition struct {
		XMLName   xml.Name `xml:"condition"`
		Attribute string   `xml:"attribute,attr"`
		Operator  string   `xml:"operator,attr"`
		Value     string   `xml:"value,attr"`
	}
	type Filter struct {
		XMLName   xml.Name    `xml:"filter"`
		Type      string      `xml:"type,attr"`
		Condition []Condition `xml:"condition"`
	}
	type Attribute struct {
		XMLName xml.Name `xml:"attribute"`
		Name    string   `xml:"name,attr"`
		Alias   string   `xml:"alias,attr"`
	}
	type Entity struct {
		XMLName    xml.Name    `xml:"entity"`
		Name       string      `xml:"name,attr"`
		Attributes []Attribute `xml:"attribute"`
		Filter     Filter      `xml:"filter"`
	}
	type FetchXml struct {
		XMLName xml.Name `xml:"fetch"`
		Entity  []Entity `xml:"entity"`
	}
	var fetchDocument FetchXml
	// fmt.Println(buf.String())

	err = xml.Unmarshal(buf.Bytes(), &fetchDocument)
	if err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
		os.Exit(1)
	}

	// for key, attr := range fetchDocument.Entity.Attributes {
	// 	fmt.Println(key)
	// 	fmt.Println(attr)
	// 	// for _, property := range attr.Properties {
	// 	// 	if property["Name"].(string) == "fmi_fetchxml" {
	// 	// 		return property["Value"].(string), nil
	// 	// 	}
	// 	// }
	// }

	return "", errors.New("No fmi_fetchxml was found.")

}

func (fetchXmlService *FetchXmlService) extractFromJSON(body io.ReadCloser) (fetchXml string, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	type Property map[string]interface{}

	type Record struct {
		Id         string     `json:"Id"`
		EntityName string     `json:"EntityName"`
		Properties []Property `json:"Properties"`
	}
	type ServiceRequestNotes struct {
		Records []Record
	}
	var serviceRequestNotes ServiceRequestNotes

	err = json.Unmarshal(buf.Bytes(), &serviceRequestNotes)
	if err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
		os.Exit(1)
	}

	for _, record := range serviceRequestNotes.Records {
		for _, property := range record.Properties {
			if property["Name"].(string) == "fmi_fetchxml" {
				return property["Value"].(string), nil
			}
		}
	}

	return "", errors.New("No fmi_fetchxml was found.")

}

func (fetchXmlService *FetchXmlService) createXmlFile(data io.ReadCloser) (err error) {
	filepath := "./assets/fetch.xml"

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Writer the body to file
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	xmlReader := bytes.NewReader(buf.Bytes())
	_, err = io.Copy(out, xmlReader)
	if err != nil {
		return err
	}

	return nil
}
