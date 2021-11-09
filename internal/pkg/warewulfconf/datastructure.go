package warewulfconf

import (
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/creasty/defaults"
)

const ConfigFile = "/etc/warewulf/warewulf.conf"

type ControllerConf struct {
	Comment  string        `yaml:"comment"`
	Ipaddr   string        `yaml:"ipaddr"`
	Netmask  string        `yaml:"netmask"`
	Network  string        `yaml:"network,omitempty"`
	Fqdn     string        `yaml:"fqdn,omitempty"`
	Warewulf *WarewulfConf `yaml:"warewulf"`
	Dhcp     *DhcpConf     `yaml:"dhcp"`
	Tftp     *TftpConf     `yaml:"tftp"`
	Nfs      *NfsConf      `yaml:"nfs"`
}

type WarewulfConf struct {
	Port              int  `yaml:"port"`
	Secure            bool `yaml:"secure"`
	UpdateInterval    int  `yaml:"update interval"`
	AutobuildOverlays bool `yaml:"autobuild overlays"`
	Syslog            bool `yaml:"syslog"`
}

type DhcpConf struct {
	Enabled     bool   `yaml:"enabled"`
	Template    string `yaml:"template"`
	RangeStart  string `yaml:"range start"`
	RangeEnd    string `yaml:"range end"`
	SystemdName string `yaml:"systemd name"`
	ConfigFile  string `yaml:"config file"`
}

type TftpConf struct {
	Enabled     bool   `yaml:"enabled"`
	TftpRoot    string `yaml:"tftproot"`
	SystemdName string `yaml:"systemd name"`
}

type NfsConf struct {
	Enabled     bool     `default:"true" yaml:"enabled"`
	Exports     []string `yaml:"exports"`
	ExportsExtended []*NfsExportConf `yaml:"exports extended"`
	SystemdName string   `yaml:"systemd name"`
}

func (s *NfsConf) UnmarshalYAML(unmarshal func(interface{}) error) error {
    if err := defaults.Set(s); err != nil {
		return err
	}

    type plain NfsConf
    if err := unmarshal((*plain)(s)); err != nil {
        return err
    }

    return nil
}

type NfsExportConf struct {
	Path     string  `yaml:"path"`
	Options  string  `yaml:"options"`
	Mount    bool    `yaml:"mount"`
}

func init() {
	//TODO: Check to make sure nodes.conf is found
	if !util.IsFile(ConfigFile) {
		wwlog.Printf(wwlog.ERROR, "Configuration file not found: %s\n", ConfigFile)
		// fail silently as this also called by bash_completion
		return
	}
}
