package vultr

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccVultrSSHKey(t *testing.T) {

	rName := fmt.Sprintf("%s-%d-terraform", acctest.RandString(3), acctest.RandInt())
	rSSH, _, err := acctest.RandSSHKeyPair("foobar")
	name := "data.vultr_ssh_key.my_key"
	if err != nil {
		t.Fatalf("Error generating test SSH key pair: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVultrSSHKeyConfig_basic(rName, rSSH),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", rName),
					resource.TestCheckResourceAttrSet(name, "ssh_key"),
					resource.TestCheckResourceAttrSet(name, "date_created"),
				),
			},
			{
				Config:      testAccCheckVultrSSHKey_noResult(rName),
				ExpectError: regexp.MustCompile(fmt.Sprintf(".*%s: %s: no results were found", name, name)),
			},
		},
	})
}

func testAccCheckVultrSSHKeyConfig_basic(name, ssh string) string {
	return fmt.Sprintf(`
		data "vultr_ssh_key" "my_key" {
    		filter {
    			name = "name"
    			values = ["${vultr_ssh_key.foo.name}"]
			}
  		}

		resource "vultr_ssh_key" "foo" {
			name = "%s"
			ssh_key = "%s"
		}
		`, name, ssh)
}

func testAccCheckVultrSSHKey_noResult(name string) string {
	return fmt.Sprintf(`
		data "vultr_ssh_key" "my_key" {
   			filter {
   				name = "name"
   				values = ["%s"]
			}
 		}`, name)
}