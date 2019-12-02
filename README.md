# SGo-Scraper
Download an entire selected Suicide Girls album.
Just run by appending the URL of the album.
- Automatically creates and sort download paths.
- Hopefuls compatible.

## Requirements
- Suicide Girls account.
- Go 1.9

## Usage
```
git clone https://github.com/Osuka42g/SGo-Scraper.git
cd SGo-Scraper
go get
cp .env.example .env
```

Open `.env` and fill _SESSIONIDTOKEN_ with [your own Token](#getting-token).

```
go build
./SGo-Scraper http://suicidegirls.com/full-url-to-the-suicidegirls-album
```

If you want to also compress the downloaded files, append `-z` at the end as argument.

## Getting Token
1. Using Google Chrome, log in into your Suicide Girls account.
2. Pop out the developers console and go to _Application_ tab.
3. At the left, go to Storage -> Cookies -> https://suicidegirls.com
4. Scroll down until you find _sessid_ cookie.
5. Copy the value from _Value_ column.

##### Why this is needed?
In order to access to full albums, credentials to the site are required.
Current implementation cannot just login to Suicide Girls because the login is captcha protected; so we are accesing the crawler with the cookie from our login.

🍻