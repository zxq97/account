package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/zxq97/account/api"
	"github.com/zxq97/account/internal/service/accountbff"
	"github.com/zxq97/gotool/config"
	"github.com/zxq97/gotool/rpc"
	"google.golang.org/grpc/reflection"
)

var (
	confPath = flag.String("conf", "", "configuration file")
	conf     accountbff.AccountBffConfig
)

func main() {
	flag.Parse()
	err := config.LoadYaml(*confPath, &conf)
	if err != nil {
		panic(err)
	}
	conf.Initialize()
	etcdCli, err := conf.Etcd["etcd"].InitEtcd()
	if err != nil {
		panic(err)
	}
	conn, err := rpc.NewGrpcConn(etcdCli, conf.Svc.Name, conf.Hystrix["account"])
	if err != nil {
		panic(err)
	}
	accountBFF, err := accountbff.InitAccountBFF(&conf, conn)
	if err != nil {
		panic(err)
	}
	svc, er := rpc.NewGrpcServer(etcdCli, conf.Svc.Name+"_"+conf.Svc.Bind, conf.Svc.Bind)
	_, err = er.Register()
	if err != nil {
		panic(err)
	}
	api.RegisterAccountServer(svc, accountBFF)

	lis, err := net.Listen("tcp", conf.Svc.Bind)
	if err != nil {
		panic(err)
	}

	reflection.Register(svc)
	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		errCh <- svc.Serve(lis)
	}()
	go func() {
		errCh <- http.ListenAndServe(conf.Svc.HttpBind, nil)
	}()

	select {
	case err = <-errCh:
		er.Stop()
		log.Println("accountbff stop err", err)
	case sig := <-sigCh:
		er.Stop()
		log.Println("accountbff stop sign", sig)
	}
}
