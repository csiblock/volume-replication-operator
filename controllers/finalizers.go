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
	"fmt"

	volumegroupv1 "github.com/IBM/csi-volume-group-operator/api/v1"

	replicationv1alpha1 "github.com/csi-addons/volume-replication-operator/api/v1alpha1"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
)

const (
	volumeReplicationFinalizer = "replication.storage.openshift.io"
	pvcReplicationFinalizer    = "replication.storage.openshift.io/pvc-protection"
	vgReplicationFinalizer     = "replication.storage.openshift.io/vg-protection"
)

// addFinalizerToVR adds the VR finalizer on the VolumeReplication instance.
func (r *VolumeReplicationReconciler) addFinalizerToVR(ctx context.Context, logger logr.Logger, vr *replicationv1alpha1.VolumeReplication,
) error {
	if !contains(vr.ObjectMeta.Finalizers, volumeReplicationFinalizer) {
		logger.Info("adding finalizer to volumeReplication object", "Finalizer", volumeReplicationFinalizer)
		vr.ObjectMeta.Finalizers = append(vr.ObjectMeta.Finalizers, volumeReplicationFinalizer)
		if err := r.Client.Update(ctx, vr); err != nil {
			return fmt.Errorf("failed to add finalizer (%s) to VolumeReplication resource"+
				" (%s/%s) %w",
				volumeReplicationFinalizer, vr.Namespace, vr.Name, err)
		}
	}

	return nil
}

// removeFinalizerFromVR removes the VR finalizer from the VolumeReplication instance.
func (r *VolumeReplicationReconciler) removeFinalizerFromVR(ctx context.Context, logger logr.Logger, vr *replicationv1alpha1.VolumeReplication) error {
	if contains(vr.ObjectMeta.Finalizers, volumeReplicationFinalizer) {
		logger.Info("removing finalizer from volumeReplication object", "Finalizer", volumeReplicationFinalizer)
		vr.ObjectMeta.Finalizers = remove(vr.ObjectMeta.Finalizers, volumeReplicationFinalizer)
		if err := r.Client.Update(ctx, vr); err != nil {
			return fmt.Errorf("failed to remove finalizer (%s) from VolumeReplication resource"+
				" (%s/%s), %w",
				volumeReplicationFinalizer, vr.Namespace, vr.Name, err)
		}
	}

	return nil
}

// addFinalizerToPVC adds the VR finalizer on the PersistentVolumeClaim.
func (r *VolumeReplicationReconciler) addFinalizerToPVC(ctx context.Context, logger logr.Logger, pvc *corev1.PersistentVolumeClaim) error {
	if !contains(pvc.ObjectMeta.Finalizers, pvcReplicationFinalizer) {
		logger.Info("adding finalizer to PersistentVolumeClaim object", "Finalizer", pvcReplicationFinalizer)
		pvc.ObjectMeta.Finalizers = append(pvc.ObjectMeta.Finalizers, pvcReplicationFinalizer)
		if err := r.Client.Update(ctx, pvc); err != nil {
			return fmt.Errorf("failed to add finalizer (%s) to PersistentVolumeClaim resource"+
				" (%s/%s) %w",
				pvcReplicationFinalizer, pvc.Namespace, pvc.Name, err)
		}
	}

	return nil
}

// removeFinalizerFromPVC removes the VR finalizer on PersistentVolumeClaim.
func (r *VolumeReplicationReconciler) removeFinalizerFromPVC(ctx context.Context, logger logr.Logger, pvc *corev1.PersistentVolumeClaim,
) error {
	if contains(pvc.ObjectMeta.Finalizers, pvcReplicationFinalizer) {
		logger.Info("removing finalizer from PersistentVolumeClaim object", "Finalizer", pvcReplicationFinalizer)
		pvc.ObjectMeta.Finalizers = remove(pvc.ObjectMeta.Finalizers, pvcReplicationFinalizer)
		if err := r.Client.Update(ctx, pvc); err != nil {
			return fmt.Errorf("failed to remove finalizer (%s) from PersistentVolumeClaim resource"+
				" (%s/%s), %w",
				pvcReplicationFinalizer, pvc.Namespace, pvc.Name, err)
		}
	}

	return nil
}

// addFinalizerToVG adds the VR finalizer on the VolumeGroup.
func (r *VolumeReplicationReconciler) addFinalizerToVG(ctx context.Context, logger logr.Logger, vg *volumegroupv1.VolumeGroup) error {
	if !contains(vg.ObjectMeta.Finalizers, vgReplicationFinalizer) {
		logger.Info("adding finalizer to VolumeGroup object", "Finalizer", vgReplicationFinalizer)
		vg.ObjectMeta.Finalizers = append(vg.ObjectMeta.Finalizers, vgReplicationFinalizer)
		if err := r.Client.Update(ctx, vg); err != nil {
			return fmt.Errorf("failed to add finalizer (%s) to VolumeGroup resource"+
				" (%s/%s) %w",
				vgReplicationFinalizer, vg.Namespace, vg.Name, err)
		}
	}

	return nil
}

// removeFinalizerFromVG removes the VR finalizer on VolumeGroup.
func (r *VolumeReplicationReconciler) removeFinalizerFromVG(ctx context.Context, logger logr.Logger, vg *volumegroupv1.VolumeGroup,
) error {
	if contains(vg.ObjectMeta.Finalizers, vgReplicationFinalizer) {
		logger.Info("removing finalizer from VolumeGroup object", "Finalizer", vgReplicationFinalizer)
		vg.ObjectMeta.Finalizers = remove(vg.ObjectMeta.Finalizers, vgReplicationFinalizer)
		if err := r.Client.Update(ctx, vg); err != nil {
			return fmt.Errorf("failed to remove finalizer (%s) from VolumeGroup resource"+
				" (%s/%s), %w",
				vgReplicationFinalizer, vg.Namespace, vg.Name, err)
		}
	}

	return nil
}
