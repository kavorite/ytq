# ytq
Command-line application for compressing posts of YouTube links in a given Discord channel (YouTube Queue).

Compiling:
1. If you haven't already, [install git](https://git-scm.com/download/), and add it to your `PATH`
1. [Install Go](https://golang.org/dl/) and add the `go` tool to your `PATH`
2. Install dependencies:
  - `go get github.com/bwmarrin/discordgo`
  - `go get github.com/jdkato/prose/tokenize`
3. `go build`
4. `ytq -chid <chan> -t <authentication token>`
