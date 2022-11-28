// Code generated by codegen; DO NOT EDIT.

package images

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_images",
		Resolver:  fetchImages,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "filename",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filename"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "require_signed_urls",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RequireSignedURLs"),
			},
			{
				Name:     "variants",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Variants"),
			},
			{
				Name:     "uploaded",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Uploaded"),
			},
		},
	}
}
