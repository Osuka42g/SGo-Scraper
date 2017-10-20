# SGo-Scraper
Download an entire selected Suicide Girls album.

### Requirements
- Suicide Girls account.
- Golang 1.9

### Installation
```
git clone https://github.com/Osuka42g/SGo-Scraper.git
cd SGo-Scraper
go get
cp .env.example .env
```

Open `.env` and fill SESSIONIDTOKEN with your own Token.
How to

```
go build
./SGo-Scraper http://suicidegirls.com/full-url-to-the-suicidegirls-album
```

### Getting Token
Using Google Chrome, log in into your Suicide Girls account.
Pop out the developers console and go to _Application_ tab.
At the left, go to Storage -> Cookies -> https://suicidegirls.com
Scroll down until you find _sessionid_ cookie.
Copy the value from _Value_ column.

### Thanks
Started with some implementation from [Gregory Schier's Blog](https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html)