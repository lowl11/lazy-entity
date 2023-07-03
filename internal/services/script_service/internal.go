package script_service

import (
	"errors"
	managerErrors "github.com/lowl11/lazyfile/data/errors"
)

const (
	defaultResourcesPath = "resources/scripts"
)

func (service *Service) readStartScripts() error {
	startFolder, err := service.manager.Folder("start")
	if err != nil {
		if errors.Is(err, managerErrors.FolderNotExist) {
			return nil
		}

		return err
	}

	files, err := startFolder.FileList()
	if err != nil {
		return err
	}

	for _, file := range files {
		service.startScripts[file.Name()] = file.String()
	}

	return nil
}

func (service *Service) readScripts() error {
	if err := service.manager.Sync(); err != nil {
		if errors.Is(err, managerErrors.FolderNotExist); err != nil {
			return nil
		}

		return err
	}

	folders, err := service.manager.FolderList()
	if err != nil {
		return err
	}

	for _, folder := range folders {
		if folder.Name() == "start" {
			continue
		}

		folderMap := make(map[string]string)

		files, err := folder.FileList()
		if err != nil {
			return err
		}

		for _, script := range files {
			folderMap[script.Name()] = script.String()
		}

		service.scripts[folder.Name()] = folderMap
	}

	return nil
}
