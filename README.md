[![Go Report Card](https://goreportcard.com/badge/github.com/flannel-dev-lab/ANSI-378-Decoder)](https://goreportcard.com/report/github.com/flannel-dev-lab/ANSI-378-Decoder)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e06105c32cf6440783c256085b708d46)](https://www.codacy.com/manual/vmanikes/ANSI-378-Decoder?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=flannel-dev-lab/ANSI-378-Decoder&amp;utm_campaign=Badge_Grade)
# ANSI-378 Decoder

Breaks down a ANSI INCITS 378-2004 Finger Minutiae Format

## Information

Supports go versions > 1.12.1. To install add this line to your `go.mod` file 

`github.com/flannel-dev-lab/ANSI-378-Decoder`

-   **Get Record headers**
    ```go
    RecordHeaders(fmd []byte) (*RecordHeader, error)
    ```
    Takes in a `fmd` byte array which is an array of bytes and returns the address of `RecordHeader`
    
    The following headers are returned
    ```go
    type RecordHeader struct {
    	FormatIdentifier    string
    	VersionNumber       string
    	RecordLength        uint16
    	CBEFFOwner          uint16
    	CBEFFType           uint16
    	EquipmentCompliance uint16
    	EquipmentID         uint16
    	ImageSizeX          uint16
    	ImageSizeY          uint16
    	ImageResolutionX    uint16
    	ImageResolutionY    uint16
    	Views               uint8
    	Reserved            byte
    }
    ```


-   **Get View Data**
    ```go
    ViewRecords(fmd []byte, views uint8) ([]ViewRecord, error)
    ```
    Takes in a `fmd` byte array which is an array of bytes and `views` (Can be obtained from record header) and returns an 
    array of `ViewRecords`
    
