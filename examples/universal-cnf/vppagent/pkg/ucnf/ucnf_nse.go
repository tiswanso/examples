package ucnf

import (
	"context"
	"os"

	"github.com/danielvladco/k8s-vnet/pkg/nseconfig"
	"github.com/davecgh/go-spew/spew"
	"github.com/networkservicemesh/networkservicemesh/sdk/common"
	"github.com/sirupsen/logrus"
	"github.com/tiswanso/examples/api/ipam/ipprovider"
	"github.com/tiswanso/examples/examples/universal-cnf/vppagent/pkg/config"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

//
type UcnfNse struct {
	processEndpoints *config.ProcessEndpoints
}

func (ucnf *UcnfNse) Cleanup() {
	ucnf.processEndpoints.Cleanup()
}

func NewUcnfNse(configPath string, verify bool, backend config.UniversalCNFBackend, ceAddons config.CompositeEndpointAddons, ctx context.Context) *UcnfNse {
	cnfConfig := &nseconfig.Config{}
	f, err := os.Open(configPath)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		err = f.Close()
		logrus.Errorf("closing file failed %v", err)
	}()
	err = nseconfig.NewConfig(yaml.NewDecoder(f), cnfConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := backend.NewUniversalCNFBackend(); err != nil {
		logrus.Fatal(err)
	}

	if verify {
		spew.Dump(cnfConfig)
		return nil
	}

	configuration := common.FromEnv()
	newIpamService := func(addr string) config.IpamService {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			logrus.Fatal("unable to connect to ipam server: %v", err)
		}

		ipamAllocator := ipprovider.NewAllocatorClient(conn)
		ipamService := config.IpamServiceImpl{
			IpamAllocator:     ipamAllocator,
			RegisteredSubnets: make(chan *ipprovider.Subnet),
			Ctx:               ctx,
		}
		go func() {
			logrus.Info("begin the ipam leased subnet renew process")
			if err := ipamService.Renew(func(err error) {
				if err != nil {
					logrus.Error("unable to renew the subnet", err)
				}
			}); err != nil {
				logrus.Error(err)
			}
		}()
		return &ipamService
	}
	//add logic here

	pe := config.NewProcessEndpoints(backend, cnfConfig.Endpoints, configuration, ceAddons, newIpamService)

	ucnfnse := &UcnfNse{
		processEndpoints: pe,
	}

	logrus.Infof("Starting endpoints")
	// defer pe.Cleanup()

	if err := pe.Process(); err != nil {
		logrus.Fatalf("Error processing the new endpoints: %v", err)
	}
	return ucnfnse
}
