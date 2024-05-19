# API Reference

## Endpoints

### Web

`https://akl.gg`

### API

`https://api.akl.gg`

## Auth

- Bearer token system
- Discord OAuth

## Routes

### Web

### API

#### Public

Accessible without authentication

- `/` -- banner
  - `layouts` -- ["name"] layout list
  - `layout/:name` -- {"slf"} specific layout (SLF)
  - `analyzers` -- ["name"] analyzer list
  - `analyze/:name?corpus=:corp` -- [x.y] metrics list
  - `corpora` -- ["name"] corpus list
  - `corpus/:name` -- stats about corpus?

#### Protected

Requires authentication

## Layouts

Using `https://github.com/akl-infra/slf` as type, format, and converter repository
`layouts` git submodule in `api` repo

## Corpora

`corpora/:name/:type.json` in `api` repo

## Analyzers

### Initialization

Need to load corpora for all Analyzers

#### Mini

Default analyzer, reimplementation of cmini in Go

Need to load `table.json` or embed values into source. `LoadTable()` func and `Table` global exist in `mini` package namespace.


### Metrics
