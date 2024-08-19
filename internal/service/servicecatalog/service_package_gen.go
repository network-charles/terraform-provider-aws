// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	servicecatalog_sdkv2 "github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
		},
		{
			Factory:  dataSourceLaunchPaths,
			TypeName: "aws_servicecatalog_launch_paths",
			Name:     "Launch Paths",
		},
		{
			Factory:  dataSourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourcePortfolioConstraints,
			TypeName: "aws_servicecatalog_portfolio_constraints",
			Name:     "Portfolio Constraints",
		},
		{
			Factory:  dataSourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceProvisioningArtifacts,
			TypeName: "aws_servicecatalog_provisioning_artifacts",
			Name:     "Provisioning Artifacts",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBudgetResourceAssociation,
			TypeName: "aws_servicecatalog_budget_resource_association",
			Name:     "Budget Resource Association",
		},
		{
			Factory:  resourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
			Name:     "Constraint",
		},
		{
			Factory:  resourceOrganizationsAccess,
			TypeName: "aws_servicecatalog_organizations_access",
			Name:     "Organizations Access",
		},
		{
			Factory:  resourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourcePortfolioShare,
			TypeName: "aws_servicecatalog_portfolio_share",
			Name:     "Portfolio Share",
		},
		{
			Factory:  resourcePrincipalPortfolioAssociation,
			TypeName: "aws_servicecatalog_principal_portfolio_association",
			Name:     "Principal Portfolio Association",
		},
		{
			Factory:  resourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceProductPortfolioAssociation,
			TypeName: "aws_servicecatalog_product_portfolio_association",
			Name:     "Product Portfolio Association",
		},
		{
			Factory:  resourceProvisionedProduct,
			TypeName: "aws_servicecatalog_provisioned_product",
			Name:     "Provisioned Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceProvisioningArtifact,
			TypeName: "aws_servicecatalog_provisioning_artifact",
			Name:     "Provisioning Artifact",
		},
		{
			Factory:  resourceServiceAction,
			TypeName: "aws_servicecatalog_service_action",
			Name:     "Service Action",
		},
		{
			Factory:  resourceTagOption,
			TypeName: "aws_servicecatalog_tag_option",
			Name:     "Tag Option",
		},
		{
			Factory:  resourceTagOptionResourceAssociation,
			TypeName: "aws_servicecatalog_tag_option_resource_association",
			Name:     "Tag Option Resource Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceCatalog
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicecatalog_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return servicecatalog_sdkv2.NewFromConfig(cfg,
		servicecatalog_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
