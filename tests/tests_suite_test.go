// Copyright 2018 Red Hat, Inc.
// Copyright 2014 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests_test

import (
	"flag"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ginkgo_reporters "kubevirt.io/qe-tools/pkg/ginkgo-reporters"

	clusterapi "github.com/k8snetworkplumbingwg/ovs-cni/tests/cluster"
)

var kubeconfig *string
var clusterApi *clusterapi.ClusterAPI

func TestPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	reporters := make([]Reporter, 0)
	if ginkgo_reporters.JunitOutput != "" {
		reporters = append(reporters, ginkgo_reporters.NewJunitReporter())
	}
	RunSpecsWithDefaultAndCustomReporters(t, "Plugin Suite", reporters)
}

var _ = BeforeSuite(func() {
	flag.Parse()

	// Change to root directory some test expect that
	os.Chdir("../")

	clusterApi = clusterapi.NewClusterAPI(*kubeconfig)
	clusterApi.RemoveTestNamespace()
	clusterApi.CreateTestNamespace()
})

var _ = AfterSuite(func() {
	clusterApi.RemoveTestNamespace()
})

func init() {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
}
