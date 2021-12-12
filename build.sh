#!/bin/bash


#linux
env GOOS=linux GOARCH=amd64 go build -o ./build/MangaDownloader_linux_amd64
echo MangaDownloader_linux_amd64
env GOOS=linux GOARCH=386 go build -o ./build/MangaDownloader_linux_x86
echo MangaDownloader_linux_x86
env GOOS=linux GOARCH=arm64 go build -o ./build/MangaDownloader_linux_arm64
echo MangaDownloader_linux_arm64

#darwin (macos)
env GOOS=darwin GOARCH=amd64 go build -o ./build/MangaDownloader_darwin_amd64
echo MangaDownloader_darwin_amd64
env GOOS=darwin GOARCH=arm64 go build -o ./build/MangaDownloader_darwin_arm64
echo MangaDownloader_darwin_arm64

#windows
env GOOS=windows GOARCH=amd64 go build -o ./build/MangaDownloader_windows_amd64
echo MangaDownloader_windows_amd64
env GOOS=windows GOARCH=386 go build -o ./build/MangaDownloader_windows_386
echo MangaDownloader_windows_386

#freebsd
env GOOS=freebsd GOARCH=386 go build -o ./build/MangaDownloader_freebsd_386
echo MangaDownloader_freebsd_386
env GOOS=freebsd GOARCH=amd64 go build -o ./build/MangaDownloader_freebsd_amd64
echo MangaDownloader_freebsd_amd64