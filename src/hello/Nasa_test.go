package hello

import (
	"testing"
)

func TestAPOD(t *testing.T) {
	client := NasaClient{ApiKey: "JOdbdzShYX1MxEflQ0V0u9rNhBorfReMx4CGwg0k"}
	apod, err := client.GetAPOD()
	if apod.Date == "" {
		t.Errorf("wrong json")
	}
	if err != nil {
		t.Errorf("parse apod failed")
	}
}
