package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "os"
)
type TCPHeader struct {
    SourcePort      uint16
    DestinationPort uint16
    SequenceNumber  uint32
    AckNumber       uint32
    DataOffset      uint8
    Reserved        uint8
    Flags           uint16
    WindowSize      uint16
    Checksum        uint16
    UrgentPointer   uint16
}

func main() {
    file, err := os.Open("data.bin")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read binary data from file
    data := make([]byte, 24)
    _, err = file.Read(data)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Parse TCP header
    header, err := parseTCPHeader(data)
    if err != nil {
        fmt.Println("Error parsing TCP header:", err)
        return
    }

    // Print parsed fields
    fmt.Println("Source Port:", header.SourcePort)
    fmt.Println("Destination Port:", header.DestinationPort)
    fmt.Println("Sequence Number:", header.SequenceNumber)
    fmt.Println("Evil Bit:", header.Flags&0x01)
}

func parseTCPHeader(data []byte) (*TCPHeader, error) {
    header := TCPHeader{}
    err := binary.Read(bytes.NewReader(data), binary.BigEndian, &header)
    if err != nil {
        return nil, err
    }
    return &header, nil
}
