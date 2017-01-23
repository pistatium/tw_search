# tw_search
Twitter search cli tool written by go lang.

## Installing

```
go get pistatium/tw_search
go build main.go
```

## HowToUse

tw_search requires Twitter OAuth keys.

Get keys from twitter developer pages.

### Configure
OAuth keys can set by Environments or arguments.

__Set by Environments__

```bash
export TS_AT=XXXXXXXXXXXXXXX
export TS_AS=XXXXXXXXXXXXXXX
export TS_CK=XXXXXXXXXXXXXXX
export TS_CS=XXXXXXXXXXXXXXX
```
### search 

```
./main SEARCH_KEYWORD
```
