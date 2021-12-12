# Pagina Archivio
Json Sono caricati in pagina tra '.concat(' & ') \<script>'
https://www.mangaworld.in/archive?page=1

Questa è la struttura importante
 ``` 
 "o":{
     "w":[
         2[
             "s0-13",
             2,
             {
               "results":2125, //totale manga
               "mangas":[
                   {
                       "extraTitles":[
                           "Onii-chan is done for",
                           "お兄ちゃんはおしまい！",
                           "Onii-chan wa Oshimai"
                       ]
                       "author":[
                        "Nekotoufu"
                       ],
                       "artist":[
                        "Nekotoufu"
                       ],
                       "genres":[
                           {
                               "name":"commedia",
                           }
                           ...
                       ]
                       "references":[
                            "mangaupdatesId":146713,
                            "mangadexId":22615,
                            "anilistId":100080,
                            "malId":112592,
                       ],
                       "fansub":{
                        "name":"Betta Scans",
                        "link":"https://www.facebook.com/BettaScans/"
                        },
                       "vm18":false,
                       "animeLink":null,
                       "_id":"5faa2373de75a27e9df97f61",
                       "title":" Onii-chan Is Done For!",
                       "status":"ONGOING",
                       "type":"DOUJINSHI",
                       "trama":"**", //la trama è qui
                       "year":2017,
                       "volumesCount":null,
                       "chaptersCount":null,
                       "image":"/mangas/5faa2373de75a27e9df97f61.png",
                       "slug":"onii-chan-is-done-for",
                       "slugFolder":"onii-chan-is-done-for",
                       "linkId":2070,  
                   },
                   ... //ripetuto per ogni manga in pagina
               ],
             },
         ],
     ]
 }
 ```

# Pagina Manga
si trova sempre nello stesso punto nel file html del precedente
https://www.mangaworld.in/manga/2070/onii-chan-is-done-for
la pagina dei manga può essere vista anche così
https://www.mangaworld.in/manga/`numero id` quindi può essere ottenuto ciclando da 170 a 2332 (ultimo attuale). così da saltare le query all archivio

```
"o":
    "w":[
         [
            "s0-2",
            0,
            "manga":{
                // Uguale a "manga" del json precedente
            },
         ]
         3[
            "s0-83",
            3,
            {
                "pages":{
                    "volumes":[ // se è manga normale
                        {
                            "volume":{
                                "_id":"5fb01fa0c34214221761baef",
                                "manga":"5faa2373de75a27e9df97f61",
                                "name":"Volume 04",
                                "image":"/volumes/5fb01fa0c34214221761baef.png",
                            },
                            "chapters":[
                                {
                                "_id":"5fd1f736ea2ffd76512a3647",
                                "pages":[
                                        "1.png",
                                        "2.png",
                                        "3.png",
                                        "4.png"
                                    ],
                                "name":"Capitolo 33.5",


                                },
                                ...
                            ],
                        },
                        ...
                    ],
                    "singleChapters":[ // se è oneshot
                        {
                        "_id":"5fc7aedae951863e97b3f121",
                        "pages":[
                           "1.jpg",
                           "2.jpg",
                            ],
                        "name":"Oneshot",

                        },
                        ...
                    ],
                },
            },
         ],

```