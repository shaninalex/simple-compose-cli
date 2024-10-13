package main

import (
	"context"
	"fmt"

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
)

func main() {

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

	for k, v := range projects {
		fmt.Printf("-- %s ( %s )\n", k, v.ConfigPath)

		for _, c := range v.Containers {
			fmt.Println("-", c.Name)
		}
	}

	// docker compose --file '/job-assignment-app/docker-compose.dev.yml' --project-name 'jobs' stop
	// docker compose --file '/fastapi_sqlalchemy/docker-compose.yml' --project-name 'fastapi_sqlalchemy' start
}
