# ASCII Art Web Application

## Description

This project is a web-based ASCII Art generator developed in Go. It allows users to input text and select a font style ("standard", "shadow", or "thinkertoy") to generate ASCII art dynamically through a browser interface. The application is built with modular handlers for serving the homepage, processing ASCII art generation, and serving static assets. Error handling and input validation are included to ensure a robust user experience.

---

## Authors

- hamza ghomari
- ilyas mohammed abid

---

## Usage: How to Run

1. **Clone the repository** and make sure you are in the project root directory.
2. **Install Go** (if not already installed).
3. **Run the server:**
    ```sh
    go run main.go
    ```
4. **Open your browser and visit:**
    ```
    http://localhost:8080
    ```
5. **Usage:**
    - Enter your text (up to 10,000 characters).
    - Choose a font style from the dropdown ("standard", "shadow", "thinkertoy").
    - Submit the form to view the ASCII art result.

---

## Implementation Details: Algorithm

- **Routing and Handlers:**
    - `/` (GET): Serves the main HTML page with form for user input.
    - `/ascii-art` (POST): Receives text and font selection, validates input, generates ASCII art via the `asciart.Fmain` function, and renders the result.
    - `/static/`: Serves static files such as CSS.
    - Error pages are rendered for 400, 404, 405, and server errors.

- **Input Processing:**
    - Checks that input length does not exceed 10,000 characters.
    - Handles only POST requests for ASCII art generation endpoint.
    - Sanitizes and processes form values.

- **ASCII Art Generation:**
    - Delegates text and font to `asciart.Fmain(text, banner)` which returns the ASCII art string and HTTP status code.
    - Renders the ASCII art on the page using Go's `html/template` package.

- **Error Handling:**
    - Centralized `render` function displays error messages and HTTP status codes based on error type (400, 404, 405, 500).

- **Static Files:**
    - Static assets (e.g., CSS) are served from the `/static/` route.
    - Security checks ensure only files (not directories) are served, and missing files return proper error codes.

- **Performance Logging:**
    - Logs the time taken to process ASCII art generation requests.

---