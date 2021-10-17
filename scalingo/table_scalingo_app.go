package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableScalingoApp() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_app",
		Description: "An application is the program that is running on Scalingo.",
		List: &plugin.ListConfig{
			Hydrate: listApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			Hydrate:           getApp,
			ShouldIgnoreError: isNotFoundError,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "Unique id of the application."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the application."},
			{Name: "region", Type: proto.ColumnType_STRING, Description: "Region of the application."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the application."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.ID"), Description: "Unique id of the owner."},
			{Name: "owner_username", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.Username"), Description: "Username of the owner."},
			{Name: "owner_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Owner.Email"), Description: "Email of the owner."},
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "Creation date of the application."},
			{Name: "updated_at", Type: proto.ColumnType_DATETIME, Description: "Last time the application has been updated."},
			{Name: "git_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("GitUrl"), Description: "URL to the GIT remote to access your application."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Url"), Description: "URL used to access your app."},
			{Name: "base_url", Type: proto.ColumnType_STRING, Description: "URL generated by Scalingo for your app."},
			{Name: "force_https", Type: proto.ColumnType_BOOL, Description: "Activation of force HTTPS."},
			{Name: "sticky_session", Type: proto.ColumnType_BOOL, Description: "Activation of sticky session."},
			{Name: "router_logs", Type: proto.ColumnType_BOOL, Description: "Activation of the router logs in your app logs."},
			{Name: "last_deployed_at", Type: proto.ColumnType_DATETIME, Description: "Date of the last deployment attempt."},
			{Name: "last_deployed_by", Type: proto.ColumnType_STRING, Description: "User who attempted the last deployment."},
			{Name: "stack_id", Type: proto.ColumnType_STRING, Description: "Id of the stack used."},
		},
	}
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	apps, err := client.AppsList()
	if err != nil {
		return nil, err
	}
	for _, app := range apps {
		d.StreamListItem(ctx, app)
	}
	return nil, nil
}

func getApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals
	name := quals["name"].GetStringValue()
	result, err := client.AppsShow(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}
