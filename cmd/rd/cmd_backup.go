package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/brotli/go/cbrotli"
	"github.com/schollz/progressbar/v3"

	"github.com/rythoris/rd"
)

type BackupCommand struct {
	List     *ListBackupCommand     `arg:"subcommand:list"     help:"list generated backups"`
	Download *DownloadBackupCommand `arg:"subcommand:download" help:"download backup"`
	Create   *CreateBackupCommand   `arg:"subcommand:create"   help:"create a new backup"`
}

type ListBackupCommand struct{}

func (c ListBackupCommand) Run(token string) int {
	backups, err := rd.GetBackups(token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}

	for _, b := range backups {
		fmt.Printf("%s\t%s\n", b.ID, b.Created.Format(time.ANSIC))
	}
	return 0
}

type DownloadBackupCommand struct {
	ID          string `arg:"positional,required" help:"backup id"`
	Destination string `arg:"positional,required" help:"download file path" placeholder:"DST"`
	Format      string `arg:"-f,--format" default:"csv" help:"file format (possible values: csv, html)"`
}

func (c DownloadBackupCommand) Run(token string) int {
	if c.Format != "csv" && c.Format != "html" {
		fmt.Fprintf(os.Stderr, "[!] ERROR: unknown backup format: %s\n", c.Format)
		fmt.Fprintf(os.Stderr, "[-] Possible values are: html, csv\n")
		return 1
	}

	backups, err := rd.GetBackups(token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: could not get backups: %s\n", err.Error())
		return 1
	}

	found := false
	for _, b := range backups {
		if b.ID == c.ID {
			found = true
		}
	}

	if !found {
		fmt.Fprintf(os.Stderr, "[!] ERROR: invalid backup id: %s\n", c.ID)
		return 1
	}

	f, err := os.OpenFile(c.Destination, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: could not create file: %s: %s\n", c.Destination, err.Error())
		return 1
	}
	defer f.Close()

	client := http.Client{Timeout: rd.HTTPClientTimeout}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/backup/%s.%s", rd.ApiBaseURL, c.ID, c.Format), http.NoBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: http request error: %s\n", err.Error())
		return 1
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept-Encoding", "br")

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: http client error: %s\n", err.Error())
		return 1
	}

	bar := progressbar.NewOptions(
		int(res.ContentLength),
		progressbar.OptionClearOnFinish(),
	)

	if h := res.Header.Get("Content-Encoding"); h != "br" {
		fmt.Fprintf(os.Stderr, "[!] ERROR: Currently other content encodings are not supported: expected 'br' got '%s'\n", h)
		return 1
	}

	brotliReader := cbrotli.NewReader(res.Body)
	defer brotliReader.Close()

	_, err = io.Copy(io.MultiWriter(f, bar), brotliReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: io error: %s\n", err.Error())
		return 1
	}

	fmt.Fprintf(os.Stdout, "[+] Done: %s\n", c.Destination)

	return 0
}

type CreateBackupCommand struct{}

func (c CreateBackupCommand) Run(token string) int {
	if err := rd.CreateBackup(token); err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}
	fmt.Fprintf(os.Stderr, "[+] NOTE: Creating new backups may take a while... raindrop.io will send you an email when the backup is ready.\n")
	return 0
}
