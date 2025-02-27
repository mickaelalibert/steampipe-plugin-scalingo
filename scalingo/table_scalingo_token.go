package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableScalingoToken() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_token",
		Description: "An API token associated to the account.",
		List: &plugin.ListConfig{
			Hydrate: listToken,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the token."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Token name."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Token creation date."},
			{Name: "last_used_at", Type: proto.ColumnType_TIMESTAMP, Description: "Token last used date."},
		},
	}
}

func listToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_token.listToken", "connection_error", err)
		return nil, err
	}
	tokens, err := client.TokensList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_token.listToken", err)
		return nil, err
	}
	for _, token := range tokens {
		d.StreamListItem(ctx, token)
	}

	return nil, nil
}
