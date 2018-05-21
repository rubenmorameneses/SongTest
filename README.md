# SongTest
Test api for aplying to beenverified

## Instalación
If you want to compile and run, you have to download the [compiler] (https://golang.org/dl/), then to write code in it, I recommend downloading the IDE [LiteIDE] (https://github.com/ visualfc / liteide), obviously also a MySQL database instance, because this IDE is specially made for Golang and we need to install the following libraries:

- [Gorilla Mux](https://github.com/gorilla/mux): Router API. Just run in cmd:

> go get github.com/gorilla/mux

- [Gonfig](https://github.com/tkanos/gonfig): JSON configuration file management. Just run in cmd:

> go get github.com/tkanos/gonfig

- [MySQLDriver](https://github.com/go-sql-driver/mysql): MySQL driver. Just run in cmd:

> go get github.com/go-sql-driver/mysql

As Golang does not have a dependency manager such as Maven or NPM or at least a confidence 3th party library, it is necessary to install the previously described libraries, finally to download the project:
>  git clone https://github.com/rubenmorameneses/SongTest.git

## Usage
The API consists of 3 calls, to consult songs either by title, genre or artist, when the code is running, you can use Postman like example to call with the following API calls sample:

- ip:port/songsByArtist/<Artist> (GET method)
- ip:port/song/<tittle_song> (GET method)
- ip:port/songByGenre/<genre> (GET method)
- ip:port/songByGenre/<start>/<top> (GET method)

It also has a configuration file (Configuration.json), in which we can change the properties to connect to the desired database:

> {
>    "Port": ":8080",
>	 "Pass": "admin",
>	 "DatabaseName": "Songs",
>	 "UserName": "admin"
> }
