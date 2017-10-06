package main

import (
	"os"
	"path/filepath"
	"runtime"
)

const column = 20

type config struct {
	MemoDir   string `toml:"memodir"`
	Editor    string `toml:"editor"`
	Column    int    `toml:column`
	SelectCmd string `toml:selectcmd`
	GrepCmd   string `toml:GrepCmd`
}

func loadConfig(cfg *config) error {
	dir := os.Getenv("HOME")
	if dir == "" && runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "memo")
		}
		dir = filepath.Join(dir, "memo")
	} else {
		dir = filepath.Join(dir, ".config", "memo")
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	file := filepath.Join(dir, "config.toml")

	_, err := os.Stat(file)
	if err == nil {
		_, err := toml.DecodeFile(file, cfg)
		return err
	}
	if !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	dir = filepath.Join(dir, "_posts")
	os.MkdirAll(dir, 0700)
	cfg.MemoDir = filepath.ToSlash(dir)
	cfg.Editor = "vim"
	cfg.Column = 20
	cfg.SelectCmd = "peco"
	cfg.GrepCmd = "grep"
	return toml.NewEncoder(f).Encode(cfg)
}
