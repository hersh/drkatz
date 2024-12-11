package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

// Structs for the YAML data
type Highlight struct {
	Link string        `yaml:"link"`
	Text template.HTML `yaml:"text"`
}

type ThumbnailRow struct {
	Thumbnail  string       `yaml:"thumbnail"`
	Guest      string       `yaml:"guest"`
	Highlights []*Highlight `yaml:"highlights"`
}

type Episode struct {
	RowClass string          `yaml:"rowClass"`
	Title    string          `yaml:"title"`
	Link     string          `yaml:"link"`
	Date     string          `yaml:"date"`
	Code     string          `yaml:"code"`
	RowSpan  int             `yaml:"rowSpan"`
	Rows     []*ThumbnailRow `yaml:"rows"`
}

type SiteData struct {
	Episodes []*Episode `yaml:"episodes"`
}

func fluffUpSiteData(data *SiteData) {
	evenOdd := []string{"even", "odd"}
	for index, episode := range data.Episodes {
		episode.RowClass = evenOdd[index%2]
		episode.RowSpan = len(episode.Rows)
	}
}

func main() {
	// Load the YAML file
	yamlFile, err := ioutil.ReadFile("site_data.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var data SiteData
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	fluffUpSiteData(&data)

	// Parse the HTML template
	tmpl, err := template.ParseFiles("site.template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Create the output HTML file
	outputFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Execute the template with the data
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Println("Site generated successfully: index.html")
}
