# Dr Katz Professional Therapist fan page source

This repo is the source code for https://chafecity.neocities.org/drkatz, a fan page
listing all Dr. Katz animated episodes, along with guests, and links to highlights.

## Contributions welcome!

* Your favorite Dr. Katz highlights
* Links to better-quality videos
* Layout/usability issues or improvements
* Suggestions for other show-related content

To contribute, you can either:

* [Create an issue](https://github.com/hersh/drkatz/issues/new) with details of the change
  you'd like to see, and I'll see what I can do, *or*

* [Create a pull request](https://github.com/hersh/drkatz/compare) if you want to write the
  changes yourself. The main table content is all in site_data.yaml. The web page HTML is in
  site.template.html.

### Highlight Guidelines:

* For highlight links, please use the same youtube link this page does, but use the "Copy
  video URL at current time" feature to get a link directly to the right part of the video.

* For highlight text, try to write text that makes it easy to find what you're looking for
  without giving away the punchlines.

## Build and test the site locally (from linux)

This uses a simple static site generator (generate_site.go). The data is stored in
site_data.yaml, and the html template is in site.template.html.

To build:

    go run generate_site.go

To build and immediately view:

    go run generate_site.go && google-chrome index.html

Last tested with go1.23.3 linux/amd64
