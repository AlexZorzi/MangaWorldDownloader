# MangaWorld Downloader 
Ein Downloader für mangaworld.ac,
dies ist ein Projekt für Bildungszwecke, das nicht mit mangaworld.ac verbunden ist. Ich bin nicht für die Nutzung und die Bilder auf mangaworld.ac verantwortlich.

[IT Übersetzung](README_IT.md)
[EN Übersetzung](README.md)

# Features
- Herunterladen von ganzen Mangas oder einzelnen Kapiteln.
- Multi-Thread Download für das Herunterladen von Seiten mit maximaler Geschwindigkeit.
- Preview-Download

# Im Falle eines CloudFlare-Captchas
Um die Sperre zu umgehen, benötigt das Programm das cf_clearance-Cookie und den User-Agent des Browsers, den Sie zum Lösen des Captchas verwendet haben.

- Token: Lösen Sie das Captcha auf www.mangaworld.ac und ermitteln Sie in der Dev Console den Wert des cf_clearance-Cookies
  ![Token](https://i.imgur.com/HYUu0M0.png)

- BenutzerAgent: Unter www.whatsmyua.acfo können Sie Ihren UserAgent sehen
  ![Useragent](https://i.imgur.com/nZZfCt1.png)

das Programm fragt Sie automatisch nach diesen beiden Parametern, wenn es sie braucht, und speichert sie dann automatisch in einer JSON-Datei, damit Sie sie
auch nach dem Beenden des Programms verwenden können.

Manuell: Sie können die Werte manuell in eine Datei namens `cred.json` mit dieser Struktur im gleichen Ordner wie die ausführbare Datei eingeben.


```
{
"useragent" : "useragent value here"
"cf_clearance" : "cf_clearance value here"
}
```
