## Transliterator API

A replacement for previous version

## File Tree

```bash
.
├── README.md
├── controllers.go
├── engine.go
├── init.go
├── main.go
├── middlewarehelpers
│   └── MiddlewareHelpers.go
├── pdfReader
│   └── pdfReader.go
├── saves
├── scripts
│   ├── govars.sh
│   └── govarsBasic.sh
├── testfiles
│   ├── aleph.txt
│   ├── alpha.txt
│   ├── greek.pdf
│   ├── greekLowercaseLetters.txt
│   ├── hebrew.docx
│   ├── hebrewLetters.txt
│   ├── test.pdf
│   └── test.txt
├── tmp
└── types.go
```

## Setup Scripts

Scripts are found in the `scripts/` directory. To start from scratch run `./scripts/basicSetup.sh`. If you've set go env variables in the past, you should run `./scripts/setup.sh`. This script will check for directories and go env vars, and set them if needed.

## API Endpoints
