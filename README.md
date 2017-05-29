[![Build Status](https://travis-ci.org/danmademe/geoip-api.svg?branch=master)](https://travis-ci.org/danmademe/geoip-api) [![Coverage Status](https://coveralls.io/repos/github/danmademe/geoip-api/badge.svg?branch=master)](https://coveralls.io/github/danmademe/geoip-api?branch=master)
# geoip-api
I convert ip's into countries and stuff


# Usage
```
./geoip-api --lang languages.json
```

**Lang is Required**

# Flags

* **lang** --- location of local language.json
* **mmdb** - location of local .mmdb file
* **gzdb** -- location of local .gzip file
* **dburl** -- location of remote file (can be mmdb or gzip)


# Language JSON

```json
[
    { "language": "en", "country": "US" },
    { "language": "en", "country": "CA" },
    { "language": "en", "country": "AU" },
    { "language": "en", "country": "GB" },
    { "language": "en", "country": "NO" },
    { "language": "es", "country": "MX" },
    { "language": "es", "country": "ES" }
]

```

# Install using
```sh
curl -Ls https://raw.githubusercontent.com/danmademe/geoip-api/master/install.sh | sudo -H sh
```
