// Copyright 2018 The Terraformer Authors.
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

// AUTO-GENERATED CODE. DO NOT EDIT.
package gcp

import (
	"context"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"google.golang.org/api/compute/v1"
)

var autoscalersAllowEmptyValues = []string{""}

var autoscalersAdditionalFields = map[string]string{}

type AutoscalersGenerator struct {
	GCPService
}

// Run on autoscalersList and create for each TerraformResource
func (g AutoscalersGenerator) createResources(ctx context.Context, autoscalersList *compute.AutoscalersListCall, zone string) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := autoscalersList.Pages(ctx, func(page *compute.AutoscalerList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				zone+"/"+obj.Name,
				obj.Name,
				"google_compute_autoscaler",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
					"zone":    zone,
				},
				autoscalersAllowEmptyValues,
				autoscalersAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each autoscalers create 1 TerraformResource
// Need autoscalers name as ID for terraform resource
func (g *AutoscalersGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, zoneLink := range g.GetArgs()["region"].(compute.Region).Zones {
		t := strings.Split(zoneLink, "/")
		zone := t[len(t)-1]
		autoscalersList := computeService.Autoscalers.List(g.GetArgs()["project"].(string), zone)
		g.Resources = g.createResources(ctx, autoscalersList, zone)
	}

	g.PopulateIgnoreKeys()
	return nil

}
