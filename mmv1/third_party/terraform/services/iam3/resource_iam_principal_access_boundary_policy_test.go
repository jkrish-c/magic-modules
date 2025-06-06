package iam3_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccIAM3PrincipalAccessBoundaryPolicy_iam3PrincipalAccessBoundaryPolicyExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAM3PrincipalAccessBoundaryPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAM3PrincipalAccessBoundaryPolicy_iam3PrincipalAccessBoundaryPolicyExample_full(context),
			},
			{
				ResourceName:            "google_iam_principal_access_boundary_policy.my-pab-policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "location", "organization", "principal_access_boundary_policy_id", "etag"},
			},
			{
				Config: testAccIAM3PrincipalAccessBoundaryPolicy_iam3PrincipalAccessBoundaryPolicyExample_update(context),
			},
			{
				ResourceName:            "google_iam_principal_access_boundary_policy.my-pab-policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "location", "organization", "principal_access_boundary_policy_id", "etag"},
			},
		},
	})
}

func testAccIAM3PrincipalAccessBoundaryPolicy_iam3PrincipalAccessBoundaryPolicyExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_principal_access_boundary_policy" "my-pab-policy" {
  organization   = "%{org_id}"
  location       = "global"
  display_name   = "test pab policy%{random_suffix}"
  principal_access_boundary_policy_id = "test-pab-policy%{random_suffix}"
}
`, context)
}

func testAccIAM3PrincipalAccessBoundaryPolicy_iam3PrincipalAccessBoundaryPolicyExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_project" "project" {
  project_id     = "tf-test%{random_suffix}"
  name           = "tf-test%{random_suffix}"
  org_id         = "%{org_id}"
  deletion_policy = "DELETE"
}

resource "google_iam_principal_access_boundary_policy" "my-pab-policy" {
  organization   = "%{org_id}"
  location       = "global"
  display_name   = "test pab policy%{random_suffix}"
  principal_access_boundary_policy_id = "test-pab-policy%{random_suffix}"
  annotations    = {"foo": "bar"}
  details {
    rules {
      description = "PAB rule%{random_suffix}"
      effect      = "ALLOW"
      resources   = ["//cloudresourcemanager.googleapis.com/projects/${google_project.project.project_id}"]
    }
    enforcement_version = "1"
  }
}
`, context)
}
