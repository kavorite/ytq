# ytq
Command-line application for compressing posts of YouTube links in a given Discord channel (the name stands for YouTube Queue).

Compiling:
1. [Install Go](https://golang.org/dl/) and add the `go` tool to your PATH.
2. Install dependencies:
  - `go get github.com/bwmarrin/discordgo`
  - `go get github.com/jdkato/prose/tokenize`
3. `go build`
4. `ytq -chid <chan> -t <authentication token>`
