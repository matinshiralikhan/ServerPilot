package agent

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

// LoadSSHKey loads an SSH private key from file
func LoadSSHKey(path string) ssh.Signer {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v", err)
	}

	return signer
}
