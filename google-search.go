package main

import (
	"os"
	"fmt"
	"time"
	"flag"
	"regexp"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
)

type Config struct {
	debug bool
	search string
	fbcookie string
	q_sleep time.Duration
	q_timeout time.Duration
	q_num int
	max_fail int
}

var config = Config{}

func config_init() {
	config.q_sleep = 300
	config.q_timeout = 5
	config.q_num = 100
	config.max_fail = 2
}


func decode_html(str string) string {
    var c_encoded = []string{"&gt;", "&lt;", "&quot;", "&amp;", "&#039;"}
    var c_decoded = []string{">", "<", "\"", "&", "'"}
	for k,_ := range c_encoded {
		str = strings.Replace( str, c_encoded[k], c_decoded[k], -1 )
	}
    return str
}


func parse(str string) []string {
	var t_links []string
	var rgxp = regexp.MustCompile( `<div class="[^"]+"><a href="/url\?q=(.+?)&sa=[^"]+">` )
	// var t_match = rgxp.FindAll([]byte(str), -1)
	var t_match = rgxp.FindAllStringSubmatch(str, -1)
	if config.debug {
		fmt.Printf("%d results\n", len(t_match))
	}

	if len(t_match) > 0 {
		for _,match := range(t_match) {
			t_links = append(t_links, match[1])
		}
	}

	return t_links
}


func doSearch(search string, page int) int {

	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()

	var gg_offset = page * config.q_num
	var gg_url = fmt.Sprintf("https://google.com/search?q=%s&num=%d&start=%d&filter=1", url.QueryEscape(search), config.q_num, gg_offset )
	if config.debug {
		fmt.Printf( "-> %s\n", gg_url)
	}

    var fb_url = fmt.Sprintf("https://developers.facebook.com/tools/debug/echo/?q=%s", url.QueryEscape(gg_url) )
	client := http.Client{ Timeout: time.Second * config.q_timeout }

	req, err := http.NewRequest("GET", fb_url, nil)
	if err != nil {
		fmt.Printf( "error: %s\n", fb_url )
		return 0
	}

	req.Header.Set("Host", "developers.facebook.com")
	req.Header.Set("User-Agent", "developers.facebook.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", config.fbcookie)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("TE", "Trailers")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Printf( "error: %s\n", fb_url )
		return 0
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Printf( "error: %s\n", fb_url )
	}

	var decoded_body = decode_html( decode_html( string(body) ) )
	// fmt.Printf(decoded_body)
	var t_links = parse( decoded_body )

	if len(t_links) > 0 {
		for _,link := range(t_links) {
			fmt.Println(link)
		}
	}

	return len(t_links)
}


func main() {

	flag.StringVar( &config.fbcookie, "c", "", "your facebook cookie" )
	flag.BoolVar( &config.debug, "d", false, "debug mode" )
	flag.StringVar( &config.search, "s", "", "search term you are looking for (required)" )
	flag.Parse()

	if config.search == "" {
		flag.Usage()
		fmt.Printf("\nsearch not found\n")
		os.Exit(-1)
	}

	if config.fbcookie == "" {
		config.fbcookie = os.Getenv("FACEBOOK_COOKIE")
	}
	if config.fbcookie == "" {
		flag.Usage()
		fmt.Printf("\nfacebook cookie not found\n")
		os.Exit(-1)
	}

	var page = 0
	var n_fail = 0

	config_init()

	for run:=true; run; {
		time.Sleep( config.q_sleep * time.Millisecond )

		var n_links = doSearch( config.search, page )

		if n_links <= 0 {
			n_fail++
		}

		if n_fail >= config.max_fail {
			run = false
			break
		}

		page++
	}
}
