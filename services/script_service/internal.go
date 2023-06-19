package script_service

import (
	"github.com/lowl11/lazyfile/folderapi"
)

const (
	defaultResourcesPath = "resources"
	defaultScriptsPath   = defaultResourcesPath + "/scripts"
	defaultStartPath     = defaultResourcesPath + "/scripts/start"
)

func (service *Service) readStartScripts() error {
	if !folderapi.Exist(service.startPath) {
		return nil
	}

	files, err := folderapi.Objects(service.startPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsFolder {
			continue
		}

		body, err := file.Read()
		if err != nil {
			return err
		}

		service.startScripts[file.Name] = string(body)
	}

	return nil
}

func (service *Service) readScripts() error {
	if !folderapi.Exist(service.scriptsPath) {
		return nil
	}

	folders, err := folderapi.Objects(service.scriptsPath)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		if !folder.IsFolder {
			continue
		}

		folderMap := make(map[string]string)

		files, err := folderapi.Objects(service.scriptsPath + "/" + folder.Name)
		if err != nil {
			return err
		}

		for _, file := range files {
			body, err := file.Read()
			if err != nil {
				return err
			}

			folderMap[file.Name] = string(body)
		}

		service.scripts[folder.Name] = folderMap
	}

	return nil
}
