# ASCII Art Web

A simple web application written in Go that converts user input into ASCII art using different banner styles.

## Features

- Web interface for submitting text and selecting a banner style
- Renders ASCII art from user input
- Custom error pages for invalid requests
- Serves static file Css from the `static` directory


### Running the Server

1. Clone the repository.
2. Make sure you are in the project root directory.
3. Run:

   ```sh
   go run ascii-art-web/main.go
   ```

4. Open your browser and go to [http://localhost:8080](http://localhost:8080)

### Usage

- Enter your text and select a banner style on the main page.
- Submit the form to see your text rendered as ASCII art.

## Notes

- Static files are served from the `static` directory.
- Templates are in the `templates` directory.
- Error handling uses a custom error page (`error.html`).

## Authors

- aoutrgua
- boulhaj
- melghama

