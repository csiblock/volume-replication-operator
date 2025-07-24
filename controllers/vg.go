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

package controllers

import (
	"context"

	volumegroupv1 "github.com/IBM/csi-volume-group-operator/api/v1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

// getVGDataSource get vg content, vg object from the request.
func (r VolumeReplicationReconciler) getVGDataSource(ctx context.Context, logger logr.Logger, req types.NamespacedName) (*volumegroupv1.VolumeGroup, *volumegroupv1.VolumeGroupContent, error) {
	vg := &volumegroupv1.VolumeGroup{}
	err := r.Client.Get(ctx, req, vg)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Error(err, "VG not found", "VG Name", req.Name)
		}

		return nil, nil, err
	}

	vgc := &volumegroupv1.VolumeGroupContent{}
	volumeGroupContentName := *vg.Spec.Source.VolumeGroupContentName
	namespacedVGC := types.NamespacedName{Name: volumeGroupContentName, Namespace: vg.Namespace}
	err = r.Client.Get(ctx, namespacedVGC, vgc)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Error(err, "VolumeGroupContent not found", "VolumeGroupContent Name", volumeGroupContentName)
		}
		return vg, nil, err
	}

	return vg, vgc, nil
}
