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
	"encoding/json"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ppv1alpha1 "github.com/akakakakakaa/python-package-operator/api/v1alpha1"
	"github.com/akakakakakaa/python-package-operator/util"
	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
)

// GlobalPythonPackageReconciler reconciles a GlobalPythonPackage object
type GlobalPythonPackageReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackages,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackages/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pythonpackage.github.io,resources=globalpythonpackages/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete

func (r *GlobalPythonPackageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("gpp", req.NamespacedName)

	gpp := &ppv1alpha1.GlobalPythonPackage{}
	if err := r.Get(context.TODO(), req.NamespacedName, gpp); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("GlobalPythonPackage resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get GlobalPythonPackage ["+req.Name+"].")
		return ctrl.Result{}, err
	}

	//set patch helper
	patchHelper, err := patch.NewHelper(gpp, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	defer func() {
		if err := patchHelper.Patch(context.TODO(), gpp); err != nil {
			logger.Error(err, "Failed to patch GlobalPythonPackage.")
		}
	}()

	gppcFound := &ppv1alpha1.GlobalPythonPackageCollection{}
	gppcName := gpp.Spec.GlobalPythonPackageCollectionName
	if err := r.Get(context.TODO(), types.NamespacedName{Name: gppcName}, gppcFound); err != nil {
		if errors.IsNotFound(err) {
			logger.Error(err, "Cannot found globalPythonPackageCollection ["+gppcName+"].")
			gpp.Status.Reason = "Cannot found globalPythonPackageCollection"
		} else {
			logger.Error(err, "Failed to get globalPythonPackageCollection ["+gppcName+"].")
			gpp.Status.Reason = "Failed to get globalPythonPackageCollection"
		}
		gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
		gpp.Status.Message = err.Error()
		return ctrl.Result{}, err
	} else {
		ctrl.SetControllerReference(gppcFound, gpp, r.Scheme)
	}

	switch gpp.Status.Status {
	case "":
		found := &batchv1.Job{}
		err = r.Get(context.TODO(), types.NamespacedName{Name: jobNameForGPP(gpp), Namespace: util.Namespace()}, found)
		if err != nil && errors.IsNotFound(err) {
			logger.Info("Create job.")
			job, err := r.jobForGPP(gppcFound, gpp)
			if err != nil {
				logger.Error(err, "Failed to get job.")
				gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
				gpp.Status.Reason = "Failed to get job"
				gpp.Status.Message = err.Error()
				return ctrl.Result{}, err
			}
			ctrl.SetControllerReference(gpp, job, r.Scheme)

			if err := r.Create(context.TODO(), job); err != nil {
				logger.Error(err, "Failed to create job.")
				gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
				gpp.Status.Reason = "Failed to create job"
				gpp.Status.Message = err.Error()
				return ctrl.Result{}, err
			}

			logger.Info("Create job success.")
			gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeCreated
			gpp.Status.Reason = "Create job success."
		} else if err != nil {
			logger.Error(err, "Failed to get job.")
			gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
			gpp.Status.Reason = "Failed to create job"
			gpp.Status.Message = err.Error()
			return ctrl.Result{}, err
		} else {
			logger.Error(err, "Duplicate job name exists.")
			gpp.Status.Status = ppv1alpha1.GlobalPythonPackageCollectionStatusTypeError
			gpp.Status.Reason = "Duplicate job name exists"
			gpp.Status.Message = err.Error()
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *GlobalPythonPackageReconciler) jobForGPP(gppc *ppv1alpha1.GlobalPythonPackageCollection, gpp *ppv1alpha1.GlobalPythonPackage) (*batchv1.Job, error) {
	gppSpecJson, err := json.Marshal(gpp.Spec)
	if err != nil {
		return nil, err
	}

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobNameForGPP(gpp),
			Namespace: util.Namespace(),
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:  jobContainerNameForGPP(gpp),
						Image: os.Getenv("GLOBALPYTHONPACKAGE_JOB_NAME"),
						Args:  []string{},
						Env: []v1.EnvVar{{
							Name:  "GPP_JSON_STRING",
							Value: string(gppSpecJson),
						}},
						VolumeMounts: []v1.VolumeMount{{
							MountPath: "/mnt",
							Name:      persistentVolumeClaimNameForGPPC(gppc),
						}},
					}},
					RestartPolicy: "Never",
				},
			},
		},
	}

	return job, nil
}

func jobNameForGPP(gpp *ppv1alpha1.GlobalPythonPackage) string {
	return gpp.Name + "-job"
}

func jobContainerNameForGPP(gpp *ppv1alpha1.GlobalPythonPackage) string {
	return jobNameForGPP(gpp) + "-container"
}

// SetupWithManager sets up the controller with the Manager.
func (r *GlobalPythonPackageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ppv1alpha1.GlobalPythonPackage{}).
		Owns(&batchv1.Job{}).
		Complete(r)
}
