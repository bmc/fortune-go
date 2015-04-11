# Fortune program in Go

## Intro

_fortune_ is a stripped-down implementation of the classic BSD Unix _fortune_ 
command: It selects a random fortune cookie saying from a file and displays it
on standard output.

Unlike the BSD fortune command (or my own Python version, at 
<https://github.com/bmc/fortune>, this version does not use an index file. We 
have loads of memory these days, and fortunes files aren't that big, so it's 
feasible to load the whole text file in memory, parse it on the fly, and 
randomly choose a resulting fortune.

This version is written in Go, mostly because, why not?

## Installation

The easiest way to install this program is via the `go` command. Make sure
your `GOPATH` is set to something, and ensure that `$GOPATH/bin` is in your
`PATH`. Then:

    go get github.com/bmc/fortune-go

A program called `fortune-go` will end up in your `$GOPATH/bin` directory.

## Usage

    fortune [/path/to/fortune/cookie/file]
    fortune -h|--help

If you don't specify a fortune cookie file path (see below), _fortune_
defaults to the contents of the `FORTUNE_FILE` environment variable. If
neither the argument nor the `FORTUNE_FILE` variable is present, _fortune_
aborts.

## Fortune Cookie File Format

A fortune cookie file is a text file full of quotes. The format is simple:
The file consists of paragraphs separated by lines containing a single '%'
character. For example::

    A little caution outflanks a large cavalry.
        -- Bismarck
    %
    A little retrospection shows that although many fine, useful software
    systems have been designed by committees and built as part of multipart
    projects, those software systems that have excited passionate fans are
    those that are the products of one or a few designing minds, great
    designers. Consider Unix, APL, Pascal, Modula, the Smalltalk interface,
    even Fortran; and contrast them with Cobol, PL/I, Algol, MVS/370, and
    MS-DOS.
        -- Fred Brooks, Jr.
    %
    A man is not old until regrets take the place of dreams.
        -- John Barrymore

You're more than welcome to _my_ fortune cookie file. It's over here:
<https://github.com/bmc/fortunes>.

## License and Copyright

See the accompanying License file.
