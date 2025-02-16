module github.com/DataDog/datadog-agent/pkg/security/seclwin

go 1.23.0

replace github.com/DataDog/datadog-agent/pkg/security/secl => ../secl

require github.com/DataDog/datadog-agent/pkg/security/secl v0.56.0-rc.3

require (
	github.com/alecthomas/participle v0.7.1 // indirect
	github.com/jellydator/ttlcache/v3 v3.3.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
)
