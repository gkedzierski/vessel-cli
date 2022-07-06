package config

import "fmt"

type EnvironmentConfig struct {
	Remote     RemoteConfig `yaml:remote`
	Forwarding []string     `yaml:forwarding`
}

type RemoteConfig struct {
	Hostname     string `yaml:"hostname"`
	User         string `yaml:"user"`
	IdentityFile string `yaml:"identityfile"`
	Port         uint   `yaml:"port"`
	RemotePath   string `yaml:"path"`
	Alias        string `yaml:alias,omitempty`
}

func (c *EnvironmentConfig) Valid() (bool, error) {
	if len(c.Forwarding) < 1 {
		return false, fmt.Errorf("no forwarding ports are defined")
	}

	if c.Remote.Hostname == "" {
		return false, fmt.Errorf("no remote hostname (to SSH in with) defined")
	}

	if c.Remote.User == "" {
		return false, fmt.Errorf("no remote user (to SSH in with) defined")
	}

	if c.Remote.IdentityFile == "" {
		return false, fmt.Errorf("no remote indentity file (SSH key) defined")
	}

	if c.Remote.RemotePath == "" {
		return false, fmt.Errorf("no remote file path to sync with defined")
	}

	if c.Remote.Port == 0 {
		return false, fmt.Errorf("no remote port (to SSH into) defined")
	}

	return true, nil
}
