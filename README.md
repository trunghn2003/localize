# localize

A lightweight Go library for model-based translation (multi-language) inspired by Laravel's HasTranslations.

## Features

- Store multilingual fields in a single JSONB column (PostgreSQL)
- GORM support (custom Value/Scan)
- Easy `Get(locale)` and `Set(locale, value)`
- Fallback language supported

## Install

```bash
go get github.com/trunghn2003/localize
