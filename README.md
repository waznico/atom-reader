![Go](https://github.com/waznico/atom-reader/workflows/Go/badge.svg?branch=master)

# Atom Reader
A simple Atom RSS reader as test project to fetch latest news of heise

# Fundamentals
This app uses spf13/cobra framework to create cli environment.
For displaying pretty tables within the CLI output I use jedib0t/go-pretty.

# Usage
This tool supports the following commands
```
lsa         - List 10 headlines of each feed which is configured in feeds.json
ls          - Lists the given amount of headlines of the given feed alias (configured in feeds.json)
describe    - Shows details of the given article of a specific feed
```

## Examples
If you want to get the 20 latest feed entries of heise Developer:
```
./atom-reader ls heise_dev 20
```

If you want to get further details to a feed entry of heise Developer:
```
./atom-reader describe heise_dev http://heise.de/-4858729
```