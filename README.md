# ytq
Command-line application for compressing posts of YouTube links in a given Discord channel (YouTube Queue).

Compiling:
1. If you haven't already, [install git](https://git-scm.com/download/), and add it to your `PATH`
2. `git clone https://github.com/kavorite/ytq` (this repository)
3. [Install Go](https://golang.org/dl/) and add the `go` tool to your `PATH`
4. Install dependencies:
    - `go get github.com/bwmarrin/discordgo`
    - `go get github.com/jdkato/prose/tokenize`
5. `cd` to the root of your local clone and `go build`
6. `ytq -chid <chan> -t <authentication token>`

I assume no responsibility for consequences of failure to comply with the Discord ToS. This software has no warranty, it is provided “as is.” It is your responsibility to validate the behavior of the routines and their accuracy using the source code provided, or to purchase support and warranties from commercial redistributors. 
