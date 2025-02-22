// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infrastructure

import (
	"context"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/extensions/pkg/controller/infrastructure"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type actuator struct {
	client     client.Client
	restConfig *rest.Config
	scheme     *runtime.Scheme
	gardenID   string
}

// NewActuator creates a new Actuator that updates the status of the handled Infrastructure resources.
func NewActuator(mgr manager.Manager, gardenID string) infrastructure.Actuator {
	return &actuator{
		client:     mgr.GetClient(),
		restConfig: mgr.GetConfig(),
		scheme:     mgr.GetScheme(),
		gardenID:   gardenID,
	}
}

func (a *actuator) Reconcile(ctx context.Context, log logr.Logger, config *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	return a.reconcile(ctx, log, config, cluster)
}

func (a *actuator) Delete(ctx context.Context, log logr.Logger, config *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	return a.delete(ctx, log, config, cluster)
}
