Abstracticon
============

This is a toy project written to play with the image package from Go. It
includes a go gettable library and a cli tool.

Installation
------------

```console
$ go get github.com/mhcerri/abstracticon
```

Usage
-----

```console
$ abstracticon -h
usage: abstracticon [options]
Reads data from stdin and generates an icon. Options:
  -hash string
    	Hash function. Options: md5 (default "md5")
  -mirrored
    	Mirror the image left to right. (default true)
  -multiplier int
    	Number of pixels that represents a point. (default 8)
  -output string
    	Output file. (default "output.png")
  -points int
    	Number of points in the height and width of the image. (default 8)
  -transparent
    	Use transparent background.
```

Example
-------

```console
$ echo -n 'example' | abstracticon -output example.png
$ display example.png
```

![example.png](/example.png?raw=true)
