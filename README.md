
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

## Download

 * [Linux amd64](../../releases/latest/download/gpgencrypt-linux-amd64.tar.gz)

 * [MacOS](../../releases/latest/download/gpgencrypt-darwin-amd64.tar.gz)

  * [Windows](../../releases/latest/download/gpgencrypt-windows-amd64.tar.gz)


On MacOS, open the download archive, ctrl-click `gpgencrypt` and open the application to accept running the application.  You can then use it from the command line.
