package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "jolt [command]",
	Short: "A task execution tool",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func experiment() {

	//ino_hash := make(map[uint64]string)

	err := filepath.WalkDir("/git", func(path string, d os.DirEntry, err error) error {
		info, error := d.Info()
		if error != nil {
			return nil
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		hash := sha1.New()

		f, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer f.Close()

		sys := info.Sys()
		if sys == nil {
			log.Fatal("error: sys")
			return nil
		}

		_, ok := sys.(*syscall.Stat_t)
		if !ok {
			log.Fatal("error: stat")
			return nil
		}

		var digest string
		ok = false // digest, ok := ino_hash[stat.Ino]
		if !ok {
			io.Copy(hash, f)
			digest = hex.EncodeToString(hash.Sum(nil))
			// ino_hash[stat.Ino] = digest
		}

		fmt.Println(digest, path)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
