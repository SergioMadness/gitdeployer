package config;

type Server struct {
	Name string;
	Path string;
	DefaultBranch string;
	GitUrl string;
	GitLogin string;
	GitPassword string;
}

func CreateServer(name, path, defaultBranch, gitUrl, gitLogin, gitPassword string) *Server {
	result := new(Server);
	
	result.Name = name;
	result.Path = path;
	result.DefaultBranch = defaultBranch;
	result.GitUrl = gitUrl;
	result.GitLogin = gitLogin;
	result.GitPassword = gitPassword;
	
	return result;
}

func RemoveServer(name string) bool {
	result := false;
	
	return result;
}

func IsServerExists(name string) bool {
	result := false;
	
	return result;
}