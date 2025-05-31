# greycli

GreyCLI is the official command-line interface and terminal UI for [Greymattr.ai](https://greymattr.ai)'s decentralized AI inference network, Fist.

- Submit jobs to P2P AI nodes
- Monitor job status, credits, and output hashes
- Built in Go, powered by TUI magic

## Installation

### From source

First, ensure you have Go 1.18 or later installed. Then:

```bash
git clone https://github.com/Leathal1/greycli.git
cd greycli
go build -o greycli .
```

This will produce a `greycli` (or `greycli.exe` on Windows) binary in the project root.

To embed version metadata, use `-ldflags`. For example:

```bash
go build -ldflags "\
-X main.version=$(git describe --tags --always) \
-X main.commit=$(git rev-parse HEAD) \
-X main.date=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
-X main.builtBy=$(whoami)" \
-o greycli .
```

### Download pre-built binaries

Pre-compiled binaries are available on the GitHub Releases page:
https://github.com/Leathal1/greycli/releases

## Release (maintainers)

We use [goreleaser](https://goreleaser.com) to automate cross-platform builds, archives, checksums, and publishing to GitHub.

1. Install goreleaser. On macOS:
   ```bash
   brew install goreleaser/tap/goreleaser
   ```
2. Commit any changes and create a version tag:
   ```bash
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```
3. Run goreleaser:
   ```bash
   goreleaser release --rm-dist
   ```

This will build all targets defined in `dist/config.yaml`, generate archives, checksums, and publish a GitHub Release.

For snapshot builds (local, no publish):
```bash
goreleaser release --snapshot --skip-publish --rm-dist
```
