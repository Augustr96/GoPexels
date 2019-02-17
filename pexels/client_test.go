package pexels

import (
	"os"
	"testing"
)

var TOKEN = os.Getenv("PexelsToken")

var c = NewClient(TOKEN)

func TestClient_SearchPhotos(t *testing.T) {
	result, err := c.SearchPhotos("Sea", 15, 1)
	if err != nil {
		t.Fatalf("search error: %v", err)
	}
	if result.Page == 0 {
		t.Fatal("search result wrong")
	}
}

func TestClient_CuratedPhotos(t *testing.T) {
	result, err := c.CuratedPhotos(15, 1)
	if err != nil {
		t.Fatalf("get curatedPhotos error: %v", err)
	}
	if result.Page ==  0 {
		t.Fatal("curated photo content wrong")
	}
}

func TestClient_GetPhoto(t *testing.T) {
	result, err := c.GetPhoto(10000)
	if err != nil {
		t.Fatalf("get photo error: %v", err)
	}
	if result.Id != 10000 {
		t.Fatal("get photo content wrong")
	}
}

func TestClient_GetRandomPhoto(t *testing.T) {
	result, err := c.GetRandomPhoto()
	if err != nil {
		t.Fatalf("get random photo error: %v", err)
	}
	if result.Id == 0 {
		t.Fatal("get random photo wrong")
	}
}

func TestClient_SearchVideo(t *testing.T) {
	result, err := c.SearchVideo("sea", 15, 1)
	if err != nil {
		t.Fatalf("search videos error: %v", err)
	}
	if result.Page == 0 {
		t.Fatal("search videos content wrong")
	}
}

func TestClient_PopularVideo(t *testing.T) {
	result, err := c.PopularVideo(15, 1)
	if err != nil {
		t.Fatalf("get popular videos error: %v", err)
	}
	if result.Page == 0 {
		t.Fatal("get popular videos wrong")
	}
}

func TestClient_GetRandomVideo(t *testing.T) {
	result, err := c.GetRandomVideo()
	if err != nil {
		t.Fatalf("get random video error: %v", err)
	}
	if result.Id == 0 {
		t.Fatal("get random video wrong")
	}
}
