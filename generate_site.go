package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

// Structs for the YAML data
type Highlight struct {
	Link string        `yaml:"link" json:"link,omitempty"`
	Text template.HTML `yaml:"text" json:"text"`
}

type ThumbnailRow struct {
	Thumbnail  string       `yaml:"thumbnail" json:"-"` // Exclude from search
	Guest      string       `yaml:"guest" json:"guest,omitempty"`
	Highlights []*Highlight `yaml:"highlights" json:"highlights,omitempty"`
}

type Episode struct {
	RowClass string          `yaml:"rowClass" json:"-"` // Exclude from search
	Title    string          `yaml:"title" json:"title"`
	Link     string          `yaml:"link" json:"-"` // Exclude from search
	Date     string          `yaml:"date" json:"date"`
	Code     string          `yaml:"code" json:"code"`
	RowSpan  int             `yaml:"rowSpan" json:"-"` // Exclude from search
	Rows     []*ThumbnailRow `yaml:"rows" json:"rows"`
}

type SiteData struct {
	Episodes     []*Episode `yaml:"episodes"`
	EpisodesJSON template.JS
}

func fluffUpSiteData(data *SiteData) {
	evenOdd := []string{"even", "odd"}
	for index, episode := range data.Episodes {
		episode.RowClass = evenOdd[index%2]
		episode.RowSpan = len(episode.Rows)
	}

	// Convert episodes to JSON for client-side search
	jsonData, err := json.Marshal(data.Episodes)
	if err != nil {
		log.Fatalf("Error marshaling episodes to JSON: %v", err)
	}

	// Clean HTML entities in the JSON string to prevent template escaping issues
	jsonString := string(jsonData)
	jsonString = strings.ReplaceAll(jsonString, "\\u0026", "&")
	jsonString = strings.ReplaceAll(jsonString, "\\u003c", "<")
	jsonString = strings.ReplaceAll(jsonString, "\\u003e", ">")

	data.EpisodesJSON = template.JS(jsonString)
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
