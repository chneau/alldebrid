# alldebrid

Simple AllDebrid API client

## Example of use

```go
package main

import (
    "log"

    "github.com/chneau/alldebrid"
)

func main() {
    client := alldebrid.New()
    _ = client
    logs, err := client.Connect("username", "password")
    if err != nil {
        panic(err)
    }
    log.Println(logs)
    link, err := client.GetDownloadLink("https://mega-downloader.example/superweirdlink.mp4")
    if err != nil {
        panic(err)
    }
    log.Println(link)
}

```
