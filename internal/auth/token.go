package auth

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func readTokenFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func TokenValidator(key string, c echo.Context) (bool, error) {
	if token, err := readTokenFromFile("token.txt"); err != nil {
		return false, err
	} else {
		return key == token, nil
	}
}
