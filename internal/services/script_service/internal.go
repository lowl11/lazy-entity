package script_service

import "github.com/lowl11/lazyfile/folderapi"

func (event *Service) readStartScripts() error {
	if !folderapi.Exist("resources/scripts/start") {
		return nil
	}

	files, err := folderapi.Objects("resources/scripts/start")
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

		event.startScripts[file.Name] = string(body)
	}

	return nil
}

func (event *Service) readScripts() error {
	if !folderapi.Exist("resources/scripts") {
		return nil
	}

	folders, err := folderapi.Objects("resources/scripts/")
	if err != nil {
		return err
	}

	for _, folder := range folders {
		if !folder.IsFolder {
			continue
		}

		folderMap := make(map[string]string)

		files, err := folderapi.Objects("resources/scripts/" + folder.Name)
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

		event.scripts[folder.Name] = folderMap
	}

	return nil
}
