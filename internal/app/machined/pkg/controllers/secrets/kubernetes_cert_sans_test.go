// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//nolint:dupl
package secrets_test

import (
	"fmt"
	"net"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/stretchr/testify/suite"
	"github.com/talos-systems/go-retry/retry"
	"inet.af/netaddr"

	"github.com/talos-systems/talos/internal/app/machined/pkg/controllers/ctest"
	secretsctrl "github.com/talos-systems/talos/internal/app/machined/pkg/controllers/secrets"
	"github.com/talos-systems/talos/pkg/machinery/resources/k8s"
	"github.com/talos-systems/talos/pkg/machinery/resources/network"
	"github.com/talos-systems/talos/pkg/machinery/resources/secrets"
)

type KubernetesCertSANsSuite struct {
	ctest.DefaultSuite
}

func TestKubernetesCertSANsSuite(t *testing.T) {
	suite.Run(t, &KubernetesCertSANsSuite{
		DefaultSuite: ctest.DefaultSuite{
			AfterSetup: func(suite *ctest.DefaultSuite) {
				suite.Require().NoError(suite.Runtime().RegisterController(&secretsctrl.KubernetesCertSANsController{}))
			},
		},
	})
}

func (suite *KubernetesCertSANsSuite) TestReconcile() {
	rootSecrets := secrets.NewKubernetesRoot(secrets.KubernetesRootID)

	var err error

	rootSecrets.TypedSpec().CertSANs = []string{"example.com"}
	rootSecrets.TypedSpec().APIServerIPs = []net.IP{net.ParseIP("10.4.3.2"), net.ParseIP("10.2.1.3")}
	rootSecrets.TypedSpec().DNSDomain = "cluster.remote"
	rootSecrets.TypedSpec().Endpoint, err = url.Parse("https://some.url:6443/")
	suite.Require().NoError(err)
	rootSecrets.TypedSpec().LocalEndpoint, err = url.Parse("https://localhost:6443/")
	suite.Require().NoError(err)

	suite.Require().NoError(suite.State().Create(suite.Ctx(), rootSecrets))

	hostnameStatus := network.NewHostnameStatus(network.NamespaceName, network.HostnameID)
	hostnameStatus.TypedSpec().Hostname = "foo"
	hostnameStatus.TypedSpec().Domainname = "example.com"
	suite.Require().NoError(suite.State().Create(suite.Ctx(), hostnameStatus))

	nodeAddresses := network.NewNodeAddress(
		network.NamespaceName,
		network.FilteredNodeAddressID(network.NodeAddressAccumulativeID, k8s.NodeAddressFilterNoK8s),
	)
	nodeAddresses.TypedSpec().Addresses = []netaddr.IPPrefix{
		netaddr.MustParseIPPrefix("10.2.1.3/24"),
		netaddr.MustParseIPPrefix("172.16.0.1/32"),
	}
	suite.Require().NoError(suite.State().Create(suite.Ctx(), nodeAddresses))

	suite.AssertWithin(10*time.Second, 100*time.Millisecond, func() error {
		certSANs, err := ctest.Get[*secrets.CertSAN](
			suite,
			resource.NewMetadata(
				secrets.NamespaceName,
				secrets.CertSANType,
				secrets.CertSANKubernetesID,
				resource.VersionUndefined,
			),
		)
		if err != nil {
			if state.IsNotFoundError(err) {
				return retry.ExpectedError(err)
			}

			return err
		}

		spec := certSANs.TypedSpec()

		suite.Assert().Equal(
			[]string{
				"example.com",
				"foo",
				"foo.example.com",
				"kubernetes",
				"kubernetes.default",
				"kubernetes.default.svc",
				"kubernetes.default.svc.cluster.remote",
				"localhost",
				"some.url",
			}, spec.DNSNames,
		)
		suite.Assert().Equal("[10.2.1.3 10.4.3.2 172.16.0.1]", fmt.Sprintf("%v", spec.IPs))

		return nil
	})

	ctest.UpdateWithConflicts(suite, rootSecrets, func(rootSecrets *secrets.KubernetesRoot) error {
		var err error
		rootSecrets.TypedSpec().Endpoint, err = url.Parse("https://some.other.url:6443/")

		return err
	})

	suite.AssertWithin(10*time.Second, 100*time.Millisecond, func() error {
		var certSANs resource.Resource
		certSANs, err := ctest.Get[*secrets.CertSAN](
			suite,
			resource.NewMetadata(
				secrets.NamespaceName,
				secrets.CertSANType,
				secrets.CertSANKubernetesID,
				resource.VersionUndefined,
			),
		)
		if err != nil {
			return err
		}

		spec := certSANs.(*secrets.CertSAN).TypedSpec()

		expectedDNSNames := []string{
			"example.com",
			"foo",
			"foo.example.com",
			"kubernetes",
			"kubernetes.default",
			"kubernetes.default.svc",
			"kubernetes.default.svc.cluster.remote",
			"localhost",
			"some.other.url",
		}

		if !reflect.DeepEqual(spec.DNSNames, expectedDNSNames) {
			return retry.ExpectedErrorf("expected %v, got %v", expectedDNSNames, spec.DNSNames)
		}

		return nil
	})
}
