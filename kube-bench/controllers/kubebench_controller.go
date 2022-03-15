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
	"time"

	"github.com/ruiyaoOps/kubeeye-plugins/kube-bench/pkg/audit"
	kubeErr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kubeeyev1alpha1 "github.com/ruiyaoOps/kubeeye-plugins/kube-bench/api/v1alpha1"
)

// KubeBenchReconciler reconciles a KubeBench object
type KubeBenchReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubeeye.kubesphere.io,resources=kubebenches,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubeeye.kubesphere.io,resources=kubebenches/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubeeye.kubesphere.io,resources=kubebenches/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=nodes,verbs=*
//+kubebuilder:rbac:groups="",resources=namespaces,verbs=*
//+kubebuilder:rbac:groups="",resources=events,verbs=*
//+kubebuilder:rbac:groups=batch,resources=*,verbs=*
//+kubebuilder:rbac:groups=apps,resources=*,verbs=*
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=*,verbs=*
//+kubebuilder:rbac:groups=storage.k8s.io,resources=*,verbs=*

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KubeBench object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *KubeBenchReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logs := log.FromContext(ctx)
	logs.Info("starting KubeBench audit")

	kubebench := &kubeeyev1alpha1.KubeBench{}
	if err := r.Get(ctx, req.NamespacedName, kubebench); err != nil {
		if kubeErr.IsNotFound(err) {
			logs.Info("Cluster resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
	}

	KubeBenchAudit := audit.KubeBenchAudit(logs)

	keControls := formatResults(KubeBenchAudit)

	kubebench.Status.Controls = keControls

	// update CR status,
	if err := r.Status().Update(ctx, kubebench); err != nil {
		logs.Error(err, "Update CR Status failed")
		return ctrl.Result{}, err
	}

	logs.Info("KubeBench audit completed")

	// If auditPeriod is not set, set the default value to 24h
	if kubebench.Spec.AuditPeriod == "" {
		kubebench.Spec.AuditPeriod = "24h"
	}

	reconcilePeriod, err := time.ParseDuration(kubebench.Spec.AuditPeriod)
	if err != nil {
		logs.Error(err, "AuditPeriod setting is invalid")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: reconcilePeriod}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeBenchReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeeyev1alpha1.KubeBench{}).
		Complete(r)
}

func formatResults(KubeBenchAudit audit.KubeBenchResponse) (keControls []kubeeyev1alpha1.Controls) {
	for _, controls := range KubeBenchAudit.Controls {
		var keControl kubeeyev1alpha1.Controls
		keControl.Version = controls.Version
		keControl.Text = controls.Text
		keControl.Summary = controls.Summary
		for _, group := range controls.Groups {
			var keGroup kubeeyev1alpha1.Group
			keGroup.ID = group.ID
			keGroup.Text = group.Text
			for _, checks := range group.Checks {
				if checks.State != "PASS" {
					var keCheck kubeeyev1alpha1.Check
					keCheck.ID = checks.ID
					keCheck.Text = checks.Text
					keCheck.Remediation = checks.Remediation
					keCheck.State = checks.State
					keGroup.Checks = append(keGroup.Checks, keCheck)
				}
			}
			keControl.Groups = append(keControl.Groups, keGroup)
		}
		keControls = append(keControls, keControl)
	}

	return keControls
}
