package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/open-cluster-management/observability-e2e-test/utils"
)

var _ = Describe("Observability:", func() {
	BeforeEach(func() {
		hubClient = utils.NewKubeClient(
			testOptions.HubCluster.MasterURL,
			testOptions.KubeConfig,
			testOptions.HubCluster.KubeContext)

		dynClient = utils.NewKubeClientDynamic(
			testOptions.HubCluster.MasterURL,
			testOptions.KubeConfig,
			testOptions.HubCluster.KubeContext)
	})

	It("should have metric data in grafana console (grafana/g0)", func() {
		Eventually(func() error {
			err, _ = utils.ContainManagedClusterMetric(testOptions, "node_memory_MemAvailable_bytes", "5m", []string{`"__name__":"node_memory_MemAvailable_bytes"`})
			return err
		}, EventuallyTimeoutMinute*5, EventuallyIntervalSecond*5).Should(Succeed())

		if err != nil {
			utils.PrintAllMCOPodsStatus(testOptions)
		}
	})
})
