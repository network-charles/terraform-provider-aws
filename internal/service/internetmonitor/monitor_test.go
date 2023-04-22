package internetmonitor_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/internetmonitor"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfinternetmonitor "github.com/hashicorp/terraform-provider-aws/internal/service/internetmonitor"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func TestAccInternetMonitorMonitor_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_internetmonitor_monitor.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, internetmonitor.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitorConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "internetmonitor", regexp.MustCompile(`monitor/.+$`)),
					resource.TestCheckResourceAttr(resourceName, "monitor_name", rName),
					resource.TestCheckResourceAttr(resourceName, "taffic_percentage_to_monitor", "1"),
					resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccMonitorConfig_status(rName, "INACTIVE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "status", "INACTIVE"),
				),
			},
		},
	})
}

func TestAccInternetMonitorMonitor_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_internetmonitor_monitor.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, internetmonitor.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitorConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfinternetmonitor.ResourceMonitor(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccInternetMonitorMonitor_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_internetmonitor_monitor.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, internetmonitor.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckMonitorDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitorConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccMonitorConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccMonitorConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func testAccCheckMonitorDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).InternetMonitorConn()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_internetmonitor_monitor" {
				continue
			}

			_, err := tfinternetmonitor.FindMonitor(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("InternetMonitor Monitor %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckMonitorExists(ctx context.Context, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No InternetMonitor Monitor ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).InternetMonitorConn()

		_, err := tfinternetmonitor.FindMonitor(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccMonitorConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_internetmonitor_monitor" "test" {
  monitor_name                 = %[1]q
  taffic_percentage_to_monitor = 1
}
`, rName)
}

func testAccMonitorConfig_status(rName, status string) string {
	return fmt.Sprintf(`
resource "aws_internetmonitor_monitor" "test" {
  monitor_name                 = %[1]q
  taffic_percentage_to_monitor = 1
  status                       = %[2]q
}
`, rName, status)
}

func testAccMonitorConfig_tags1(rName string, tagKey1 string, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_internetmonitor_monitor" "test" {
  monitor_name                 = %[1]q
  taffic_percentage_to_monitor = 1

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccMonitorConfig_tags2(rName string, tagKey1 string, tagValue1 string, tagKey2 string, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_internetmonitor_monitor" "test" {
  monitor_name                 = %[1]q
  taffic_percentage_to_monitor = 1

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
