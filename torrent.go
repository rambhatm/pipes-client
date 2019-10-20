package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
)

var (
	builtinAnnounceList = [][]string{
		{"udp://tracker.openbittorrent.com:80"},
		{"udp://tracker.publicbt.com:80"},
		{"udp://tracker.istole.it:6969"},
	}
)

//Creates a magnet link for the file
func CreateMagnet(path string) (magnet string) {

	mi := metainfo.MetaInfo{
		AnnounceList: make([][]string, 0),
	}
	mi.SetDefaults()
	mi.Comment = ""
	mi.CreatedBy = "pipes"

	info := metainfo.Info{
		PieceLength: 256 * 1024,
	}
	err := info.BuildFromFilePath(path)
	if err != nil {
		log.Fatal(err)
	}

	mi.InfoBytes, err = bencode.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	info, err = mi.UnmarshalInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error unmarshalling info: %s", err)
		os.Exit(1)
	}

	magnet = mi.Magnet(info.Name, mi.HashInfoBytes()).String()

	return
}
