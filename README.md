[![Go Report Card](https://goreportcard.com/badge/github.com/yombu360/ANSI-378-Decoder)](https://goreportcard.com/report/github.com/yombu360/ANSI-378-Decoder)

# ANSI-378 Decoder

Breaks down a ANSI INCITS 378-2004 Finger Minutiae Format

## Method Information

Supports go versions > 1.12.1. To install add this line to your `go.mod` file 

`github.com/yombu360/ANSI-378-Decoder`

- **Get Record headers**
    ```
    GetFMDRecordHeaders(fmdByteArray []byte) (recordHeaders map[string]interface{}, err error)
    ```
    Takes in a `fmdByteArray` which is an array of bytes and returns the `recordHeader`
    
- **Get View Data**
    ```
    GetViewRecords(fmdByteArray []byte, views uint8) (fingerViewRecords map[string][]interface{}, err error)
    ```
    Takes in a `fmdByteArray` which is an array of bytes and `views` (Can be obtained from record header) and returns a 
    map of `fingerViewRecords`
    
    Sample Output:
    ```
    map[
        Finger View 1:
            [
                map[
                    Finger Position:0 
                    Finger Quality:86 
                    Impression Type:0 
                    Minutiae Count:45 
                    View Number:0
                ] 
                map[
                    Angle:120 
                    Minutiae Type:1 
                    Quality:100 
                    Reserved:0 
                    X coordinate:80 
                    Y coordinate:291
                ]
                :
                :
            ]
        :
        :
        :
        Finger View N:
            :
            :
        
    ```

