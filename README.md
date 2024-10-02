# Web Scraper para Extraer Precios y Títulos

Este programa en Go extrae precios y otros elementos HTML de una página web utilizando la biblioteca `goquery`. 

## Descripción

El scraper busca elementos HTML con la clase `.price` en la página especificada y devuelve los precios encontrados. Puedes modificar el selector para extraer otros tipos de información según tus necesidades.

## Requisitos

- [Go](https://golang.org/doc/install) instalado en tu máquina.
- La biblioteca `goquery`. Puedes instalarla usando el siguiente comando:

bash
go get -u github.com/PuerkitoBio/goquery

## USO
Para extraer información de una página diferente, debes cambiar la URL en la función main. Abre el archivo main.go en un editor de texto y busca la siguiente línea:
`blogTitle, err := Getanything("URLEXAMPLE")`

Además, cambia la etiqueta con la cual se identifica en el HTML de la página web en la siguiente línea:
`doc.Find(".price").Each(func(i int, s *goquery.Selection) {
		titles.WriteString("-" + s.Text() + "\n")
	})`
