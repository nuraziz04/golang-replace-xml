package model

import "encoding/xml"

type Data struct {
	EnableBoxcarring string     `xml:"enableBoxcarring,attr"`
	Name             string     `xml:"name,attr"`
	ServiceNamespace string     `xml:"serviceNamespace,attr"`
	Transports       string     `xml:"transports,attr"`
	Configes         []Config   `xml:"config"`
	Queries          []Query    `xml:"query"`
	Resources        []Resource `xml:"resource"`
}

type Config struct {
	EnableOData string   `xml:"enableOData,attr"`
	Id          string   `xml:"id,attr"`
	Property    Property `xml:"property"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Query struct {
	XMLName            xml.Name `xml:"query"`
	ID                 string   `xml:"id,attr"`
	ReturnUpdatedCount bool     `xml:"returnUpdatedRowCount,attr"`
	UseConfig          string   `xml:"useConfig,attr"`
	SQL                string   `xml:"sql"`
	Result             Result   `xml:"result"`
	Params             []Param  `xml:"param"`
}

type Result struct {
	Element      string         `xml:"element,attr"`
	RowName      string         `xml:"rowName,attr"`
	UseColNums   bool           `xml:"useColumnNumbers,attr"`
	ElementChild []ElementChild `xml:"element"`
}

type ElementChild struct {
	Column  string `xml:"column,attr"`
	Name    string `xml:"name,attr"`
	XSDType string `xml:"xsdType,attr"`
}

type Param struct {
	Name    string `xml:"name,attr"`
	SQLType string `xml:"sqlType,attr"`
}

type Resource struct {
	Method    string    `xml:"method,attr"`
	Path      string    `xml:"path,attr"`
	CallQuery CallQuery `xml:"call-query"`
}

type CallQuery struct {
	Href      string      `xml:"href,attr"`
	WithParam []WithParam `xml:"with-param"`
}

type WithParam struct {
	Name       string `xml:"name,attr"`
	QueryParam string `xml:"query-param,attr"`
}

type Queries struct {
	XMLName xml.Name `xml:"root"`
	Data    Data     `xml:"data"`
}
