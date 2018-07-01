package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/liamg/raft/config"
	"gitlab.com/liamg/raft/gui"
	"gitlab.com/liamg/raft/pty"
	"gitlab.com/liamg/raft/terminal"
	"go.uber.org/zap"
)

func loadConfig() config.Config {

	home := os.Getenv("HOME")
	if home == "" {
		return config.DefaultConfig
	}

	places := []string{
		fmt.Sprintf("%s/.raft.yml", home),
		fmt.Sprintf("%s/.raft/config.yml", home),
		fmt.Sprintf("%s/.config/raft/config.yml", home),
	}

	for _, place := range places {
		if b, err := ioutil.ReadFile(place); err == nil {
			if c, err := config.Parse(b); err == nil {
				return *c
			} else {
				fmt.Printf("Invalid config at %s: %s\n", place, err)
			}
		}
	}

	return config.DefaultConfig
}

func getLogger(conf config.Config) (*zap.SugaredLogger, error) {

	var logger *zap.Logger
	var err error
	if conf.DebugMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to create logger: %s", err)
	}
	return logger.Sugar(), nil
}

func main() {

	// parse this
	conf := loadConfig()

	logger, err := getLogger(conf)
	if err != nil {
		fmt.Printf("Failed to create logger: %s\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Infof("Allocating pty...")
	pty, err := pty.NewPtyWithShell()
	if err != nil {
		logger.Fatalf("Failed to allocate pty: %s", err)
	}

	logger.Infof("Creating terminal...")
	terminal := terminal.New(pty, logger, conf.ColourScheme)

	g := gui.New(conf, terminal, logger)
	if err := g.Render(); err != nil {
		logger.Fatalf("Render error: %s", err)
	}

}
