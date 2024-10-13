package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Project struct {
	Name       string
	ConfigPath string
	Containers []Container
}

type Container struct {
	ID   string
	Name string
}

var (
	projects = make(map[string]Project)
	actions  = []string{"start", "stop"}
)

func main() {

	// docker compose --file '/job-assignment-app/docker-compose.dev.yml' --project-name 'jobs' stop
	// docker compose --file '/fastapi_sqlalchemy/docker-compose.yml' --project-name 'fastapi_sqlalchemy' start
	list := flag.Bool("list", false, "Projects list")
	name := flag.String("name", "", "Project name.")
	flag.Parse()

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		projectName := ""
		configPath := ""

		if name, ok := ctr.Labels["com.docker.compose.project"]; ok {
			projectName = name
		}
		if path, ok := ctr.Labels["com.docker.compose.project.config_files"]; ok {
			configPath = path
		}

		project, ok := projects[projectName]
		if !ok {
			project = Project{
				Name:       projectName,
				ConfigPath: configPath,
				Containers: []Container{},
			}
		}

		container := Container{
			ID:   ctr.ID,
			Name: ctr.Names[0],
		}
		project.Containers = append(project.Containers, container)

		projects[projectName] = project
	}

	if *list {
		for k, v := range projects {
			fmt.Printf("%s\t( %s )\n", k, v.ConfigPath)
		}
		return
	}

	if *name == "" {
		fmt.Println("project name should be provided. See --help")
		return
	}

	action := os.Args[len(os.Args)-1]

	if !slices.Contains(actions, action) {
		fmt.Println("action should be 'run' or 'start'")
		return
	}

	project, ok := projects[*name]
	if !ok {
		fmt.Printf("project \"%s\" does not exists\n", *name)
		return
	}

	cmd := exec.Command(
		"docker", "compose",
		"--file", project.ConfigPath,
		"--project-name", project.Name,
		action,
	)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
