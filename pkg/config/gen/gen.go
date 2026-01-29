package main

import (
	cfg "github.com/conductorone/baton-formal/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("formal", cfg.Config)
}
