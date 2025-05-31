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
2. Commit any changes, tag the new release, and push to GitHub:
   ```bash
   # switch to main branch and update
   git checkout main
   git pull origin main

   # commit any pending changes (e.g., version bump, changelog)
   git add .
   git commit -m "chore: release vX.Y.Z"

   # create an annotated tag for the release
   git tag -a vX.Y.Z -m "Release vX.Y.Z"

   # push commits and tags
   git push origin main --tags
   ```
3. Run goreleaser:
   ```bash
   goreleaser release --clean --config dist/config.yaml
   ```

This will build all targets defined in `dist/config.yaml`, generate archives, checksums, and publish a GitHub Release.

For snapshot builds (local, no publish):
```bash
goreleaser release --snapshot --clean --config dist/config.yaml
```

### Homebrew Tap

We maintain a Homebrew tap in this repository under the `Formula/` directory. On each release, Goreleaser will update the Homebrew formula and push it back to this repo.

To tap and install Homebrew formula:
```bash
brew tap Leathal1/greycli
brew install greycli
```

### Scoop Bucket

We maintain a Scoop bucket in the `scoop/` directory of this repository. On each release, GoReleaser will update the Scoop manifest and push it back.

To add the bucket and install:
```powershell
scoop bucket add greycli https://github.com/Leathal1/greycli.git
scoop install greycli
```
