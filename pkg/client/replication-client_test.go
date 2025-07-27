/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"errors"
	"testing"

	"github.com/csi-addons/volume-replication-operator/pkg/client/fake"

	replicationlib "github.com/csi-addons/spec/lib/go/replication"
	"github.com/stretchr/testify/require"
)

func TestEnableVolumeReplication(t *testing.T) {
	t.Parallel()

	mockedEnableReplication := &fake.ReplicationClient{
		EnableVolumeReplicationMock: func(_, _ string, _, _ map[string]string) (*replicationlib.EnableVolumeReplicationResponse, error) {
			return &replicationlib.EnableVolumeReplicationResponse{}, nil
		},
	}
	client := mockedEnableReplication

	resp, err := client.EnableVolumeReplication("", "", nil, nil)
	require.Equal(t, &replicationlib.EnableVolumeReplicationResponse{}, resp)
	require.NoError(t, err)

	// return error
	mockedEnableReplication = &fake.ReplicationClient{
		EnableVolumeReplicationMock: func(_, _ string, _, _ map[string]string) (*replicationlib.EnableVolumeReplicationResponse, error) {
			return nil, errors.New("failed to enable mirroring")
		},
	}

	client = mockedEnableReplication

	resp, err = client.EnableVolumeReplication("", "", nil, nil)
	require.Nil(t, resp)
	require.Error(t, err)
}

func TestDisableVolumeReplication(t *testing.T) {
	t.Parallel()

	mockedDisableReplication := &fake.ReplicationClient{
		DisableVolumeReplicationMock: func(_, _ string, _, _ map[string]string) (*replicationlib.DisableVolumeReplicationResponse, error) {
			return &replicationlib.DisableVolumeReplicationResponse{}, nil
		},
	}
	client := mockedDisableReplication

	resp, err := client.DisableVolumeReplication("", "", nil, nil)
	require.Equal(t, &replicationlib.DisableVolumeReplicationResponse{}, resp)
	require.NoError(t, err)

	// return error
	mockedDisableReplication = &fake.ReplicationClient{
		DisableVolumeReplicationMock: func(_, _ string, _, _ map[string]string) (*replicationlib.DisableVolumeReplicationResponse, error) {
			return nil, errors.New("failed to disable mirroring")
		},
	}

	client = mockedDisableReplication

	resp, err = client.DisableVolumeReplication("", "", nil, nil)
	require.Nil(t, resp)
	require.Error(t, err)
}

func TestPromoteVolume(t *testing.T) {
	t.Parallel()
	// return success response
	mockedPromoteVolume := &fake.ReplicationClient{
		PromoteVolumeMock: func(_, _ string, _ bool, _, _ map[string]string) (*replicationlib.PromoteVolumeResponse, error) {
			return &replicationlib.PromoteVolumeResponse{}, nil
		},
	}
	client := mockedPromoteVolume

	resp, err := client.PromoteVolume("", "", false, nil, nil)
	require.Equal(t, &replicationlib.PromoteVolumeResponse{}, resp)
	require.NoError(t, err)

	// return error
	mockedPromoteVolume = &fake.ReplicationClient{
		PromoteVolumeMock: func(_, _ string, _ bool, _, _ map[string]string) (*replicationlib.PromoteVolumeResponse, error) {
			return nil, errors.New("failed to promote volume")
		},
	}

	client = mockedPromoteVolume

	resp, err = client.PromoteVolume("", "", false, nil, nil)
	require.Nil(t, resp)
	require.Error(t, err)
}

func TestDemoteVolume(t *testing.T) {
	t.Parallel()
	// return success response
	mockedDemoteVolume := &fake.ReplicationClient{
		DemoteVolumeMock: func(_, _ string, _, _ map[string]string) (*replicationlib.DemoteVolumeResponse, error) {
			return &replicationlib.DemoteVolumeResponse{}, nil
		},
	}
	client := mockedDemoteVolume

	resp, err := client.DemoteVolume("", "", nil, nil)
	require.Equal(t, &replicationlib.DemoteVolumeResponse{}, resp)
	require.NoError(t, err)

	// return error
	mockedDemoteVolume = &fake.ReplicationClient{
		DemoteVolumeMock: func(_, _ string, _, _ map[string]string) (*replicationlib.DemoteVolumeResponse, error) {
			return nil, errors.New("failed to demote volume")
		},
	}

	client = mockedDemoteVolume

	resp, err = client.DemoteVolume("", "", nil, nil)
	require.Nil(t, resp)
	require.Error(t, err)
}

func TestResyncVolume(t *testing.T) {
	t.Parallel()
	// return success response
	mockedResyncVolume := &fake.ReplicationClient{
		ResyncVolumeMock: func(_, _ string, _, _ map[string]string) (*replicationlib.ResyncVolumeResponse, error) {
			return &replicationlib.ResyncVolumeResponse{}, nil
		},
	}
	client := mockedResyncVolume

	resp, err := client.ResyncVolume("", "", nil, nil)
	require.Equal(t, &replicationlib.ResyncVolumeResponse{}, resp)
	require.NoError(t, err)

	// return error
	mockedResyncVolume = &fake.ReplicationClient{
		ResyncVolumeMock: func(_, _ string, _, _ map[string]string) (*replicationlib.ResyncVolumeResponse, error) {
			return nil, errors.New("failed to resync volume")
		},
	}

	client = mockedResyncVolume

	resp, err = client.ResyncVolume("", "", nil, nil)
	require.Nil(t, resp)
	require.Error(t, err)
}
