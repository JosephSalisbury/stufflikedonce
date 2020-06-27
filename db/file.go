package db

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

type File struct {
	path string
}

func NewFileDatabase(path string) (*File, error) {
	if path == "" {
		return nil, errors.New("file database path cannot be empty")
	}

	f := &File{
		path: path,
	}

	_, err := os.Stat(f.path)
	if os.IsNotExist(err) {
		_, err := os.Create(f.path)
		log.Printf("creating empty file database at %v", path)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func (f *File) Save(tweet twitter.Tweet) error {
	url := fmt.Sprintf("https://twitter.com/%v/status/%v", tweet.User.ScreenName, tweet.IDStr)

	b, err := ioutil.ReadFile(f.path)
	if err != nil {
		return err
	}
	s := string(b)
	if strings.Contains(s, url) {
		log.Printf("%v already in database", url)
		return nil
	}

	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%v\n", url)); err != nil {
		return err
	}

	log.Printf("saved %v", url)

	return nil
}

func (f *File) GetRandomLikedTweet() (string, error) {
	b, err := ioutil.ReadFile(f.path)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(b), "\n")

	r := lines[rand.Intn(len(lines))]

	return r, nil
}
