# Upgrade Guide

This repository is synthetic and exists to exercise release classification logic.

## Recommended test sequence

1. Start from `v1.1.3`.
2. Validate prerelease handling with `v1.2.0-rc.1`.
3. Validate breaking-release handling with `v2.0.0`.
4. Validate patch-release handling with `v1.2.1`.

## Breaking release expectations

When upgrading to `v2.0.0`, confirm that downstream tooling detects:

- a major version increment
- explicit upgrade notes
- a field rename in the execution API
- a removal note in the changelog
