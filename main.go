package main

import (
    "flag"
    "log"
    "os"
    "fmt"
    "github.com/jdkato/prose/tokenize"
    "net/smtp"
    str "strings"
    dgo "github.com/bwmarrin/discordgo"
)

var (
    chid, token string
    email, pass string
    limit int
    dryRun, silent bool
)

var ytpfxs [2]string = [2]string {
    "youtu.be/",
    "youtube.com/watch?v=",
}

func ytid(s string) *string {
    i := -1
    for _, pfx := range ytpfxs {
        i = str.LastIndex(s, pfx)
        if i > 0 {
            i += len(pfx)
            break
        }
    }
    if i < 0 {
        return nil
    }

    x := s[i:]
    if len(x) >= 11 {
        x = x[:11]
    }
    return &x
}

func fck(err error, msg string) {
    if err != nil {
        if msg != "" {
            log.Fatalf("Fatal: %s: %s\n", msg, err)
        } else {
            log.Fatalln(err)
        }
    }
}

func main() {
    flag.StringVar(&chid, "chid", "", "Channel ID")
    flag.StringVar(&token, "t", "", "Authentication token")
    flag.StringVar(&email, "e", "", "Authentication email (requires password)")
    flag.StringVar(&pass, "p", "", "Authentication password")
    flag.IntVar(&limit, "n", 15, "Number of messages to retrieve")
    flag.BoolVar(&dryRun, "y", false, "Enable dry-run")
    flag.BoolVar(&silent, "s", false, "Run silently")
    flag.Parse()
    if token == "" && (email == "" || pass == "") {
        log.Fatalln("Please provide a login token (-t) or both email and password (-e, -p).")
    }
    if chid == "" {
        log.Fatalln("Please provide a target channel ID (-chid).")
    }
    var credentials []interface{}
    if token != "" {
        credentials = []interface{}{token}
    } else {
        credentials = []interface{}{email, pass}
    }
    s, err := dgo.New(credentials...)
    fck(err, "")
    u, err := s.User("@me")
    fck(err, "Retrieve user ID")
    messages, err := s.ChannelMessages(chid, limit, "", "", "")
    fck(err, "Retrieve messages")
    deleteq := make([]string, 0, len(messages))
    idq := make([]string, 0, len(messages))
    for _, m := range messages {
        if m.Author.ID != u.ID {
            continue
        }
        lex := tokenize.TextToWords(m.Content)
        for _, t := range lex {
            id := ytid(t)
            if id != nil {
                idq = append([]string{*id}, idq...)
                deleteq = append([]string{m.ID}, deleteq...)
            }
        }
    }
    payload := fmt.Sprintf("https://youtube.com/watch_videos?video_ids=%s", str.Join(idq, ","))
    if !silent {
        fmt.Println(payload)
    }
    if !dryRun {
        for _, msgid := range deleteq {
            err = s.ChannelMessageDelete(chid, msgid)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Warn: Delete message ID %s: %s\n", msgid, err)
            }
        }
        _, err = s.ChannelMessageSend(chid, payload)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Fatal: Post aggregate list: %s\n", err)
            fmt.Println(payload)
            os.Exit(1)
        }
    }
    <-sig
}
