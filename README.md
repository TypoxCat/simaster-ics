# Simaster ICS Generator

Convert saved SIMASTER class and exam schedules from HTML to ICS files for
calendar applications such as Google Calendar. The web converter runs entirely
in the browser; uploaded files are not sent to a server.

## This fork

This project is based on [SimasterICSGen](https://github.com/refeed/SimasterICSGen)
by `refeed` and is licensed under the Apache License 2.0. This fork adds support
for newer SIMASTER HTML formats, a redesigned minimal interface, improved file
download handling, and a server fix for serving the web application.

## How to use

### CLI

```
Usage of SimasterICSGen:
  -input string
        (Mandatory) The HTML file of a SIMASTER schedule page
  -type string
        Schedule type: exam or class
  -output string
        The ICS output (default "result.ics")
```

Steps:
1. Download the executable from this repository's releases page.
2. Go to the relevant SIMASTER class or exam schedule page.
3. Save the page (press `CTRL` + `S`)
4. Run the converter, for example:
```
./simastericsgen -input "Simaster Jadwal Ujian.html" -type exam
```
5. The result will be saved as `result.ics`.

And finally, you can import the `result.ics` file from a calendar app like Google
Calendar.

### Web version

The web version is a static browser app, so other users only need to open the
published GitHub Pages URL. No Go installation, CLI tool, or server is needed
for end users.

To deploy it:

1. Push this repository to GitHub using the `main` branch.
2. In **Settings > Pages**, select **Deploy from a branch**.
3. Select the `gh-pages` branch and the `/ (root)` folder, then save.
4. Wait for the `Deploy web app to GitHub Pages` action to finish.

The app will be available at:

```text
https://<your-username>.github.io/<repository-name>/
```

Every push to `main` rebuilds and republishes the web app automatically. Users
can open that URL, upload their saved SIMASTER HTML file, and download the
generated ICS file directly from their browser.

To build and preview it locally:

```sh
make buildpages
cd output
python3 -m http.server 8080
```

Then open <http://localhost:8080>. The generated filename follows the uploaded
HTML filename, for example `Jadwal Kuliah.html` becomes `Jadwal Kuliah.ics`.

## License

This project is distributed under the Apache License 2.0. See [LICENSE.txt](LICENSE.txt).
