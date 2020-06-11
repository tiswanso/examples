package vppagent

import (
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/sirupsen/logrus"
	"go.ligato.io/vpp-agent/v3/proto/ligato/vpp"
)

// UniversalCNFVPPAgentBackend is the VPP CNF backend struct
type SrCNFVPPAgentBackend struct {
	ucnfBackend UniversalCNFVPPAgentBackend
}

func (srcnf *SrCNFVPPAgentBackend) NewDPConfig() *vpp.ConfigData {
	return srcnf.ucnfBackend.NewDPConfig()
}

// NewUniversalCNFBackend initializes the VPP CNF backend
func (srcnf *SrCNFVPPAgentBackend) NewUniversalCNFBackend() error {
	return srcnf.ucnfBackend.NewUniversalCNFBackend()
}

func (srcnf *SrCNFVPPAgentBackend) ProcessClient(dpconfig interface{}, ifName string, conn *connection.Connection) error {
	return srcnf.ucnfBackend.ProcessClient(dpconfig, ifName, conn)
}

func (srcnf *SrCNFVPPAgentBackend) ProcessEndpoint(dpconfig interface{}, serviceName, ifName string, conn *connection.Connection) error {
	err := srcnf.ucnfBackend.ProcessEndpoint(dpconfig, serviceName, ifName, conn)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"serviceName": serviceName,
		"ifName":      ifName,
	}).Infof("SrCNFVPPAgentBackend ProcessEndpoint--doing SR stuff")

	return nil
}

func (srcnf *SrCNFVPPAgentBackend) GetEndpointIfID(serviceName string) string {
	return srcnf.ucnfBackend.GetEndpointIfID(serviceName)
}

func (srcnf *SrCNFVPPAgentBackend) ProcessDPConfig(dpconfig interface{}) error {
	return srcnf.ucnfBackend.ProcessDPConfig(dpconfig)
}
