# Changelog

All notable changes to this synthetic repository will be documented in this file.

## [1.2.1] - 2026-04-18

### Fixed

- Corrected the sandbox RPC health response.
- Clarified the upgrade guide for patch adoption.

## [2.0.0] - 2026-04-18

### Changed

- Raised the minimum supported database schema version to `2`.
- Renamed the execution API field `blockGasLimit` to `maxBlockGas`.

### Removed

- Removed compatibility behavior for legacy transaction envelopes.

### Upgrade Notes

- This is an intentionally breaking release for validation coverage.

## [1.2.0-rc.1] - 2026-04-18

### Added

- Added release candidate plumbing for the sandbox consensus engine.
- Added explicit prerelease notes for worker classification testing.

## [1.1.3] - 2026-04-18

### Fixed

- Tightened consensus timing defaults.
- Updated changelog wording for a maintenance-style release.
