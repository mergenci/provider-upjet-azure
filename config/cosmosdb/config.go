// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cosmosdb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures cosmodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_cosmosdb_account", func(r *config.Resource) {
		delete(r.References, "geo_location.location")
	})
	p.AddResourceConfigurator("azurerm_cosmosdb_sql_container", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "SQLDatabase",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_mongo_collection", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "MongoDatabase",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_cassandra_table", func(r *config.Resource) {
		delete(r.References, "cassandra_keyspace_id")
		r.References["cassandra_keyspace_id"] = config.Reference{
			Type:      "CassandraKeySpace",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"analytical_storage_ttl"},
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_gremlin_graph", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "GremlinDatabase",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_function", func(r *config.Resource) {
		r.References["container_id"] = config.Reference{
			Type:      "SQLContainer",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_stored_procedure", func(r *config.Resource) {
		r.References["database_name"] = config.Reference{
			Type: "SQLDatabase",
		}
		r.References["container_name"] = config.Reference{
			Type: "SQLContainer",
		}
	})

	p.AddResourceConfigurator("azurerm_cosmosdb_sql_trigger", func(r *config.Resource) {
		r.References = config.References{
			"container_id": config.Reference{
				Type:      "SQLContainer",
				Extractor: rconfig.ExtractResourceIDFuncPath,
			},
		}
	})
}
