package migrations

import "embed"

//go:embed sql/*.sql
var EmbedMigrationsFs embed.FS
