# Dr Katz Professional Therapist fan page source

This repo is the source code for https://neocities.org/chafecity/drkatz, a fan page
listing all Dr. Katz animated episodes, along with guests, and links to highlights.

This uses a simple static site generator (generate_site.go). The data is stored in
site_data.yaml, and the html template is in site.template.html.

To build:

    go run generate_site.go

To build and immediately view:

    go run generate_site.go && google-chrome index.html

Last tested with go1.23.3 linux/amd64
