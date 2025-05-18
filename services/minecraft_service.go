package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"Bohrium/models"
)

type MinecraftService struct{}

func NewMinecraftService() *MinecraftService {
	return &MinecraftService{}
}

func (s *MinecraftService) GetUUID(username string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("player not found")
	}

	var profile models.ProfileResponse
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return "", err
	}

	return profile.ID, nil
}

func (s *MinecraftService) GetSkinURL(uuid string) (string, error) {
	url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("profile not found")
	}

	var skinResp models.SkinResponse
	if err := json.NewDecoder(resp.Body).Decode(&skinResp); err != nil {
		return "", err
	}

	for _, prop := range skinResp.Properties {
		if prop.Name == "textures" {
			decoded, err := base64.StdEncoding.DecodeString(prop.Value)
			if err != nil {
				return "", err
			}
			var textures struct {
				Textures struct {
					SKIN struct {
						URL string `json:"url"`
					} `json:"SKIN"`
				} `json:"textures"`
			}
			if err := json.Unmarshal(decoded, &textures); err != nil {
				return "", err
			}
			return textures.Textures.SKIN.URL, nil
		}
	}
	return "", fmt.Errorf("no skin found")
}

func (s *MinecraftService) GetHeadURL(uuid string) string {
	return fmt.Sprintf("https://crafatar.com/avatars/%s", uuid)
} 
