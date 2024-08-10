package utils

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

func ReadPostgresDetails() (host, port, user, dbName, password string, err error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return "", "", "", "", "", err
	}

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`^(POSTGRES_HOST|POSTGRES_PORT|POSTGRES_USER|POSTGRES_DB|POSTGRES_PASSWORD)=(.*)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			switch matches[1] {
			case "POSTGRES_HOST":
				host = matches[2]
			case "POSTGRES_PORT":
				port = matches[2]
			case "POSTGRES_USER":
				user = matches[2]
			case "POSTGRES_DB":
				dbName = matches[2]
			case "POSTGRES_PASSWORD":
				password = matches[2]
			}
		}
	}

	if host == "" || port == "" || user == "" || dbName == "" || password == "" {
		return "", "", "", "", "", errors.New("missing required environment postgres variables")
	}

	return host, port, user, dbName, password, nil
}

func ReadJWTSecret() (jwtSecret []byte, err error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return []byte{}, err
	}

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`^(JWT_SECRET)=(.*)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			switch matches[1] {
			case "JWT_SECRET":
				jwtSecret = []byte(matches[2])
			}
		}
	}

	if len(jwtSecret) == 0 {
		return []byte{}, errors.New("missing required environment variable JWT_SECRET")
	}

	return
}

func ReadBackendURL() (url string, err error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`^(BACKEND_URL)=(.*)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			switch matches[1] {
			case "BACKEND_URL":
				url = matches[2]
			}
		}
	}

	if len(url) == 0 {
		return "", errors.New("missing required environment variable JWT_SECRET")
	}

	return
}

func ReadGmailDetails() (email, password string, err error) {
	content, err := os.ReadFile(".env")
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`^(GMAIL_EMAIL|GMAIL_APP_PASSWORD)=(.*)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			switch matches[1] {
			case "GMAIL_EMAIL":
				email = matches[2]
			case "GMAIL_APP_PASSWORD":
				password = matches[2]
			}
		}
	}

	if email == "" || password == "" {
		return "", "", errors.New("missing required environment variables GMAIL_EMAIL and GMAIL_APP_PASSWORD")
	}

	return email, password, nil
}
