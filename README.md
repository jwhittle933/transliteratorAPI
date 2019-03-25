## Transliterator API

A replacement for previous version

## File Tree

```bash
├── README.md
├── controllers
│   ├── TransliterateController.go
│   └── UploadController.go
├── engines
│   └── engine.go
├── init
│   └── Init.go
├── middlewarehelpers
│   └── MiddlewareHelpers.go
├── scripts
│   ├── basicSetup.sh
│   └── setup.sh
└── transliterate.go
```

## Setup Scripts

Scripts are found in the `scripts/` directory. To start from scratch run `./scripts/basicSetup.sh`. If you've set go env variables in the past, you should run `./scripts/setup.sh`. This script will check for directories and go env vars, and set them if needed.

## API Endpoints
