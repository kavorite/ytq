# ytq
Command-line application for compressing posts of YouTube links in a given Discord channel (YouTube Queue).

Compiling:
1. If you haven't already, [install git](https://git-scm.com/download/), and add it to your `PATH`
2. `git clone https://github.com/kavorite/ytq` (this repository)
3. [Install Go](https://golang.org/dl/) and add the `go` tool to your `PATH`
4. Install dependencies:
  - `go get github.com/bwmarrin/discordgo`
  - `go get github.com/jdkato/prose/tokenize`
5. `go build`
6. `ytq -chid <chan> -t <authentication token>`
