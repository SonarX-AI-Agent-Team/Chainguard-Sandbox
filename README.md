# agent-validation-sandbox

`agent-validation-sandbox` is a small synthetic repository shaped like a blockchain client.

It exists to validate release, changelog, and alerting behavior without depending on a production codebase.

## Layout

The repository intentionally mirrors a subset of the `bnb-chain/bsc` structure:

- `params/config.go`
- `core/vm/instructions.go`
- `core/vm/contracts.go`
- `internal/ethapi/api.go`
- `consensus/sandbox/consensus.go`
- `docs/upgrade-guide.md`
- `.github/workflows/ci.yml`
- `.github/workflows/release.yml`

## Intended validation flow

1. Publish tagged releases from this repository.
2. Run the worker against the `sandbox` chain entry.
3. Inspect resulting DB rows and alerts for maintenance, prerelease, breaking, and patch behavior.
