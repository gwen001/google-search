<h1 align="center">google-search</h1>

<h4 align="center">Performs searches on Google and display the resulting URLs, as simple as that!</h4>

<p align="center">
    <img src="https://img.shields.io/badge/go-v1.13-blue" alt="go badge">
    <img src="https://img.shields.io/badge/python-v3-blue" alt="python badge">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="MIT license badge">
    <a href="https://twitter.com/intent/tweet?text=https%3a%2f%2fgithub.com%2fgwen001%2fgoogle-search%2f" target="_blank"><img src="https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fgwen001%2Fgoogle-search" alt="twitter badge"></a>
</p>

<!-- <p align="center">
    <img src="https://img.shields.io/github/stars/gwen001/google-search?style=social" alt="github stars badge">
    <img src="https://img.shields.io/github/watchers/gwen001/google-search?style=social" alt="github watchers badge">
    <img src="https://img.shields.io/github/forks/gwen001/google-search?style=social" alt="github forks badge">
</p> -->

---

## Important note

‼ Google searches are unstable as f**k ‼

## Requirement

This script uses the tricky method developed s0md3v (see [goop by s0md3v](https://github.com/s0md3v/goop)) so it requires a Facebook cookie.
Yeah I know, sounds weird but it works!

For this purpose I use a Facebook test account and set the cookie in an environment variable:  
(don't even try that one it's fake 😉)  
```
export FACEBOOK_COOKIE="datr=q1V0Y8TgRJBF11vXCB2cl; sb=tEV2Y6okO_z0DRjFJ7jjEsXk; c_user=10409039010190; xs=16%3A1utfaZis3V5yq%3A2%3A16695433%3A-1%3A-1; fr=0rB0KLvdL4UqePy.AXXorCVM%HkrQkP7L4_VLSExs.BjW0.ZW.AA.0.0.BjkW4.AWm2gIsce6; wd=1680x937; dpr=4; presence=C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A16686438%2C%22v%22%3A1%7D"
```

## Install

```
go install github.com/gwen001/google-search@latest
```

or

```
git clone https://github.com/gwen001/google-search
cd google-search
go install
```

## Usage

```
Usage of google-search:
  -c string
    	your facebook cookie
  -d	debug mode
  -s string
    	search term you are looking for (required)
```

---

<img src="https://raw.githubusercontent.com/gwen001/google-search/master/preview.gif" />

---

## Install (python version)

```
git clone https://github.com/gwen001/google-search
cd google-search
pip3 install -r requirements.txt
```

## Usage (python version)

```
$ python3 google-search.py -t "site:10degres.net"
```

```
usage: google-search.py [-h] [-s SEARCH] [-d] [-c FBCOOKIE]

options:
  -h, --help            show this help message and exit
  -s SEARCH, --search SEARCH
                        search term (required)
  -d, --decode          urldecode the results
  -c FBCOOKIE, --fbcookie FBCOOKIE
                        your facebook cookie
```

---

Feel free to [open an issue](/../../issues/) if you have any problem with the script.  

