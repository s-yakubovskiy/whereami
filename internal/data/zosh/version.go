package zosh

import (
	"fmt"
	"time"
)

var (
	Version     string
	Commit      string
	showVersion bool
)

type ZoshVersion struct {
	Version    string `json:"version"`
	CommitHash string `json:"commit_hash"`
	Date       string `json:"date"`
}

func NewZoshVersion() *ZoshVersion {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if Version == "" {
		Version = "dev-version"
	}
	if Commit == "" {
		Commit = "dev-commit"
	}

	return &ZoshVersion{
		Version:    Version,
		CommitHash: Commit,
		Date:       currentTime,
	}
}

func (zv *ZoshVersion) GetVersion() string {
	return zv.Version
}

func (zv *ZoshVersion) GetCommitHash() string {
	return zv.CommitHash
}

func (zv *ZoshVersion) GetDate() string {
	return zv.Date
}

func (zv *ZoshVersion) PrintVersion() {
	fmt.Println("\nBuild Info:")
	fmt.Println("  Version:", Version)
	fmt.Println("  Commit:", Commit)
}
