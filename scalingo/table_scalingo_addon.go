package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableScalingoAddon() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_addon",
		Description: "An addon is a database or similar associated to an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listAddon,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    getAddon,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the addon."},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application which owns the addon."},
			{Name: "resource_id", Type: proto.ColumnType_STRING, Description: "Resource reference."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Current status of the addon."},
			{Name: "provisioned_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the addon has been created."},
			{Name: "deprovisioned_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the addon has been removed/upgraded."},
			{Name: "provider_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.ID"), Description: "Id of the provider."},
			{Name: "provider_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.Name"), Description: "Name of the provider."},
			{Name: "provider_logo_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.LogoURL"), Description: "Logo url of the provider."},
			{Name: "provider_hds_available", Type: proto.ColumnType_BOOL, Transform: transform.FromField("AddonProvider.HDSAvailable"), Description: "Is this addon HDS compliantName of the provider."},
			{Name: "provider_short_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.ShortDescription"), Description: "Short description of the provider."},
			{Name: "provider_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("AddonProvider.Description"), Description: "Description of the provider."},
			{Name: "plan_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.ID"), Description: "Id of the plan."},
			{Name: "plan_display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.DisplayName"), Description: "Display name of the plan."},
			{Name: "plan_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.Name"), Description: "Name of the plan."},
			{Name: "plan_description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.Description"), Description: "Description of the plan."},
			{Name: "plan_price", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Plan.Price"), Description: "Price of the plan."},
			{Name: "plan_sku", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.SKU"), Description: "SKU of the plan."},
			{Name: "plan_on_demand", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Plan.OnDemand"), Description: "Is this addon available only on demand."},
			{Name: "plan_disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Plan.Disabled"), Description: "Is this plan disabled."},
			{Name: "plan_disabled_alternative_plan_id", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Plan.DisabledAlternativePlanID"), Description: "ID of other plan if this plan disabled."},
			{Name: "plan_hds_available", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Plan.HDSAvailable"), Description: "is this plan HDS compliant."},
		},
	}
}

func listAddon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_addon.listAddon", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	addons, err := client.AddonsList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_addon.listAddon", err)
		return nil, err
	}
	for _, addon := range addons {
		d.StreamListItem(ctx, addon)
	}
	return nil, nil
}

func getAddon(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_addon.getAddon", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()

	result, err := client.AddonShow(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_addon.getAddon", err)
		return nil, err
	}
	return result, nil
}
