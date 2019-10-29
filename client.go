package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	shareDir  = "./share"
	targetDir = "./target"
)

func getMagnets(user string, node string) (magnetList []string) {

	resp, err := http.Get("https://copymyfile.herokuapp.com/GetPipe?user=" + user + "&node=" + node)
	if err != nil {
		log.Fatal("http get err ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)

		res := make([]map[string]string, 0)
		json.Unmarshal(body, &res)

		for _, pipe := range res {
			magnetList = append(magnetList, pipe["data"])
		}
	}

	return
}

func setMagnets(user string, node string) {
	magnet := CreateMagnet(shareDir)
	pipeData := url.Values{
		"user": {user},
		"node": {node},
		"data": {magnet},
	}

	resp, err := http.PostForm("https://copymyfile.herokuapp.com/SetPipe", pipeData)
	if err != nil {
		log.Fatal("http post err ", err)
	}
	if resp.StatusCode == http.StatusCreated {
		log.Println("pipe created")
	} else {
		log.Printf("%#v", resp)
	}

}

func main() {

	//setMagnets("clTest", "n1")
	mags := getMagnets("clTest", "n1")
	log.Println(mags)

	/*
		btClientConfig := torrent.NewDefaultClientConfig()
		btClientConfig.DataDir = targetDir

		btClient, err := torrent.NewClient(btClientConfig)
		if err != nil {
			log.Fatal("BT client create error", err)
		}
		defer btClient.Close()
				magnetList := getMagnets()
				for _, magnet := range magnetList {
					t, _ := btClient.AddMagnet(magnet)
					<-t.GotInfo()
					t.DownloadAll()
				}

			btClient.WaitAll()
	*/

}
