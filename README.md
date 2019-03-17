# sitemapaths

**sitemapaths** lists paths on `sitemap.xml`.

## Install

`go get -u github.com/hareku/sitemapaths/cmd/sitemapaths`

## Usage

`sitemapaths [global options] [sitemap.xml URL]`

```
GLOBAL OPTIONS:
   --separator value  separator of paths for output (default: "\n")
   --help, -h         show help
   --version, -v      print the version
```

### Simple usage

```
$ sitemapaths https://example.com/sitemap.xml
/
/first-post
/second-post
```

### With separator usage

```
$ sitemapaths --separator " " https://example.com/sitemap.xml
/ /first-post /second-post
```
