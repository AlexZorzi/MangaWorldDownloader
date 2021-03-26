package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var WG sync.WaitGroup
var method int = 1
var DEBUG bool = true
var COOKIE string = ""
var USERAGENT string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36"
var PROXY []string
var LATESTMANGAIdLink int64

// COOKIE needs only cf_clearance
/*
	method:
			0 database Full sync
			1 single sync
			2 telegram event listener

*/
func main() {
	//SyncProxyList()
	//ChangeProxy()
	getCred()
	LATESTMANGAIdLink = GetLatestMangaIndex()
	fmt.Println("##################################################")
	fmt.Println("#            MangaWorld Downloader               #")
	fmt.Println("#                                                #")
	fmt.Println("#            Version: 1.0                        #")
	fmt.Println("#            Latest Update: 26/03/2021           #")
	fmt.Println("#                                                #")
	fmt.Println("##################################################")
	single()
	WG.Wait()
}

func getCred(){

	code, err := GetWebPage("https://www.mangaworld.cc/")
	if(err == nil){
		if(code.StatusCode == 200){
			return
		}
	}
	if _, err := os.Stat("./cred.json"); !os.IsNotExist(err) {
		data, _ := ioutil.ReadFile("./cred.json")
		var cred Cred
		json.Unmarshal(data, &cred)
		COOKIE = cred.Clearance
		USERAGENT = cred.Useragent
	}else {
		fmt.Println("The file cred.json dosen't exist!, creating one...")
		ChangeCookie()
		ChangeUserAgent()
		SaveCookies()
	}
}

func single() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Manga Url/ID to Check: ")
	Url, _ := reader.ReadString('\n')
	UrlClean := strings.ReplaceAll(Url, "\n", "")
	// sample url https://www.mangaworld.cc/manga/1807/murim-login/
	var id int64
	for _, idFind := range strings.Split(UrlClean, "/") {
		idInt, err := strconv.ParseInt(idFind, 10, 64)
		if err == nil {
			id = idInt
			break
		}
	}
	if id == 0 && id >= 170 && id <= LATESTMANGAIdLink { // check if an id is found
		fmt.Println("Error: Url Not Valid, Quitting...")
		os.Exit(0)
	}
	fmt.Println("Info: Anime id = " + strconv.Itoa(int(id)))
	mangasaved, err := GetManga(int(id))
	if err != nil {
		println(err.Error())
	}
	fmt.Println("Selected: "+mangasaved.Title)
	fmt.Print("0 for single 1 for whole: ")
	whole, _ := reader.ReadString('\n')
	wholecl := strings.ReplaceAll(whole, "\n", "")
	if wholecl == "0" {
		if len(mangasaved.Pages.Volumes) > 0 {
			fmt.Println("Pick: What Volume You wanna add/update? ")
			var aviableVolumes []struct {
				Volume   Volume    `json:"volume"`
				Chapters []Chapter `json:"chapters"`
			}
			for index, volume := range mangasaved.Pages.Volumes {
				println(index, ")    PickOption: Volume "+volume.Volume.Name)
				aviableVolumes = append(aviableVolumes, volume)
			}
			pick, _ := reader.ReadString('\n')
			pickClean, _ := strconv.Atoi(strings.ReplaceAll(pick, "\n", ""))
			selectedVolume := aviableVolumes[pickClean]
			fmt.Println("Pick: What Chapter You wanna add/update? ")
			var aviableChapters []Chapter
			for index, chapter := range selectedVolume.Chapters {
				println(index, ")  PickOption: Chapters "+chapter.Name)
				aviableChapters = append(aviableChapters, chapter)
			}
			pick, _ = reader.ReadString('\n')
			pickClean, _ = strconv.Atoi(strings.ReplaceAll(pick, "\n", ""))
			selectedChapter := aviableChapters[pickClean]
			fmt.Println("Starting Chapter Download")
			DownloadChapter(selectedChapter, selectedVolume.Volume, mangasaved)
			DownloadPreview(mangasaved)
		} else {
			fmt.Println("Pick: What Chapter You wanna add/update? ")
			var aviableChapters []Chapter
			for index, chapter := range mangasaved.Pages.SingleChapters {
				println(index, ")  PickOption: Chapters "+chapter.Name)
				aviableChapters = append(aviableChapters, chapter)
			}
			pick, _ := reader.ReadString('\n')
			pickClean, _ := strconv.Atoi(strings.ReplaceAll(pick, "\n", ""))
			selectedChapter := aviableChapters[pickClean]
			fmt.Println("Starting Chapter Download")
			DownloadChapterNoVolume(selectedChapter, mangasaved)
			DownloadPreview(mangasaved)

		}
	} else if wholecl == "1" {
		DownloadManga(mangasaved)
	} else {
		println("Error")
	}
}