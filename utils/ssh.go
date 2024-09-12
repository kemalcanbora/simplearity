package utils

import (
	"fmt"
	"os"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type SSHClient struct {
	client *goph.Client
}

func NewSSHClient(host, user string) (*SSHClient, error) {
	fmt.Printf("Connecting to %s as %s\n", host, user)

	password, err := getPassword()
	if err != nil {
		return nil, fmt.Errorf("error reading password: %w", err)
	}

	client, err := goph.NewConn(&goph.Config{
		User:     user,
		Addr:     host,
		Port:     22,
		Auth:     goph.Password(string(password)),
		Callback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return &SSHClient{client: client}, nil
}

func (s *SSHClient) Close() error {
	return s.client.Close()
}

func (s *SSHClient) UploadFile(localPath, remotePath string) error {
	fmt.Printf("Uploading %s to %s\n", localPath, remotePath)
	err := s.client.Upload(localPath, remotePath)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func (s *SSHClient) ExecuteCommand(command string) (string, error) {
	fmt.Printf("Executing command: %s\n", command)
	output, err := s.client.Run(command)
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %w", err)
	}
	return string(output), nil
}

func getPassword() (string, error) {
	fmt.Print("Enter your password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(password), nil
}
