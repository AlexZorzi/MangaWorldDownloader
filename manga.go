package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/buger/jsonparser"
	"net/http"
	"strconv"
)

type Manga struct {
	ID          string   `json:"_id"`
	IDLink      int      `json:"linkId"`
	Title       string   `json:"title"`
	ExtraTitles []string `json:"extraTitles"`
	Author      []string `json:"author"`
	Artist      []string `json:"artist"`
	Genres      []struct {
		Slug string `json:"slug"`
	}
	References struct {
		MangaupdatesId int `json:"mangaupdatesId"`
		MangadexId     int `json:"mangadexId"`
		AnilistId      int `json:"anilistId"`
		MalId          int `json:"malId"`
	}
	Fansub struct {
		Name string `json:"name"`
		Link string `json:"Link"`
	}
	Vm18          bool   `json:"vm18"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Trama         string `json:"trama"`
	Year          int    `json:"year"`
	ChaptersCount string `json:"chaptersCount"`
	VolumesCount  string `json:"volumesCount"`
	Image         string `json:"image"`
	Slug          string `json:"slug"`
	SlugFolder    string `json:"slugFolder"`
	ReadMode      string `json:"readMode"`
	Pages         Pages
}

//o.w[1][2].globalData.latestMangas
func bufToJson(response *http.Response) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	BodyStr := buf.String()
	jsondata := Between(BodyStr, ".concat(", ")</script>")
	data := []byte(jsondata)
	return data
}

func GetLatestMangaIndex() int64 {
	url := "https://www.mangaworld.io/"
	var latest int64
	for {
		resp, err := GetWebPage(url)
		data := bufToJson(resp)
		latestMangaTemp, err := jsonparser.GetInt(data, "o", "w", "[1]", "[2]", "globalData", "latestMangas", "[0]", "linkId")
		if err != nil {
			//println("CloudFlare Error refreshing proxy")
			//ChangeProxy()
			println("CloudFlare Error refreshing cookie")
			ChangeCookie()
			ChangeUserAgent()
			SaveCookies()
		} else {
			latest = latestMangaTemp
			break
		}
	}
	return latest
}

func GetManga(ID int) (Manga, error) {
	url := "https://www.mangaworld.io/manga/" + strconv.Itoa(ID)
	var mangadata []byte
	var readdata []byte
	for {
		resp, err := GetWebPage(url)
		if resp.StatusCode == 404 {
			return Manga{}, errors.New("Error 404 skip of ID")
		}
		data := bufToJson(resp)
		mangadataTemp, _, _, err := jsonparser.Get(data, "o", "w", "[0]", "[2]", "manga")
		readdataTemp, _, _, err2 := jsonparser.Get(data, "o", "w", "[3]", "[2]", "pages")
		if err2 != nil {
			readdataTemp, _, _, err2 = jsonparser.Get(data, "o", "w", "[4]", "[2]", "pages")
		}
		if err != nil && err2 != nil {
			//println("CloudFlare Error refreshing proxy")
			//ChangeProxy()
			println(url)
			println("CloudFlare Error refreshing cookie")
			ChangeCookie()
			ChangeUserAgent()
			SaveCookies()
		} else {
			mangadata = mangadataTemp
			readdata = readdataTemp
			break
		}
	}
	var manga Manga
	json.Unmarshal(mangadata, &manga)

	var pages Pages
	json.Unmarshal(readdata, &pages)

	manga.Pages = pages
	return manga, nil
}
