package main

import (
	"log"

	"github.com/anacrolix/torrent"
)

const (
	shareDir = "./share"
)

func main() {
	btClientConfig := torrent.NewDefaultClientConfig()
	btClientConfig.DataDir = shareDir

	btClient, err := torrent.NewClient(btClientConfig)
	if err != nil {
		log.Fatal("BT client create error", err)
	}

	defer btClient.Close()
}
