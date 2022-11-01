// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type APICertsSpec -type CertSANSpec -type EtcdCertsSpec -type EtcdRootSpec -type KubeletSpec -type KubernetesCertsSpec -type KubernetesRootSpec -type OSRootSpec -type TrustdCertsSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package secrets

import (
	"net/netip"
	"net/url"
)

// DeepCopy generates a deep copy of APICertsSpec.
func (o APICertsSpec) DeepCopy() APICertsSpec {
	var cp APICertsSpec = o
	if o.CA != nil {
		cp.CA = o.CA.DeepCopy()
	}
	if o.Client != nil {
		cp.Client = o.Client.DeepCopy()
	}
	if o.Server != nil {
		cp.Server = o.Server.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of CertSANSpec.
func (o CertSANSpec) DeepCopy() CertSANSpec {
	var cp CertSANSpec = o
	if o.IPs != nil {
		cp.IPs = make([]netip.Addr, len(o.IPs))
		copy(cp.IPs, o.IPs)
	}
	if o.DNSNames != nil {
		cp.DNSNames = make([]string, len(o.DNSNames))
		copy(cp.DNSNames, o.DNSNames)
	}
	return cp
}

// DeepCopy generates a deep copy of EtcdCertsSpec.
func (o EtcdCertsSpec) DeepCopy() EtcdCertsSpec {
	var cp EtcdCertsSpec = o
	if o.Etcd != nil {
		cp.Etcd = o.Etcd.DeepCopy()
	}
	if o.EtcdPeer != nil {
		cp.EtcdPeer = o.EtcdPeer.DeepCopy()
	}
	if o.EtcdAdmin != nil {
		cp.EtcdAdmin = o.EtcdAdmin.DeepCopy()
	}
	if o.EtcdAPIServer != nil {
		cp.EtcdAPIServer = o.EtcdAPIServer.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of EtcdRootSpec.
func (o EtcdRootSpec) DeepCopy() EtcdRootSpec {
	var cp EtcdRootSpec = o
	if o.EtcdCA != nil {
		cp.EtcdCA = o.EtcdCA.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of KubeletSpec.
func (o KubeletSpec) DeepCopy() KubeletSpec {
	var cp KubeletSpec = o
	if o.Endpoint != nil {
		cp.Endpoint = new(url.URL)
		*cp.Endpoint = *o.Endpoint
		if o.Endpoint.User != nil {
			cp.Endpoint.User = new(url.Userinfo)
			*cp.Endpoint.User = *o.Endpoint.User
		}
	}
	if o.CA != nil {
		cp.CA = o.CA.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of KubernetesCertsSpec.
func (o KubernetesCertsSpec) DeepCopy() KubernetesCertsSpec {
	var cp KubernetesCertsSpec = o
	if o.APIServer != nil {
		cp.APIServer = o.APIServer.DeepCopy()
	}
	if o.APIServerKubeletClient != nil {
		cp.APIServerKubeletClient = o.APIServerKubeletClient.DeepCopy()
	}
	if o.FrontProxy != nil {
		cp.FrontProxy = o.FrontProxy.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of KubernetesRootSpec.
func (o KubernetesRootSpec) DeepCopy() KubernetesRootSpec {
	var cp KubernetesRootSpec = o
	if o.Endpoint != nil {
		cp.Endpoint = new(url.URL)
		*cp.Endpoint = *o.Endpoint
		if o.Endpoint.User != nil {
			cp.Endpoint.User = new(url.Userinfo)
			*cp.Endpoint.User = *o.Endpoint.User
		}
	}
	if o.LocalEndpoint != nil {
		cp.LocalEndpoint = new(url.URL)
		*cp.LocalEndpoint = *o.LocalEndpoint
		if o.LocalEndpoint.User != nil {
			cp.LocalEndpoint.User = new(url.Userinfo)
			*cp.LocalEndpoint.User = *o.LocalEndpoint.User
		}
	}
	if o.CertSANs != nil {
		cp.CertSANs = make([]string, len(o.CertSANs))
		copy(cp.CertSANs, o.CertSANs)
	}
	if o.APIServerIPs != nil {
		cp.APIServerIPs = make([]netip.Addr, len(o.APIServerIPs))
		copy(cp.APIServerIPs, o.APIServerIPs)
	}
	if o.CA != nil {
		cp.CA = o.CA.DeepCopy()
	}
	if o.ServiceAccount != nil {
		cp.ServiceAccount = o.ServiceAccount.DeepCopy()
	}
	if o.AggregatorCA != nil {
		cp.AggregatorCA = o.AggregatorCA.DeepCopy()
	}
	return cp
}

// DeepCopy generates a deep copy of OSRootSpec.
func (o OSRootSpec) DeepCopy() OSRootSpec {
	var cp OSRootSpec = o
	if o.CA != nil {
		cp.CA = o.CA.DeepCopy()
	}
	if o.CertSANIPs != nil {
		cp.CertSANIPs = make([]netip.Addr, len(o.CertSANIPs))
		copy(cp.CertSANIPs, o.CertSANIPs)
	}
	if o.CertSANDNSNames != nil {
		cp.CertSANDNSNames = make([]string, len(o.CertSANDNSNames))
		copy(cp.CertSANDNSNames, o.CertSANDNSNames)
	}
	return cp
}

// DeepCopy generates a deep copy of TrustdCertsSpec.
func (o TrustdCertsSpec) DeepCopy() TrustdCertsSpec {
	var cp TrustdCertsSpec = o
	if o.CA != nil {
		cp.CA = o.CA.DeepCopy()
	}
	if o.Server != nil {
		cp.Server = o.Server.DeepCopy()
	}
	return cp
}
