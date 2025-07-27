# ASCII Art Web Dockerize

This project demonstrates how to containerize a Go web application that generates ASCII art using Docker.

## Docker Features

- Multi-stage Docker build for small, secure images
- Copies Go binary and required assets (`templates`, `static`) into the final image
- Exposes port 8080 for web access

## Quick Start

### Build the Docker Image

```sh
docker build -t ascii-art-web .
```

### Run the Container

```sh
docker run -d -p 8080:8080 ascii-art-web
```

Open [http://localhost:8080](http://localhost:8080) in your browser.

## Project Structure

- `Dockerfile` – Docker build instructions
- `main.go` – Go web server source code
- `templates/` – HTML templates
- `static/` – Static assets (CSS, JS)

## How It Works

- The Dockerfile uses a multi-stage build to compile the Go app and copy only the necessary files to a minimal Alpine image.
- The server runs on port 8080 inside the container.
- Static and template files are included for full web functionality.

## Authors

- boulhaj
- mhilli