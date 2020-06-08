package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
	"github.com/tiswanso/examples/examples/universal-cnf/vppagent/pkg/config"
	"sync"

	"github.com/networkservicemesh/networkservicemesh/sdk/common"
	"github.com/networkservicemesh/networkservicemesh/sdk/endpoint"
	"github.com/sirupsen/logrus"
)

type srAppsConnectComposite struct {
	sync.RWMutex
	myEndpointName    string
	nsConfig          *common.NSConfiguration
	ipamEndpoint  *endpoint.IpamEndpoint
	backend       config.UniversalCNFBackend
	myNseNameFunc fnGetNseName
}

func (vxc *srAppsConnectComposite) SetMyNseName(request *networkservice.NetworkServiceRequest) {
	vxc.Lock()
	defer vxc.Unlock()
	if vxc.myEndpointName == "" {
		nseName := vxc.myNseNameFunc()
		logrus.Infof("Setting SR-Apps connect composite endpoint name to \"%s\"--req contains \"%s\"", nseName, request.GetConnection().GetNetworkServiceEndpointName())
		if request.GetConnection().GetNetworkServiceEndpointName() != "" {
			vxc.myEndpointName = request.GetConnection().GetNetworkServiceEndpointName()
		} else {
			vxc.myEndpointName = nseName
		}
	}
}

func (vxc *srAppsConnectComposite) GetMyNseName() string {
	vxc.Lock()
	defer vxc.Unlock()
	return vxc.myEndpointName
}

func (vxc *srAppsConnectComposite) Request(ctx context.Context,
	request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	logger := logrus.New() // endpoint.Log(ctx)
	logger.WithFields(logrus.Fields{
		"endpointName":              request.GetConnection().GetNetworkServiceEndpointName(),
		"networkServiceManagerName": request.GetConnection().GetSourceNetworkServiceManagerName(),
	}).Infof("srAppsConnectComposite Request handler")

	logger.Infof("NSC client IP %s", request.GetConnection().GetContext().GetIpContext().SrcIpAddr)
	logger.Infof("srAppsConnectComposite request done")
	//return incoming, nil
	if endpoint.Next(ctx) != nil {
		return endpoint.Next(ctx).Request(ctx, request)
	}
	return request.GetConnection(), nil
}

func (vxc *srAppsConnectComposite) Close(ctx context.Context, conn *connection.Connection) (*empty.Empty, error) {
	// remove from connections
	// TODO: should we be removing all peer connections here or no?
	if endpoint.Next(ctx) != nil {
		return endpoint.Next(ctx).Close(ctx, conn)
	}
	return &empty.Empty{}, nil
}

// Name returns the composite name
func (vxc *srAppsConnectComposite) Name() string {
	return "vL3 NSE"
}

// NewVppAgentComposite creates a new VPP Agent composite
func newSrAppsConnectComposite(configuration *common.NSConfiguration, ipamCidr string, backend config.UniversalCNFBackend, remoteIpList []string, getNseName fnGetNseName) *srAppsConnectComposite {

	newSrAppsConnectComposite := &srAppsConnectComposite{
		nsConfig:          configuration,
		myEndpointName:    "",
		backend:           backend,
		myNseNameFunc:     getNseName,
	}

	logrus.Infof("newSrAppsConnectComposite returning")

	return newSrAppsConnectComposite
}
