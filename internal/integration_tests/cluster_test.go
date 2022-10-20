package integration_tests

import (
	. "github.com/neo4j/helm-charts/internal/helpers"
	"github.com/neo4j/helm-charts/internal/model"
	"github.com/neo4j/helm-charts/internal/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstallNeo4jClusterInGcloud(t *testing.T) {
	if model.Neo4jEdition != "enterprise" {
		t.Skip()
		return
	}
	t.Parallel()

	var closeables []Closeable
	addCloseable := func(closeableList ...Closeable) {
		for _, closeable := range closeableList {
			closeables = append([]Closeable{closeable}, closeables...)
		}
	}

	err := labelNodes(t)
	addCloseable(func() error {
		return removeLabelFromNodes(t)
	})
	if !assert.NoError(t, err) {
		return
	}

	clusterReleaseName := model.NewReleaseName("cluster-" + TestRunIdentifier)
	headlessService := clusterHeadLessService{model.NewHeadlessServiceReleaseName(clusterReleaseName), nil}
	//defaultHelmArgs := []string{"--set", "volumes.data.volume.setOwnerAndGroupWritableFilePermissions=true",
	//	"--set", "volumes.logs.volume.setOwnerAndGroupWritableFilePermissions=true"}
	defaultHelmArgs := []string{}
	defaultHelmArgs = append(defaultHelmArgs, model.DefaultNeo4jName...)
	defaultHelmArgs = append(defaultHelmArgs, model.DefaultClusterSize...)
	core1HelmArgs := append(defaultHelmArgs, model.ImagePullSecretArgs...)
	core1HelmArgs = append(core1HelmArgs, model.NodeSelectorArgs...)
	core2HelmArgs := append(defaultHelmArgs, model.PriorityClassNameArgs...)
	core1 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 1), core1HelmArgs}
	core2 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 2), core2HelmArgs}
	core3 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 3), defaultHelmArgs}
	cores := []clusterCore{core1, core2, core3}

	t.Cleanup(func() { cleanupTest(t, AsCloseable(closeables)) })

	t.Logf("Starting setup of '%s'", t.Name())

	closeable, err := prepareK8s(t, clusterReleaseName)
	addCloseable(closeable)
	if !assert.NoError(t, err) {
		return
	}
	cleanPriorityClass, err := createPriorityClass(t, clusterReleaseName)
	addCloseable(cleanPriorityClass)
	if !assert.NoError(t, err) {
		return
	}
	// Install one core synchronously, if all cores are installed simultaneously they run into conflicts all trying to create a -auth secret
	result := core1.Install(t)
	addCloseable(result.Closeable)
	if !assert.NoError(t, result.error) {
		return
	}

	componentsToParallelInstall := []helmComponent{core2, core3, headlessService}
	closeablesNew, err := performBackgroundInstall(t, componentsToParallelInstall, clusterReleaseName)
	if !assert.NoError(t, err) {
		return
	}
	addCloseable(closeablesNew...)

	for _, core := range cores {
		err = run(t, "kubectl", "--namespace", string(core.Name().Namespace()), "rollout", "status", "--watch", "--timeout=180s", "statefulset/"+core.Name().String())
		if !assert.NoError(t, err) {
			return
		}
	}
	closeablesNew, err = performBackgroundInstall(t, componentsToParallelInstall, clusterReleaseName)
	if !assert.NoError(t, err) {
		return
	}
	addCloseable(closeablesNew...)

	t.Logf("Succeeded with setup of '%s'", t.Name())

	//subTests, err := clusterTests(loadBalancer.Name())
	//if !assert.NoError(t, err) {
	//	return
	//}
	//subTests = append(subTests, nodeSelectorTests(core1.Name())...)
	//subTests = append(subTests, headLessServiceTests(headlessService.Name())...)
	//runSubTests(t, subTests)

	t.Logf("Succeeded running all tests in '%s'", t.Name())
}

func TestInstallNeo4jClusterWithApocConfigInGcloud(t *testing.T) {
	if model.Neo4jEdition != "enterprise" {
		t.Skip()
		return
	}

	//if we make this in parallel with the other cluster tests , it will fail
	// we need to wait for this cluster test to complete so that the other cluster test can complete
	//t.Parallel()

	var closeables []Closeable
	addCloseable := func(closeableList ...Closeable) {
		for _, closeable := range closeableList {
			closeables = append([]Closeable{closeable}, closeables...)
		}
	}

	clusterReleaseName := model.NewReleaseName("apoc-cluster-" + TestRunIdentifier)
	loadBalancer := clusterLoadBalancer{model.NewLoadBalancerReleaseName(clusterReleaseName), nil}

	apocCustomArgs := append(model.CustomApocImageArgs, resources.ApocClusterTestConfig.HelmArgs()...)
	core1 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 1), apocCustomArgs}
	core2 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 2), apocCustomArgs}
	core3 := clusterCore{model.NewCoreReleaseName(clusterReleaseName, 3), apocCustomArgs}
	cores := []clusterCore{core1, core2, core3}

	t.Cleanup(func() { cleanupTest(t, AsCloseable(closeables)) })

	t.Logf("Starting setup of '%s'", t.Name())

	closeable, err := prepareK8s(t, clusterReleaseName)
	addCloseable(closeable)
	if !assert.NoError(t, err) {
		return
	}

	// Install one core synchronously, if all cores are installed simultaneously they run into conflicts all trying to create a -auth secret
	result := core1.Install(t)
	addCloseable(result.Closeable)
	if !assert.NoError(t, result.error) {
		return
	}

	componentsToParallelInstall := []helmComponent{core2, core3, loadBalancer}
	closeablesNew, err := performBackgroundInstall(t, componentsToParallelInstall, clusterReleaseName)
	if !assert.NoError(t, err) {
		return
	}
	addCloseable(closeablesNew...)

	for _, core := range cores {
		err = run(t, "kubectl", "--namespace", string(core.Name().Namespace()), "rollout", "status", "--watch", "--timeout=180s", "statefulset/"+core.Name().String())
		if !assert.NoError(t, err) {
			return
		}
	}

	t.Logf("Succeeded with setup of '%s'", t.Name())

	subTests := apocConfigTests(loadBalancer.Name())
	if !assert.NoError(t, err) {
		return
	}
	runSubTests(t, subTests)

	t.Logf("Succeeded running all apoc config tests in '%s'", t.Name())
}
