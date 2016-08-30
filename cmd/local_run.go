package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/spf13/cobra"
)

var serviceType string

func checkLocalRunning(name string) (exists bool) {
	client, _ := GetDockerClient()
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})

	for _, container := range containers {
		if container.Names[0][1:] == name {
			exists = true
		}
	}

	return exists
}

// localRunCmd allows users to execute arbitrary bash commands within a container.
var localRunCmd = &cobra.Command{
	Use:   "run",
	Short: "run a command in an app container.",
	Long:  `Execs into container and runs bash commands.`,
	Run: func(cmd *cobra.Command, args []string) {

		if cfg.ActiveApp == "" || cfg.ActiveDeploy == "" {
			log.Fatalln("Must set ActiveApp and ActiveDeploy in drud.yaml. Use config set to achieve this.")
		}

		if len(args) < 1 {
			log.Fatalln("Must pass a command as first argument.")
		}

		if appClient == "" {
			appClient = cfg.Client
		}

		cmdSplit := strings.Split(args[0], " ")

		basePath := path.Join(homedir, ".drud", appClient, cfg.ActiveApp, cfg.ActiveDeploy)
		nameContainer := fmt.Sprintf("%s-%s-%s-%s", appClient, cfg.ActiveApp, cfg.ActiveDeploy, serviceType)

		if !checkLocalRunning(nameContainer) {
			log.Fatal("App not runnign locally. Try `drud local add`.")
		}

		composeLOC := path.Join(basePath, "docker-compose.yaml")
		if _, err := os.Stat(composeLOC); os.IsNotExist(err) {
			log.Fatalln("No docker-compose yaml for this site. Try `drud local add`.")
		}

		cmdArgs := []string{
			"-f", composeLOC,
			"exec",
			"-T", nameContainer,
		}
		cmdArgs = append(cmdArgs, cmdSplit...)

		out, err := exec.Command("docker-compose", cmdArgs...).CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Errorf("%s - %s", err.Error(), string(out)))
		}

		fmt.Println(string(out))

	},
}

func init() {
	localRunCmd.Flags().StringVarP(&appClient, "client", "c", "", "Client name")
	localRunCmd.Flags().StringVarP(&serviceType, "service", "s", "web", "Which service to send the command to. [web, db]")
	LocalCmd.AddCommand(localRunCmd)

}