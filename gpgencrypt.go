package main

/**
  Based on :
  https://gist.github.com/stuart-warren/93750a142d3de4e8fdd2#file-simple-gpg-enc-go
  https://gist.github.com/ayubmalik/a83ee23c7c700cdce2f8c5bf5f2e9f20
*/

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"

	"os"
	"strings"
)

func main() {

	recipient, fileToEnc, err := initialize()
	if err != nil {
		fmt.Println(err)

		printHelp()
		return
	}

	fmt.Println(recipient.Identities)

	f, err := os.Open(fileToEnc)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	encrypt([]*openpgp.Entity{recipient}, nil, f, buf)

	fmt.Println(buf)
}

func printHelp() {
	fmt.Println("Usage:\n")
	fmt.Println("Encrypt a file using a public keyfile on disk:")
	fmt.Println("\tgpgencrypt public-keyfile file-to-encrypt\n")
	fmt.Println("Encrypt a file using a public keyfile from the clipboard (mac example):")
	fmt.Println("\tpbpaste | gpgencrypt file-to-encrypt")
}

func initialize() (*openpgp.Entity, string, error) {
	if len(os.Args) == 3 {
		recipient, err := readEntityFromFile(os.Args[1])
		if err != nil {
			printHelp()
			panic(err)
		}
		return recipient, os.Args[2], nil
	} else if len(os.Args) == 2 {
		fi, err := os.Stdin.Stat()
		if err != nil {
			printHelp()
			panic(err)
		}
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			bytes, _ := ioutil.ReadAll(os.Stdin)
			str := string(bytes)

			//log.Println(str)

			recipient, err := readEntity(str)
			if err != nil {
				fmt.Println("Error reading public key from pipe.  Input: ")
				fmt.Println(str)

				panic(err)
			}
			return recipient, os.Args[1], nil
		} else {
			printHelp()
			panic(errors.New("Either supply two arguments or one argument and piped input"))
		}
	} else {
		return nil, "", errors.New("Invalid number of arguments")
	}
}

func encrypt(recipient []*openpgp.Entity, signer *openpgp.Entity, reader io.Reader, writer io.Writer) error {
	armored, err := armor.Encode(writer, "PGP MESSAGE", make(map[string]string))
	if err != nil {
		return err
	}
	defer armored.Close()

	wc, err := openpgp.Encrypt(armored, recipient, signer, &openpgp.FileHints{IsBinary: false}, nil)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, reader); err != nil {
		return err
	}
	return wc.Close()
}

func readEntityFromFile(name string) (*openpgp.Entity, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}

func readEntity(key string) (*openpgp.Entity, error) {
	block, err := armor.Decode(strings.NewReader(key))
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}
