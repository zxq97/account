package accountbff

import "github.com/zxq97/gotool/config"

type AccountBffConfig struct {
	Svc *config.SvcConf `yaml:"svc"`

	IncludeETCD    string `yaml:"include_etcd"`
	IncludeHystrix string `yaml:"include_hystrix"`

	Etcd    map[string]*config.EtcdConf
	Hystrix map[string]*config.HystrixConf

	LogPath *config.LogConf `yaml:"log_path"`
}

func (conf *AccountBffConfig) Initialize() {
	if conf.IncludeETCD != "" {
		if err := config.LoadYaml(conf.IncludeETCD, &conf.Etcd); err != nil {
			panic(err)
		}
	}
	if conf.IncludeHystrix != "" {
		if err := config.LoadYaml(conf.IncludeHystrix, &conf.Hystrix); err != nil {
			panic(err)
		}
	}
	if err := conf.Svc.InitSvc(); err != nil {
		panic(err)
	}
}
