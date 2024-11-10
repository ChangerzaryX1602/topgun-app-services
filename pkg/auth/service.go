package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		authRepository: repo,
	}
}
func GetUserDetails(accessToken string) (map[string]interface{}, error) {
	userURL := fmt.Sprintf(viper.GetString("oauth.host") + "/verify")
	bodyParams := url.Values{
		"id_token":  {accessToken},
		"client_id": {viper.GetString("oauth.client_id")},
	}
	req, err := http.NewRequest("POST", userURL, bytes.NewBufferString(bodyParams.Encode()))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error getting access token")
	}
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("response:\n", response)
	return response, nil
}
func GetAccessToken(code string) (string, error) {
	tokenURL := viper.GetString("oauth.host") + "/token"
	bodyParams := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {viper.GetString("oauth.client_id")},
		"client_secret": {viper.GetString("oauth.client_secret")},
		"code":          {code},
		"redirect_uri":  {viper.GetString("oauth.redirect_uri")},
	}
	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(bodyParams.Encode()))
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// if resp.StatusCode != http.StatusOK {
	// 	return "", fiber.NewError(fiber.StatusInternalServerError, "Error getting access token")
	// }
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("response:\n", response)
	accessToken, ok := response["id_token"].(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Access token not found in the response")
	}
	fmt.Println("accessToken:\n", accessToken)
	return accessToken, nil
}

func (s *authService) Login(code string) (*AuthResponse, error) {
	accessToken, err := GetAccessToken(code)
	if err != nil {
		fmt.Println("error get access token", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	userDetails, err := GetUserDetails(accessToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("userDetails:\n", userDetails)

	users := User{
		Email:    userDetails["email"].(string),
		Username: userDetails["name"].(string),
		Picture:  userDetails["picture"].(string),
		Sub:      userDetails["sub"].(string),
	}
	id, err := s.authRepository.Login(users.Email, users)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	AuthResponse := AuthResponse{
		ID:       id.ID,
		Email:    users.Email,
		Username: users.Username,
		Picture:  users.Picture,
		Sub:      users.Sub,
	}
	fmt.Println("AuthResponse:\n", AuthResponse)
	return &AuthResponse, nil
}
