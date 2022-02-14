/*
Copyright 2022.

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

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ppv1alpha1 "github.com/akakakakakaa/python-package-operator/api/v1alpha1"
	"github.com/akakakakakaa/python-package-operator/util"
	"k8s.io/apimachinery/pkg/api/errors"
)

// GlobalPythonPackageCollectionReconciler reconciles a GlobalPythonPackageCollection object
type GlobalPythonPackageCollectionReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackagecollections,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackagecollections/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackagecollections/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete

func (r *GlobalPythonPackageCollectionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("gppc", req.NamespacedName)

	gppc := &ppv1alpha1.GlobalPythonPackageCollection{}
	if err := r.Get(context.TODO(), req.NamespacedName, gppc); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("GlobalPythonPackageCollection resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get GlobalPythonPackageCollection ["+req.Name+"].")
		return ctrl.Result{}, err
	}

	//set patch helper
	patchHelper, err := patch.NewHelper(gppc, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	defer func() {
		if err := patchHelper.Patch(context.TODO(), gppc); err != nil {
			logger.Error(err, "Failed to patch GlobalPythonPackageCollection.")
		}
	}()

	found := &v1.PersistentVolumeClaim{}
	err = r.Get(ctx, types.NamespacedName{Name: r.getPersistentVolumeClaimNameForGPPC(gppc), Namespace: util.Namespace()}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Create persistentVolumeClaim.")
		persistentVolumeClaim := r.persistentVolumeClaimForGPPC(gppc)
		if err := r.Create(context.TODO(), persistentVolumeClaim); err != nil {
			logger.Error(err, "Failed to create persistentVolumeClaim.")
			gppc.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
			gppc.Status.Reason = "Failed to create persistentVolumeClaim."
			gppc.Status.Message = err.Error()
			return ctrl.Result{}, err
		}

		logger.Info("Create persistentVolumeClaim success.")
		gppc.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeCreated
		gppc.Status.Reason = "Create persistentVolumeClaim success."
	} else if err != nil {
		logger.Error(err, "Failed to get persistentVolumeClaim.")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *GlobalPythonPackageCollectionReconciler) persistentVolumeClaimForGPPC(gppc *ppv1alpha1.GlobalPythonPackageCollection) *v1.PersistentVolumeClaim {
	persistentVolumeClaim := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      r.getPersistentVolumeClaimNameForGPPC(gppc),
			Namespace: util.Namespace(),
		},
		Spec: gppc.Spec.PersistentVolumeClaimSpec,
	}

	ctrl.SetControllerReference(gppc, persistentVolumeClaim, r.Scheme)
	return persistentVolumeClaim
}

func (r *GlobalPythonPackageCollectionReconciler) getPersistentVolumeClaimNameForGPPC(gppc *ppv1alpha1.GlobalPythonPackageCollection) string {
	return gppc.Name + "-pvc"
}

// SetupWithManager sets up the controller with the Manager.
func (r *GlobalPythonPackageCollectionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ppv1alpha1.GlobalPythonPackageCollection{}).
		Owns(&v1.PersistentVolumeClaim{}).
		Complete(r)
}
