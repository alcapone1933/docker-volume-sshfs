package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
	"os"
)

const (
	// DefaultBasePath defines the base path within the docker plugins rootfs file system
	DefaultBasePath = "/mnt"
	// DefaultUnixSocket sets the path to the plugin socket
	DefaultUnixSocket = "/run/docker/plugins/sshfs.sock"
)

func main() {
	driver, err := newSshfsDriver(DefaultBasePath)
	if err != nil {
		log.Errorf("Failed to create the driver %s", err)
		os.Exit(1)
	}
	log.SetLevel(log.DebugLevel)
	handler := volume.NewHandler(driver)
	handler.ServeUnix(DefaultUnixSocket, 0)
}