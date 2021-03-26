# MangaWorld Downloader
Un Downloader Per mangaworld.cc,
progetto a scopo didattico non affiliato a mangaworld.cc, non sono responsabile dell'utilizzo e delle immagini hostate su mangaworld.cc

# Features
- Download di manga interi o singoli capitoli
- MultiThreaded Download per scaricare le pagine alla massima velocità
- Download della Preview

# In caso di un CloudFlare captcha
Per passare il blocco il programma avrà bisogno del cookie cf_clearance e dell'useragent del browser che avete utilzzato per risolvere il captcha

- Token: Risolvete il captcha su www.mangaworld.cc e in Dev Console prendete il valore del cookie cf_clearance 
  ![Token](https://i.imgur.com/HYUu0M0.png)
  
- UserAgent: Al sito www.whatsmyua.info è possibile vedere il proprio useragent
  ![Useragent](https://i.imgur.com/nZZfCt1.png)
  
il programma vi chiederà in automatico questi due parametri se ne avrà bisogno, poi li salverà in automatico in un file JSON così da poterli utilizzare 
anche dopo che il programma è stato chiuso.

Manualmente: è possibile inserire manualmente i valori in un file chiamato `cred.json` con questa struttura nella stessa cartella dell'eseguibile

```
{
"useragent" : "useragent value here"
"cf_clearance" : "cf_clearance value here"
}
```