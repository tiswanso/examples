package serviceregistry

import (
	"bytes"
	"fmt"
	"net"
)

type validationErrors []error

func (e validationErrors) Error() string {
	b := bytes.NewBufferString("")
	for _, err := range e {
		_, _ = fmt.Fprintf(b, "\t%s", err)
	}
	return b.String()
}

func (x *ServiceRequest) Validate() error {
	var errs validationErrors
	if x.Name == "" {
		errs = append(errs, fmt.Errorf("service name is a mandatory parameter"))
	}
	if x.ConnectivityDomain == "" {
		errs = append(errs, fmt.Errorf("connectivity domain is a mandatory parameter"))
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (x *ServiceWorkload) Validate() error {
	var errs validationErrors
	if x.ServiceName == "" {
		errs = append(errs, fmt.Errorf("service name is a mandatory parameter"))
	}
	if x.ConnectivityDomain == "" {
		errs = append(errs, fmt.Errorf("connectivity domain is a mandatory parameter"))
	}
	for _, p := range x.Ports {
		if p > 65535 || p < 1 {
			errs = append(errs, fmt.Errorf("port \"%d\" is invalid - allowed port range: 1-65535", p))
		}
	}
	var w *Workload
	for _, w = range x.Workloads {
		errs = append(errs, w.Validate())
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (x *Workload) Validate() error {
	var errs validationErrors
	errs = append(errs, x.Identifier.Validate())
	if len(x.IPAddress) > 1 {
		for _, s := range x.IPAddress {
			_, err := toIpNet(s)
			if err != nil {
				errs = append(errs, fmt.Errorf("invalid ip %s : %v", s, err))
			}
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (x *WorkloadIdentifier) Validate() error {
	var errs validationErrors
	if x.Cluster == "" {
		errs = append(errs, fmt.Errorf("cluster name is a mandatory parameter"))
	}
	if x.PodName == "" {
		errs = append(errs, fmt.Errorf("pod name is a mandatory parameter"))
	}
	if x.Name == "" {
		errs = append(errs, fmt.Errorf("workload name is a mandatory parameter"))
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func toIpNet(ipStr string) (net.IP, error) {
	ip, _, err := net.ParseCIDR(ipStr)
	return ip, err
}
