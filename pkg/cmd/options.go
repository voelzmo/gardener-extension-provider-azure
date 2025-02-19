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

package cmd

import (
	backupbucketcontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/backupbucket"
	backupentrycontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/backupentry"
	controlplanecontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/controlplane"
	csimigrationcontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/csimigration"
	dnsrecordcontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/dnsrecord"
	healthcheckcontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/healthcheck"
	infrastructurecontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/infrastructure"
	workercontroller "github.com/gardener/gardener-extension-provider-azure/pkg/controller/worker"
	controlplanewebhook "github.com/gardener/gardener-extension-provider-azure/pkg/webhook/controlplane"
	controlplaneexposurewebhook "github.com/gardener/gardener-extension-provider-azure/pkg/webhook/controlplaneexposure"
	networkwebhook "github.com/gardener/gardener-extension-provider-azure/pkg/webhook/network"

	extensionsbackupbucketcontroller "github.com/gardener/gardener/extensions/pkg/controller/backupbucket"
	extensionsbackupentrycontroller "github.com/gardener/gardener/extensions/pkg/controller/backupentry"
	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	extensionscontrolplanecontroller "github.com/gardener/gardener/extensions/pkg/controller/controlplane"
	extensionscsimigrationcontroller "github.com/gardener/gardener/extensions/pkg/controller/csimigration"
	extensionsdnsrecordcontroller "github.com/gardener/gardener/extensions/pkg/controller/dnsrecord"
	extensionshealthcheckcontroller "github.com/gardener/gardener/extensions/pkg/controller/healthcheck"
	extensionsinfrastructurecontroller "github.com/gardener/gardener/extensions/pkg/controller/infrastructure"
	extensionsworkercontroller "github.com/gardener/gardener/extensions/pkg/controller/worker"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"
	extensionscontrolplanewebhook "github.com/gardener/gardener/extensions/pkg/webhook/controlplane"
	extensionsnetworkwebhook "github.com/gardener/gardener/extensions/pkg/webhook/network"
)

// ControllerSwitchOptions are the controllercmd.SwitchOptions for the provider controllers.
func ControllerSwitchOptions() *controllercmd.SwitchOptions {
	return controllercmd.NewSwitchOptions(
		controllercmd.Switch(extensionsbackupbucketcontroller.ControllerName, backupbucketcontroller.AddToManager),
		controllercmd.Switch(extensionsbackupentrycontroller.ControllerName, backupentrycontroller.AddToManager),
		controllercmd.Switch(extensionscontrolplanecontroller.ControllerName, controlplanecontroller.AddToManager),
		controllercmd.Switch(extensionscsimigrationcontroller.ControllerName, csimigrationcontroller.AddToManager),
		controllercmd.Switch(extensionsdnsrecordcontroller.ControllerName, dnsrecordcontroller.AddToManager),
		controllercmd.Switch(extensionsinfrastructurecontroller.ControllerName, infrastructurecontroller.AddToManager),
		controllercmd.Switch(extensionsworkercontroller.ControllerName, workercontroller.AddToManager),
		controllercmd.Switch(extensionshealthcheckcontroller.ControllerName, healthcheckcontroller.AddToManager),
	)
}

// WebhookSwitchOptions are the webhookcmd.SwitchOptions for the provider webhooks.
func WebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(extensionsnetworkwebhook.WebhookName, networkwebhook.AddToManager),
		webhookcmd.Switch(extensionscontrolplanewebhook.WebhookName, controlplanewebhook.AddToManager),
		webhookcmd.Switch(extensionscontrolplanewebhook.ExposureWebhookName, controlplaneexposurewebhook.AddToManager),
	)
}
