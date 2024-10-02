package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//el parametro tiene que cambiar segun la pagina que quieras visitar
	blogTitle, err := GetLatestBlogTitles("https://gacel.cl/sandalias/?srsltid=AfmBOophxxi9XFYVEmwOAI-X2pKPEya4zi4tb-5sBDXmjvtcr5Q4vUMH")
	if err != nil {
		log.Println("Error fetching blog titles:", err)
		return
	}

	fmt.Println("Blog titles:")
	fmt.Println(blogTitle)
}

func GetLatestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch data: %v", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var titles strings.Builder
	//En la funcion buscar de `doc.Find` tiene que llevar el selector de lo que te interesa buscar en la pagina, pueden ser precios, titulos, etc...
	doc.Find(".price").Each(func(i int, s *goquery.Selection) {
		titles.WriteString("-" + s.Text() + "\n")
	})

	return titles.String(), nil
}
