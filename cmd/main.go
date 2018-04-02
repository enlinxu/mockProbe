package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"

	"github.com/songbinliu/mockProbe/pkg/action"
	"github.com/songbinliu/mockProbe/pkg/discovery"
	"github.com/songbinliu/mockProbe/pkg/registration"

	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/service"
)

var (
	targetConf   string
	opsMgrConf   string
)

func getFlags() {
	flag.StringVar(&opsMgrConf, "turboConf", "./conf/turbo.json", "configuration file of OpsMgr")
	flag.StringVar(&targetConf, "targetConf", "./conf/target.json", "configuration file of target")

	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func buildProbe(targetConf string, stop chan struct{}) (*probe.ProbeBuilder, error) {
	//1. load target configuration
	config, err := discovery.NewTargetConf(targetConf)
	if err != nil {
		return nil, fmt.Errorf("failed to load json conf:%v", err)
	}

	//2. generate various clients
	regClient := registration.NewRegistrationClient()
	discoveryClient := discovery.NewDiscoveryClient(config)
	actionHandler := action.NewActionHandler(stop)

	builder := probe.NewProbeBuilder(config.TargetType, config.ProbeCategory).
		RegisteredBy(regClient).
		DiscoversTarget(config.Identifier, discoveryClient).
		ExecutesActionsBy(actionHandler)

	return builder, nil
}

func createTapService() (*service.TAPService, error) {
	turboConfig, err := service.ParseTurboCommunicationConfig(opsMgrConf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OpsMgrConfig: %v", err)
	}

	stop := make(chan struct{})
	probeBuilder, err := buildProbe(targetConf, stop)
	if err != nil {
		return nil, fmt.Errorf("failed to create probe: %v", err)
	}

	tapService, err := service.NewTAPServiceBuilder().
		WithTurboCommunicator(turboConfig).
		WithTurboProbe(probeBuilder).
		Create()

	if err != nil {
		return nil, fmt.Errorf("error when creating TapService: %v", err.Error())
	}

	return tapService, nil
}

func main() {
	getFlags()

	tap, err := createTapService()
	if err != nil {
		glog.Errorf("failed to create tapServier: %v", err)
	}

	tap.ConnectToTurbo()
}
