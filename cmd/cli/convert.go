package main

import (
	"encoding/json"
	"os"

	"github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/slf/v2"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Key struct {
	Row    uint8  `json:"row"`
	Col    uint8  `json:"col"`
	Finger string `json:"finger"`
}

type Layout struct {
	Name  string         `json:"name"`
	Owner uint64         `json:"user"`
	Board string         `json:"board"`
	Keys  map[string]Key `json:"keys"`
	Free  []Key          `json:"free"`
}

type Owners map[uint64][]string
type Links map[string]string

func main() {
	var (
		owners Owners = make(Owners)
		links  Links  = make(Links)
		base          = "tmp/cmini/"
		prefix        = "layouts/"
	)

	repo, err := git.PlainOpen(base)
	if err != nil {
		log.Error(err)
		return
	}

	ref, err := repo.Head()
	if err != nil {
		log.Error(err)
		return
	}

	owners_data, err := os.ReadFile("owners.json")
	if err != nil {
		log.Error(err)
		return
	}
	if err := json.Unmarshal(owners_data, &owners); err != nil {
		log.Error(err)
		return
	}

	links_data, err := os.ReadFile(base + "links.json")
	if err != nil {
		log.Error(err)
		return
	}
	if err := json.Unmarshal(links_data, &links); err != nil {
		log.Error(err)
		return
	}

	entries, err := os.ReadDir(base + prefix)
	if err != nil {
		log.Error(err)
		return
	}

	storage.Init(prefix)

outer:
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := prefix + entry.Name()

		file, err := os.ReadFile(base + fileName)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fileName)
		fileIter, err := repo.Log(&git.LogOptions{
			From:     ref.Hash(),
			Order:    git.LogOrderCommitterTime,
			FileName: &fileName,
		})
		if err != nil {
			log.Error(err)
			continue
		}

		var firstCommit, lastCommit *object.Commit
		err = fileIter.ForEach(func(c *object.Commit) error {
			if lastCommit == nil {
				lastCommit = c
			}
			firstCommit = c
			return nil
		})
		if err != nil {
			log.Error(err)
			continue
		}

		firstCommitTime := firstCommit.Committer.When
		lastCommitTime := lastCommit.Committer.When

		var layout Layout
		if err = json.Unmarshal(file, &layout); err != nil {
			log.Error(err)
			continue
		}

		var ordered map[uint8]map[uint8]slf.Key = make(map[uint8]map[uint8]slf.Key, 0)
		for k, v := range layout.Keys {
			// Handle "TB" finger => "LT"
			if v.Finger == "TB" {
				v.Finger = "LT"
			}

			var slfKey slf.Key
			keyJson, err := json.Marshal(v)
			if err != nil {
				log.Error(err)
				continue outer
			}
			json.Unmarshal(keyJson, &slfKey)
			slfKey.Char = k
			if ordered[v.Row] == nil {
				ordered[v.Row] = make(map[uint8]slf.Key, 0)
			}

			ordered[v.Row][v.Col] = slfKey
		}

		for _, v := range layout.Free {
			var slfKey slf.Key
			keyJson, err := json.Marshal(v)
			if err != nil {
				log.Error(err)
				continue outer
			}

			json.Unmarshal(keyJson, &slfKey)
			if ordered[v.Row] == nil {
				ordered[v.Row] = make(map[uint8]slf.Key, 0)
			}

			ordered[v.Row][v.Col] = slfKey
		}

		var keys []slf.Key
		for row := range uint8(len(ordered)) {
			for col := range uint8(len(ordered[row])) {
				keys = append(keys, ordered[row][col])
			}
		}

		var author string
		// Handle no author for a given owner, or owner == 0
		if len(owners[layout.Owner]) == 0 || layout.Owner == 0 {
			author = "akl"
		} else {
			author = owners[layout.Owner][0]
		}

		converted := slf.Layout{
			Name:     layout.Name,
			Author:   author,
			Owner:    layout.Owner,
			Created:  firstCommitTime,
			Modified: lastCommitTime,
			Link:     links[layout.Name],
			Boards:   []string{layout.Board},
			Keys:     keys,
		}

		if err := storage.Put(converted); err != nil {
			log.Error(err)
		}
	}
}
