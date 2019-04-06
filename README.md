## Transliterator API

A replacement for previous version

## File Tree

```bash
.
├── README.md
├── controllers
│   ├── transliterateController.go
│   ├── types.go
│   └── uploadController.go
├── engine
│   ├── engine.go
│   └── engineData.go
├── init
│   └── init.go
├── main.go
├── middlewarehelpers
│   └── MiddlewareHelpers.go
├── pdfReader
│   └── pdfReader.go
├── saves
│   └── dir-[207\ 203\ 53\ 127\ 122\ 18\ 68\ 70\ 152\ 247\ 172\ 215\ 25\ 117\ 166\ 110]
│       └── unzip
│           ├── [Content_Types].xml
│           ├── _rels
│           ├── docProps
│           │   ├── app.xml
│           │   └── core.xml
│           └── word
│               ├── _rels
│               │   └── document.xml.rels
│               ├── document.xml
│               ├── fontTable.xml
│               ├── settings.xml
│               ├── styles.xml
│               ├── theme
│               │   └── theme1.xml
│               └── webSettings.xml
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
└── tmp
```

## Setup Scripts

Scripts are found in the `scripts/` directory. To start from scratch run `./scripts/basicSetup.sh`. If you've set go env variables in the past, you should run `./scripts/setup.sh`. This script will check for directories and go env vars, and set them if needed.

## API Endpoints
