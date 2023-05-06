# Translate Text using Azure Cognitive Services API

This is a simple Go program that translates a given text from one language to another using the Azure Cognitive Services API.

## Prerequisites

Before running the program, make sure you have:

- An Azure account
- Azure Cognitive Services API key
- Go installed on your machine

## Installation

1. Clone the repository:

```sh
git clone https://github.com/rafaelsiq94/go-cli-translator.git
```

2. Install dependencies:

```sh
go get github.com/joho/godotenv
```
## Usage

1. Add your Azure Cognitive Services API Key and location to a `.env` file:

```sh
touch .env
echo "KEY=your_key_here" >> .env
echo "LOCATION=your_location_here" >> .env
```

2. Run the program with the origin language and target language, followed by the text to be translated:

```sh
go run main.go pt-br en "Seja bem-vindo!"
```

3. The translated text will be displayed in the terminal.
