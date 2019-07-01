package main

import (
	"github.com/n7down/PITFTDisplays/internal/config"
	"github.com/n7down/PITFTDisplays/internal/display"
	"github.com/n7down/PITFTDisplays/internal/display/spacexdisplay"
	"time"
	//"github.com/n7down/PITFTDisplays/internal/display/githubdisplay"
	log "github.com/sirupsen/logrus"
)

func main() {
	_, err := config.Config()
	if err != nil {
		log.Error(err)
	}

	displayManager := display.NewDisplayManager()
	displayManager.AddDisplay(spacexdisplay.NewSpaceXDisplay())
	//displayManager.AddDisplay(githubdisplay.NewGithubDisplay(c))
	// TODO: render every second
	for {
		time.Sleep(time.Second)
		displayManager.Render()
	}
}