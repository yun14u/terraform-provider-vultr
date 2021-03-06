package vultr

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVultrDnsDomainDataBase(t *testing.T) {
	domain := fmt.Sprintf("%s.com", acctest.RandString(6))
	name := "data.vultr_dns_domain.my-site"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVultrDNSDomainConfig(domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttr(name, "domain", domain),
					resource.TestCheckResourceAttrSet(name, "date_created"),
				),
			},
		},
	})
}

func testAccVultrDNSDomainConfig(domain string) string {
	return fmt.Sprintf(`
			data "vultr_dns_domain" "my-site" {
				domain = "${vultr_dns_domain.my-site.id}"
			}

			resource "vultr_dns_domain" "my-site" {
				domain = "%s"
				ip = "10.0.0.0"
			}`, domain)
}
