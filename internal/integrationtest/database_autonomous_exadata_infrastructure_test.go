// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_database "github.com/oracle/oci-go-sdk/v55/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousExadataInfrastructureRequiredOnlyResource = AutonomousExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Required, acctest.Create, autonomousExadataInfrastructureRepresentation)

	AutonomousExadataInfrastructureResourceConfig = AutonomousExadataInfrastructureResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Update, autonomousExadataInfrastructureRepresentation)

	autonomousExadataInfrastructureSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
	}

	autonomousExadataInfrastructureDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domain.ad.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `tst3dbsys`, Update: `displayName2`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousExadataInfrastructureDataSourceFilterRepresentation}}
	autonomousExadataInfrastructureDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`}},
	}

	autonomousExadataInfrastructureRepresentation = map[string]interface{}{
		"availability_domain":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domain.ad.name}`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                      acctest.Representation{RepType: acctest.Required, Create: `Exadata.X8M`},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.exadata_subnet.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `tst3dbsys`, Update: `displayName2`},
		"domain":                     acctest.Representation{RepType: acctest.Optional, Create: `subnetexadata.tfvcn.oraclevcn.com`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"maintenance_window_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsRepresentation},
		"nsg_ids":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{`${oci_core_network_security_group.test_network_security_group2.id}`}},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsRepresentation = map[string]interface{}{
		"preference":     acctest.Representation{RepType: acctest.Required, Create: `NO_PREFERENCE`, Update: `CUSTOM_PREFERENCE`},
		"days_of_week":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsDaysOfWeekRepresentation},
		"hours_of_day":   acctest.Representation{RepType: acctest.Optional, Create: []string{`4`}, Update: []string{`8`}},
		"months":         []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation4}},
		"weeks_of_month": acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}, Update: []string{`2`}},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsDaysOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JANUARY`, Update: `FEBRUARY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `APRIL`, Update: `MAY`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation3 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JULY`, Update: `AUGUST`},
	}
	autonomousExadataInfrastructureMaintenanceWindowDetailsMonthsRepresentation4 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `OCTOBER`, Update: `NOVEMBER`},
	}

	AutonomousExadataInfrastructureResourceDependencies = ExadataBaseDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("vcn_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("vcn_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation))
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousExadataInfrastructureResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure"
	datasourceName := "data.oci_database_autonomous_exadata_infrastructures.test_autonomous_exadata_infrastructures"
	singularDatasourceName := "data.oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AutonomousExadataInfrastructureResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Create, autonomousExadataInfrastructureRepresentation), "database", "autonomousExadataInfrastructure", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousExadataInfrastructureDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Required, acctest.Create, autonomousExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Create, autonomousExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tst3dbsys"),
				resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutonomousExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousExadataInfrastructureRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "tst3dbsys"),
				resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Update, autonomousExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hostname"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructures", "test_autonomous_exadata_infrastructures", acctest.Optional, acctest.Update, autonomousExadataInfrastructureDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousExadataInfrastructureResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Optional, acctest.Update, autonomousExadataInfrastructureRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.domain", "subnetexadata.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.hostname"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.shape", "Exadata.X8M"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_exadata_infrastructures.0.nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructures.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Required, acctest.Create, autonomousExadataInfrastructureSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousExadataInfrastructureResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_exadata_infrastructure_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", "subnetexadata.tfvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape", "Exadata.X8M"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousExadataInfrastructureResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"maintenance_window_details",
				"create_async",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseAutonomousExadataInfrastructureDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_exadata_infrastructure" {
			noResourceFound = false
			request := oci_database.GetAutonomousExadataInfrastructureRequest{}

			tmp := rs.Primary.ID
			request.AutonomousExadataInfrastructureId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAutonomousExadataInfrastructure(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousExadataInfrastructureLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseAutonomousExadataInfrastructure") {
		resource.AddTestSweepers("DatabaseAutonomousExadataInfrastructure", &resource.Sweeper{
			Name:         "DatabaseAutonomousExadataInfrastructure",
			Dependencies: acctest.DependencyGraph["autonomousExadataInfrastructure"],
			F:            sweepDatabaseAutonomousExadataInfrastructureResource,
		})
	}
}

func sweepDatabaseAutonomousExadataInfrastructureResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousExadataInfrastructureIds, err := getAutonomousExadataInfrastructureIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousExadataInfrastructureId := range autonomousExadataInfrastructureIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousExadataInfrastructureId]; !ok {
			terminateAutonomousExadataInfrastructureRequest := oci_database.TerminateAutonomousExadataInfrastructureRequest{}

			terminateAutonomousExadataInfrastructureRequest.AutonomousExadataInfrastructureId = &autonomousExadataInfrastructureId

			terminateAutonomousExadataInfrastructureRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.TerminateAutonomousExadataInfrastructure(context.Background(), terminateAutonomousExadataInfrastructureRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousExadataInfrastructure %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousExadataInfrastructureId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousExadataInfrastructureId, autonomousExadataInfrastructureSweepWaitCondition, time.Duration(3*time.Minute),
				autonomousExadataInfrastructureSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getAutonomousExadataInfrastructureIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousExadataInfrastructureId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousExadataInfrastructuresRequest := oci_database.ListAutonomousExadataInfrastructuresRequest{}
	listAutonomousExadataInfrastructuresRequest.CompartmentId = &compartmentId
	listAutonomousExadataInfrastructuresRequest.LifecycleState = oci_database.AutonomousExadataInfrastructureSummaryLifecycleStateAvailable
	listAutonomousExadataInfrastructuresResponse, err := databaseClient.ListAutonomousExadataInfrastructures(context.Background(), listAutonomousExadataInfrastructuresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousExadataInfrastructure list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousExadataInfrastructure := range listAutonomousExadataInfrastructuresResponse.Items {
		id := *autonomousExadataInfrastructure.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousExadataInfrastructureId", id)
	}
	return resourceIds, nil
}

func autonomousExadataInfrastructureSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousExadataInfrastructureResponse, ok := response.Response.(oci_database.GetAutonomousExadataInfrastructureResponse); ok {
		return autonomousExadataInfrastructureResponse.LifecycleState != oci_database.AutonomousExadataInfrastructureLifecycleStateTerminated
	}
	return false
}

func autonomousExadataInfrastructureSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousExadataInfrastructure(context.Background(), oci_database.GetAutonomousExadataInfrastructureRequest{
		AutonomousExadataInfrastructureId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
