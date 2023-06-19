package script_service

type Service struct {
	startScripts map[string]string
	scripts      map[string]any

	resourcesPath string
	scriptsPath   string
	startPath     string
}

func New() *Service {
	return &Service{
		resourcesPath: defaultResourcesPath,
		scriptsPath:   defaultScriptsPath,
		startPath:     defaultStartPath,
	}
}
