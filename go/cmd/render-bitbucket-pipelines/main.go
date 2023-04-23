package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	monorepo := struct {
		DockerImages        []string
		HelmCharts          []string
		HelmReleases        []string
		KptInstances        []string
		KptPackages         []string
		KubernetesKustomize []string
		KubernetesRaw       []string
		TerraformModules    []string
		TerraformSubmodules []string
	}{
		DockerImages:        listDirs("./docker/images"),
		HelmCharts:          listDirs("./helm/charts"),
		HelmReleases:        listDirs("./helm/releases"),
		KptInstances:        listDirs("./kpt/instances"),
		KptPackages:         listDirs("./kpt/packages"),
		KubernetesKustomize: listDirs("./kubernetes/kustomize"),
		KubernetesRaw:       listDirs("./kubernetes/raw"),
		TerraformModules:    listDirs("./terraform/modules"),
		TerraformSubmodules: listDirs("./terraform/submodules"),
	}
	templateFilePath := "./bitbucket-pipelines.gotpl.yml"
	outputFilePath := "./bitbucket-pipelines.yml"

	templateFile, err := template.ParseFiles(templateFilePath)
	if err != nil {
		panic(fmt.Errorf("template parse error: %v", err))
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(fmt.Errorf("file create error: %v", err))
	}

	err = templateFile.Execute(outputFile, monorepo)
	if err != nil {
		panic(fmt.Errorf("template execute error: %v", err))
	}
}

func listDirs(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(fmt.Errorf("read dir error (%s): %v", path, err))
	}

	var dirs []string
	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs
}
