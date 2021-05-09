
## gpgencrypt

This is simple command line application that encrypts a file using a PGP public key.

## Usage

To encrypt a file using a public key file, run

    gpgencrypt public-keyfile file-to-encrypt

To encrypt a file using a public keyfile from the clipboard, run

	pbpaste | gpgencrypt file-to-encrypt

This example works on MacOs.  On Linux you can use [`xclip`](https://ostechnix.com/how-to-use-pbcopy-and-pbpaste-commands-on-linux/).

## Building

    go mod download golang.org/x/crypto
    go build
