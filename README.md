# manifestlib-go
Shared configuration and manifest handling library for any tool that consumes the same structure.

## What it provides

- Embedded JSON schemas for `configuration.yaml` and `manifest.yaml`
- Runtime validation against those schemas
- Generic typed access to tool blocks under `configuration.metadata.structure.<tool>`
- Raw YAML loading helpers for flexible consumption

## Functions

- `ValidateInputs(configPath, manifestPath)`
- `LoadConfiguration(path)`
- `LoadManifestRoot(path)`
- `LoadRawConfig(path)`
- `LoadRawManifest(path)`
- `(*ConfigurationFile).ToolIncludePaths(tool)`
- `(*ManifestRoot).IncludePath(tool)`

## Local Development

Until this module is published, consumers should use a local `replace` in `go.mod`:

```go
require github.com/msundalskliev/manifestlib-go v0.0.0
replace github.com/msundalskliev/manifestlib-go => ../manifestlib-go
```

## Package

- `configschema`

## Example usage

```go
err := configschema.ValidateInputs(configPath, manifestPath)
if err != nil {
    return err
}

cfg, err := configschema.LoadConfiguration(configPath)
if err != nil {
    return err
}

manifest, err := configschema.LoadManifestRoot(manifestPath)
if err != nil {
    return err
}

toolKey := "my-tool"
includes := cfg.ToolIncludePaths(toolKey)
entry := manifest.IncludePath(toolKey)

_ = includes
_ = entry
```
