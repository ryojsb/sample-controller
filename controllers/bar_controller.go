/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	samplecontrollerv1 "my.domain/sample-controller/api/v1"
)

// BarReconciler reconciles a Bar object
type BarReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=samplecontroller.my.domain,resources=bars,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=samplecontroller.my.domain,resources=bars/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=samplecontroller.my.domain,resources=bars/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Bar object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *BarReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	Log := log.FromContext(ctx)

	Log.Info("*Reconcile called*")

	var bar samplecontrollerv1.Bar

	// Golangのstructの埋め込みを利用している
	// BarReconcilerの実装ではなく、実態は埋め込んでいるclient.ClientのGetを呼んでいる
	// clientの内部キャッシュに含まれている情報を元にオブジェクトを取得してきている
	if err := r.Get(ctx, req.NamespacedName, &bar); err != nil {
		Log.Error(err, "failed to fetching foo")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	Log.Info("[Debug] Bar: " + bar.Namespace + "/" + bar.Name + "from in-memory-chache")

	message := "Hello " + bar.Spec.Message
	Log.Info("[Debug] Message: " + message)

	if bar.Status.Message != message {
		bar.Status.Message = message
		if err := r.Status().Update(ctx, &bar); err != nil {
			Log.Error(err, "Failed to update Foo status")
			return ctrl.Result{}, err
		}
		Log.Info("[Debug] Update status of Bar: " + bar.Name)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BarReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&samplecontrollerv1.Bar{}).
		Complete(r)
}
