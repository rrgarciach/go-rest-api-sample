package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FetchXml struct{}

func (fetchXml *FetchXml) getFetchXml() (err error) {
	url := "https://devapi.foundationmedicine.com/FMI.CRM.Integration.Query/api/query/getFetchXml?functionName=xgetservicerequestnotes"
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

	xml, err := extractFetchXml(resp.Body)
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

func extractFetchXml(body io.ReadCloser) (fetchXml string, err error) {
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
