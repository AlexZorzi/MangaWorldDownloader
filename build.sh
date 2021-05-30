#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o ./build/MangaDownloader_linux_amd64 
echo MangaDownloader_linux_amd64 
env GOOS=linux GOARCH=386 go build -o ./build/MangaDownloader_linux_x86 
echo MangaDownloader_linux_x86 
env GOOS=linux GOARCH=arm64 go build -o ./build/MangaDownloader_linux_arm64
echo MangaDownloader_linux_arm64
env GOOS=android GOARCH=arm64 go build -o ./build/MangaDownloader_android_arm64
echo MangaDownloader_android_arm64
env GOOS=windows GOARCH=amd64 go build -o ./build/MangaDownloader_windows_amd64
echo MangaDownloader_windows_amd64
env GOOS=windows GOARCH=386 go build -o ./build/MangaDownloader_windows_386
echo MangaDownloader_windows_386
 
