package SimpleSSH

import (
	"bytes"
	"log"

	"golang.org/x/crypto/ssh"
)

type SimpleSSH struct {
	hostname string
	username string
	password string
	config   *ssh.ClientConfig
	client   *ssh.Client
}

func New(hostname string, username string, password string) SimpleSSH {
	s := SimpleSSH{hostname, username, password, nil, nil}

	s.configure()
	s.connect()
	return s
}

func (s *SimpleSSH) Cleanup() {
	s.client.Close()
}

func (s *SimpleSSH) configure() {
	config := &ssh.ClientConfig{
		User: s.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.password),
		},
	}
	s.config = config
}

func (s *SimpleSSH) connect() {
	client, err := ssh.Dial("tcp", s.hostname+":22", s.config)
	if err != nil {
		log.Fatal("Unable to connect: " + err.Error())
	}
	s.client = client
}

func (s *SimpleSSH) initialize() *ssh.Session {
	session, err := s.client.NewSession()
	if err != nil {
		log.Fatal("Unable to create session: " + err.Error())
	}
	return session
}

func (s *SimpleSSH) Run(command string) string {
	session := s.initialize()
	defer session.Close()

	var out bytes.Buffer
	session.Stdout = &out
	err := session.Run(command)
	if err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	return out.String()
}
