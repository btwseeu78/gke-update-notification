/*
Copyright 2024.

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

package controller

import (
	"context"
	"errors"
	"github.com/go-logr/logr"

	notificationv1alpha1 "btwseeu78/gke-update-notification/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NotificationConfigReconciler reconciles a NotificationConfig object
type NotificationConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=notification.controller.renault.com,resources=notificationconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=notification.controller.renault.com,resources=notificationconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=notification.controller.renault.com,resources=notificationconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NotificationConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *NotificationConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("notificationconfig", req.NamespacedName)
	log.Info("Reconciling NotificationConfig")

	notificationConfig := &notificationv1alpha1.NotificationConfig{}
	if err := r.Client.Get(ctx, req.NamespacedName, notificationConfig); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	checkInterval := notificationConfig.Spec.CheckInterval.Duration

	metricName := notificationConfig.Spec.Metrics.Name
	metricMetadata := notificationConfig.Spec.Metrics.Labels
	metricString, err := CreateDynaMetrics(metricName, metricMetadata)
	if err != nil {
		return ctrl.Result{RequeueAfter: checkInterval}, err
	}
	log.Info("Metric Data Is", "metricName", metricName,
		"metricMetadata", metricMetadata, "metricsString", metricString)

	return ctrl.Result{RequeueAfter: checkInterval}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NotificationConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&notificationv1alpha1.NotificationConfig{}).
		Complete(r)
}

func CreateDynaMetrics(metricName string, labels map[string]string) (string, error) {
	var tempMetric string
	if metricName == "" {
		return "", errors.New("metric name is required")
	}
	tempMetric = metricName
	for k, v := range labels {
		tempMetric = tempMetric + "," + k + ":" + v
	}
	return tempMetric, nil
}
