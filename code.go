import (
    "encoding/json"
    "errors"
    "fmt"
)

type Document struct {
    Header      string   `json:"header"`
    LineItems   []string `json:"line_items"`
}

func ProcessDocument(jsonData []byte) (bool, error) {
    var doc Document

    err := json.Unmarshal(jsonData, &doc)
    if err != nil {
        logError("Invalid JSON data")
        return false, err
    }

    if doc.Header == "" || len(doc.LineItems) == 0 {
        logError("Missing header or line items")
        return false, errors.New("validation error")
    }

    err = writeToDatabase(doc)
    if err != nil {
        logError("Database write error")
        return false, err
    }

    logInfo("Document processed successfully")
    return true, nil
}

// Метод внешенего пакета
