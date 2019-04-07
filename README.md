## Transliterator API

[![Build Status](https://travis-ci.com/jwhittle933/transliteratorAPI.svg?branch=master)](https://travis-ci.com/jwhittle933/transliteratorAPI)

## File Tree

```bash
.
├── LICENSE
├── README.md
├── controllers
│   ├── transliterate.go
│   └── upload.go
├── controllers.go
├── engine
│   └── engine.go
├── init
│   └── init.go
├── main.go
├── middleware
│   └── middleware.go
├── pdfReader
│   └── pdfReader.go
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
└── types
    └── types.go
```

## Setup Scripts

Scripts are found in the `scripts/` directory. To start from scratch run `./scripts/basicSetup.sh`. If you've set go env variables in the past, you should run `./scripts/setup.sh`. This script will check for directories and go env vars, and set them if needed.

## API Endpoints
