# identicon

## Purpose

This project exists as a way for me to experiment with using different languages to solve the same task.

Note this this is not intended to be _good_ code or even idiomatic. It's just experimentation.

Some things that will be covered:

- The language’s build/execution tooling.
- The language’s third party package tooling.
- The language's ability to create command line tools.
- Basic language structures and control flow and the following specific capabilities:
  - Reading in command line arguments.
  - Image generation.
  - File I/O.

## Approach

The code will generate a simple 5 x 5, single colour [identicon](https://en.wikipedia.org/wiki/Identicon) like this:

> ![identicon generated from the word 'biggianteye'](identicon.png)

The general approach will be:

- Read in input text from the command line.
- Generate an MD5 checksum from the input.
- Take the first three bytes as the RGB values of the colour to use.
- Divide the bytes into chunks of three and discard the last byte.
- The identicon should have bilateral symmetry, so map the bytes onto a 3 x 5 grid and then extend the grid to 5 x 5 by vertically mirroring the existing grid.
- Cells with odd bytes have a colour. Cells with even bytes are white.
- Generate an image and save it as a PNG.

## Languages

- [Elixir](elixir)
- [Go](go)
