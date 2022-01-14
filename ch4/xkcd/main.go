package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

const (
	MaxComicsNum = 5
	indexFile    = "index.db"
	dataFile     = "data.db"
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Transcript       string
	IMG              string
	Title            string
	ALT              string
}

func getComic(num int) (*Comic, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get %s fail: %s", url, err)
	}
	defer resp.Body.Close()
	comic := new(Comic)
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(comic); err != nil {
		return nil, fmt.Errorf("fail to parse xkcd resp: %s", err)
	}
	return comic, nil
}

func getComics() ([]*Comic, error) {
	comics := make([]*Comic, MaxComicsNum)
	for i := 1; i <= MaxComicsNum; i += 1 {
		comic, err := getComic(i)
		if err != nil {
			return nil, err
		}
		comics[i-1] = comic
	}
	return comics, nil
}

type wordsOnLine map[string]map[int]bool

func toDB(comics []*Comic) error {
	data, err := os.Create(dataFile)
	if err != nil {
		return fmt.Errorf("create db data fail: %s", err)
	}
	defer func() {
		if err := data.Close(); err != nil {
			panic(err)
		}
	}()

	index, err := os.Create(indexFile)
	if err != nil {
		return fmt.Errorf("create db index fail: %s", err)
	}
	defer func() {
		if err := index.Close(); err != nil {
			panic(err)
		}
	}()

	wol := make(wordsOnLine)
	linePosition := map[int]struct{ Start, End int }{}
	next := 0
	for line, comic := range comics {
		b, err := json.Marshal(comic)
		if err != nil {
			return fmt.Errorf("marshal comic fail: %s", err)
		}

		b = append(b, '\n')
		dataWrited, err := data.Write(b)
		if err != nil {
			return fmt.Errorf("write to db fail: %s", err)
		}

		for _, word := range strings.Fields(comic.Transcript) {
			word = strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})
			if wol[word] == nil {
				wol[word] = make(map[int]bool)
			}
			wol[word][line] = true
		}
		linePosition[line] = struct{ Start, End int }{
			next,
			next + dataWrited,
		}
		next += dataWrited
	}

	b, err := json.Marshal(wol)
	b = append(b, '\n')
	if err != nil {
		return fmt.Errorf("marshal index fail: %s", err)
	}
	if _, err = index.Write(b); err != nil {
		return fmt.Errorf("write to index fail: %s", err)
	}

	b, err = json.Marshal(linePosition)
	if err != nil {
		return fmt.Errorf("marshal index fail: %s", err)
	}
	b = append(b, '\n')
	if _, err = index.Write(b); err != nil {
		return fmt.Errorf("write to index fail: %s", err)
	}

	return nil
}

func build() error {
	comics, err := getComics()
	if err != nil {
		return err
	}
	if err = toDB(comics); err != nil {
		return err
	}
	return nil
}

func search(keywords []string) ([]*Comic, error) {
	comics := []*Comic{}

	index, err := os.Open(indexFile)
	if err != nil {
		return nil, fmt.Errorf("open index fail: %s", err)
	}
	defer index.Close()

	data, err := os.Open(dataFile)
	if err != nil {
		return nil, fmt.Errorf("open data fail: %s", err)
	}
	defer data.Close()

	reader := bufio.NewReader(index)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("read wol of index fail: %s", err)
	}
	wol := make(wordsOnLine)
	if err = json.Unmarshal(b, &wol); err != nil {
		return nil, fmt.Errorf("parse wol of index fail: %s", err)
	}

	lp := map[int]struct{ Start, End int }{}
	b, err = reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("read lp of index fail: %s", err)
	}
	if err = json.Unmarshal(b, &lp); err != nil {
		return nil, fmt.Errorf("parse lp os index fail: %s", err)
	}

	seen := map[int]bool{}
	for _, keyword := range keywords {
		if wol[keyword] == nil {
			continue
		}

		for line := range wol[keyword] {
			if seen[line] {
				continue
			}
			seen[line] = true
			pos := lp[line]
			rawComic := make([]byte, pos.End-pos.Start)
			comic := new(Comic)
			if _, err = data.ReadAt(rawComic, int64(pos.Start)); err != nil {
				return nil, fmt.Errorf("read data fail: %s", err)
			}
			if err = json.Unmarshal(rawComic, comic); err != nil {
				return nil, fmt.Errorf("parse data fail: %s", err)
			}
			comics = append(comics, comic)
		}
	}
	return comics, nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	if len(os.Args) < 2 {
		log.Fatal("invalid arguments num")
	}

	switch os.Args[1] {
	case "build":
		if err := build(); err != nil {
			log.Fatalf("build db fail: %s", err)
		}
		log.Print("build db success")
	case "search":
		comics, err := search(os.Args[2:])
		if err != nil {
			log.Fatalf("search comic fail: %s", err)
		}
		for _, comic := range comics {
			fmt.Printf("title: %s\nurl: https://xkcd.com/%d\n%s\n\n",
				comic.Title, comic.Num, comic.Transcript)
		}
	default:
		log.Fatalf("invalid arguments")
	}
}
