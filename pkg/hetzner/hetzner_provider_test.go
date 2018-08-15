package hetzner_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/xetys/hetzner-kube/pkg/clustermanager"
	"github.com/xetys/hetzner-kube/pkg/hetzner"
)

func getProviderWithNodes(nodes []clustermanager.Node) hetzner.Provider {
	provider := hetzner.Provider{}

	provider.SetNodes(nodes)

	return provider
}

func TestProviderGetMasterNodes(t *testing.T) {
	tests := []struct {
		Name         string
		Nodes        []clustermanager.Node
		MatchedNodes []string
	}{
		{
			Name: "Single master node",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-master-1",
			},
		},
		{
			Name: "Two master nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-master-2", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-master-1",
				"kube-master-2",
			},
		},
		{
			Name: "Two etcd node that are also master",
			Nodes: []clustermanager.Node{
				{Name: "kube-etcd-1", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-2", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-3", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-etcd-1",
				"kube-etcd-2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			nodes := provider.GetMasterNodes()

			nodeNames := []string{}

			for _, node := range nodes {
				nodeNames = append(nodeNames, node.Name)
			}

			assert.Equal(t, nodeNames, tt.MatchedNodes)
		})
	}
}

func TestProviderGetEtcdNodes(t *testing.T) {
	tests := []struct {
		Name         string
		Nodes        []clustermanager.Node
		MatchedNodes []string
	}{
		{
			Name: "Single etcd node",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-etcd-1",
			},
		},
		{
			Name: "Two etcd nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-etcd-2", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-etcd-1",
				"kube-etcd-2",
			},
		},
		{
			Name: "Three etcd node some of them are also master",
			Nodes: []clustermanager.Node{
				{Name: "kube-etcd-1", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-2", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-3", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-etcd-1",
				"kube-etcd-2",
				"kube-etcd-3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			nodes := provider.GetEtcdNodes()

			nodeNames := []string{}

			for _, node := range nodes {
				nodeNames = append(nodeNames, node.Name)
			}

			assert.Equal(t, nodeNames, tt.MatchedNodes)
		})
	}
}

func TestProviderGetWorkerNodes(t *testing.T) {
	tests := []struct {
		Name         string
		Nodes        []clustermanager.Node
		MatchedNodes []string
	}{
		{
			Name: "Single worker node",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-worker-1",
			},
		},
		{
			Name: "Two worker nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
				{Name: "kube-worker-2"},
			},
			MatchedNodes: []string{
				"kube-worker-1",
				"kube-worker-2",
			},
		},
		{
			Name: "No worker nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-etcd-1", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-2", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-3", IsEtcd: true},
			},
			MatchedNodes: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			nodes := provider.GetWorkerNodes()

			nodeNames := []string{}

			for _, node := range nodes {
				nodeNames = append(nodeNames, node.Name)
			}

			assert.Equal(t, nodeNames, tt.MatchedNodes)
		})
	}
}

func TestProviderGetAllNodes(t *testing.T) {
	tests := []struct {
		Name         string
		Nodes        []clustermanager.Node
		MatchedNodes []string
	}{
		{
			Name: "One node per type",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNodes: []string{
				"kube-master-1",
				"kube-etcd-1",
				"kube-worker-1",
			},
		},
		{
			Name: "Multiple node per type",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-master-2", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-etcd-2", IsEtcd: true},
				{Name: "kube-worker-1"},
				{Name: "kube-worker-2"},
				{Name: "kube-worker-3"},
			},
			MatchedNodes: []string{
				"kube-master-1",
				"kube-master-2",
				"kube-etcd-1",
				"kube-etcd-2",
				"kube-worker-1",
				"kube-worker-2",
				"kube-worker-3",
			},
		},
		{
			Name:         "No nodes",
			Nodes:        []clustermanager.Node{},
			MatchedNodes: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			nodes := provider.GetAllNodes()

			nodeNames := []string{}

			for _, node := range nodes {
				nodeNames = append(nodeNames, node.Name)
			}

			assert.Equal(t, nodeNames, tt.MatchedNodes)
		})
	}
}

func TestProviderGetMasterNode(t *testing.T) {
	tests := []struct {
		Name        string
		Nodes       []clustermanager.Node
		MatchedNode string
	}{
		{
			Name: "Single master node",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNode: "kube-master-1",
		},
		{
			Name: "Two master nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-master-2", IsMaster: true},
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
			MatchedNode: "kube-master-1",
		},
		{
			Name: "An etcd node that is also master",
			Nodes: []clustermanager.Node{
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-etcd-2", IsMaster: true, IsEtcd: true},
				{Name: "kube-etcd-3", IsEtcd: true},
				{Name: "kube-master-1", IsMaster: true},
				{Name: "kube-worker-1"},
			},
			MatchedNode: "kube-etcd-2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			node, _ := provider.GetMasterNode()

			assert.Equal(t, node.Name, tt.MatchedNode)
		})
	}
}

func TestProviderGetMasterNodeIsMissing(t *testing.T) {
	tests := []struct {
		Name        string
		Nodes       []clustermanager.Node
		MatchedNode string
	}{
		{
			Name:  "No nodes",
			Nodes: []clustermanager.Node{},
		},
		{
			Name: "No master nodes",
			Nodes: []clustermanager.Node{
				{Name: "kube-etcd-1", IsEtcd: true},
				{Name: "kube-worker-1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			provider := getProviderWithNodes(tt.Nodes)
			_, err := provider.GetMasterNode()

			if err == nil {
				t.Error("no error ommited with no master")
			}
		})
	}
}
