package main

import (
    "encoding/csv"
    "encoding/xml"
    "os"
    "fmt"
    "reflect"
)

type Catalog struct {
    Fields []struct {
        ProductList    []Product    `xml:"product,omitempty"`
    } `xml:"catalog"`
}

type Product struct {
    Fields    []Attr    `xml:",any"`
}

type Attr struct {
    XMLName    xml.Name    `xml:""`
    Value    string        `xml:",chardata"`
}

func in_slice(array interface{}, val interface{}) (exists bool) {
    exists = false

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                exists = true
                return
            }
        }
    }

    return
}

func slice_insert(original []string, position int, value string) []string {
  l := len(original)
  target := original
  if cap(original) == l {
    target = make([]string, l+1, l+10)
    copy(target, original[:position])
  } else {
    target = append(target, "")
  }
  copy(target[position+1:], original[position:])
  target[position] = value
  return target
}

func ProcessRow(row Product) map[string]string {
    var record map[string]string

    record = make(map[string]string, len(row.Fields))

    for i := range row.Fields {
        record[row.Fields[i].XMLName.Local] = row.Fields[i].Value
    }
    return record
}

func main() {
    args := os.Args[1:]
    
    if len(args)<1 {
        fmt.Println("Please provide a feed xml (<catalog><product><[any]>) as argument")
        return;        
    }
    
    xmlFile, err := os.Open(args[0])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    decoder := xml.NewDecoder(xmlFile)
    out := csv.NewWriter(os.Stdout)
    row := Product{}
    var header []string

    for token, _ := decoder.Token(); err == nil; token, err = decoder.Token() {
        if start, ok := token.(xml.StartElement); ok && start.Name.Local == "product" {
            row = Product{}
            decoder.DecodeElement(&row, &start)
            for i := range row.Fields {
                if !in_slice(header,row.Fields[i].XMLName.Local) {
                    header = slice_insert(header,i,row.Fields[i].XMLName.Local)
                }
            }
        }
    }
    out.Write(header)
    out.Flush()

    var record []string

    xmlFile, err = os.Open(args[0])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    decoder2 := xml.NewDecoder(xmlFile)

    for token, _ := decoder2.Token(); err == nil; token, err = decoder2.Token() {
        if start, ok := token.(xml.StartElement); ok && start.Name.Local == "product" {
        
            row = Product{}
            
            decoder2.DecodeElement(&row, &start)
            
            data := ProcessRow(row)
            
            if cap(record) < len(header) {
                record = make([]string, len(header))
            }
            
            record = record[:len(header)]

            for index,name := range header {
                record[index] = data[name]
            }

            out.Write(record)
            row.Fields = row.Fields[:0]
        }
    }
    out.Flush()
}
