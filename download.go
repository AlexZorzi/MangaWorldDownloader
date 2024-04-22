package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DownloadManga(m Manga) {
	// https://cdn.mangaworld.ac/chapters/blust-5fde5e956fe15b440d71d1c6/oneshot-5fde5ea093f7544385e33e4d/1.png
	// 									 slugManga - idManga / slugChapther - idChapter / pageFilename
	DownloadPreview(m)
	// TODO: download Banner
	// Download Volumes
	for _, volumeEntity := range m.Pages.Volumes {
		for _, chapter := range volumeEntity.Chapters {
			DownloadChapter(chapter, volumeEntity.Volume, m)
		}
	}
	// insert all SingleChapters
	for _, chapter := range m.Pages.SingleChapters {
		DownloadChapterNoVolume(chapter, m)
	}
}

func DownloadPreview(m Manga) {
	PATH := filepath.FromSlash("./" + m.Slug + "/")
	WG.Add(1)
	go downloadWPath("https://cdn.mangaworld.ac"+m.Image, PATH, m.Slug+"."+strings.Split(m.Image, ".")[1])
}
func DownloadChapterNoVolume(c Chapter, m Manga) {
	WG.Wait()
	var floatC float64
	fmt.Println(c.Name)
	if c.Name == "Oneshot" {
		floatC = 0.0
	} else {
		floattry, err2 := strconv.ParseFloat(strings.Split(c.Name, " ")[1], 32)
		if err2 != nil {
			panic(err2)
		}
		floatC = floattry
	}
	chapterNumber := fmt.Sprintf("%.1F", floatC)
	for _, page := range c.Pages {
		API := "https://cdn.mangaworld.ac/chapters/" + m.SlugFolder + "-" + m.ID + "/" + c.Slug + "-" + c.ID + "/" + page
		PATH := filepath.FromSlash("./" + m.Slug + "/chapter " + chapterNumber + "/")
		WG.Add(1)
		go downloadWPath(API, PATH, page)
	}
}
func DownloadChapter(c Chapter, v Volume, m Manga) {
	floatV, err := strconv.ParseFloat(strings.Split(v.Name, " ")[1], 32)
	if err != nil {
		panic(err)
	}
	volumeNumber := fmt.Sprintf("%.1f", floatV)
	fmt.Println("Download: Volume " + v.Name + " Chapter " + c.Name)
	floatC, err2 := strconv.ParseFloat(strings.Split(c.Name, " ")[1], 32)
	if err2 != nil {
		panic(err2)
	}
	chapterNumber := fmt.Sprintf("%.1F", floatC)
	WG.Wait()
	for _, page := range c.Pages {
		API := "https://cdn.mangaworld.ac/chapters/" + m.SlugFolder + "-" + m.ID + "/" + v.Slug + "-" + v.ID + "/" + c.Slug + "-" + c.ID + "/" + page
		PATH := filepath.FromSlash("./" + m.Slug + "/volume " + volumeNumber + "/chapter " + chapterNumber + "/")
		WG.Add(1)
		go downloadWPath(API, PATH, page)
	}
}

func downloadWPath(url string, path string, filename string) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "charset=UTF-8")
	req.Header.Set("cookie", COOKIE)
	req.Header.Set("user-agent", USERAGENT)
	client := &http.Client{}
	response, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	os.MkdirAll(path, os.ModePerm)
	file, err := os.Create(path + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	WG.Done()
}

func GetWebPage(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "charset=UTF-8")
	req.Header.Set("cookie", COOKIE)
	req.Header.Set("user-agent", USERAGENT)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println(resp.Status)
		panic(err)
	}
	if DEBUG {

		println(resp.Status)
	}
	return resp, err
}

func SyncProxyList() {
	url := "https://raw.githubusercontent.com/clarketm/proxy-list/master/proxy-list-raw.txt"
	resp, err := GetWebPage(url)
	if err != nil {
		panic(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := strings.Split(string(bodyBytes), "\n")
	bodyString = bodyString[2 : len(bodyString)-1]
	PROXY = bodyString
}
func ChangeProxy() {
	os.Setenv("HTTP_PROXY", PROXY[len(PROXY)-1])
	os.Setenv("HTTPS_PROXY", PROXY[len(PROXY)-1])
	println("Proxy changed in " + PROXY[len(PROXY)-1])
	PROXY = PROXY[0 : len(PROXY)-2]
}

func ChangeCookie() {
	reader := bufio.NewReader(os.Stdin)
	print("Cookie Error, supply new cookie: ")
	cookie, _ := reader.ReadString('\n')
	cookienew := strings.ReplaceAll(cookie, "\n", "")
	COOKIE = "cf_clearance=" + cookienew
}

func ChangeUserAgent() {
	reader := bufio.NewReader(os.Stdin)
	print("supply new USERAGENT: ")
	useragent, _ := reader.ReadString('\n')
	useragentnew := strings.ReplaceAll(useragent, "\n", "")
	USERAGENT = useragentnew
}

func SaveCookies() {
	ioutil.WriteFile("cred.json", []byte("{\n \"useragent\":\""+USERAGENT+"\", \n \"cf_clearance\":\""+COOKIE+"\" \n}"), 0755)
}
