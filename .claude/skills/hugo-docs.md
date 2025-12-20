---
description: Build and manage Hugo documentation site
args: [command]
---

# Hugo Documentation Build Skill

This skill helps you build and manage the Hugo-based documentation site for go-swagger.

## Documentation Architecture

### Source Structure
- **Documentation source**: `./docs/` - All markdown content organized hierarchically
  - Uses `_index.md` files for section landing pages
  - Front matter includes: title, date, draft status, weight for ordering
  - Subdirectories: about, contributing, faq, features, generate, install, reference, tutorial, usage

### Hugo Configuration
- **Hugo root**: `./hack/doc-site/hugo/`
- **Main config**: `hugo.yaml` - Primary Hugo configuration
  - Theme: hugo-book (v9 from alex-shpak/hugo-book)
  - Base URL: https://goswagger.io/go-swagger
  - Key feature: Mounts `../../../docs` as content directory

- **Dynamic config**: `goswagger.yaml` - Generated at build time with:
  - Latest release version (from git tags)
  - Required Go version (from go.mod)
  - Version message for the documentation set

### Customizations

#### Theme Mounts
1. **go-swagger-assets** - Custom SCSS styles
   - `_custom.scss` for style overrides
   - `mermaid.json` for diagram configuration

2. **go-swagger-static** - Static assets
   - `logo.png` - Project logo
   - `favicon.png` - Site favicon
   - Additional images (cloudsmith.png, slack.png)

#### Custom Layouts
Located in `layouts/`:
- **Shortcodes**:
  - `forkme.html` - "Fork me on GitHub" ribbon
  - `rawhtml.html` - Embed raw HTML

- **Partials** (`partials/docs/inject/`):
  - `footer.html` - Custom footer with copyright and attribution
  - `menu-before.html`, `menu-after.html` - Menu customizations
  - `toc-before.html` - Table of contents customizations

- **Markup renderers** (`_default/_markup/`):
  - `render-link.html` - Custom link rendering for internal references
  - `render-image.html` - Custom image rendering

### Hugo Configuration Highlights

```yaml
theme: hugo-book
enableEmoji: true
enableGitInfo: true
enableRobotsTXT: true

# Key mounts:
- source: '../../../docs'           # Markdown source
  target: 'content'
  excludeFiles: 'presentations/**'

- source: '../../../docs/presentations'
  target: 'static/presentations'    # Raw HTML presentations

params:
  BookTheme: "dark"
  BookToC: true
  BookRepo: 'https://github.com/go-swagger/go-swagger'
  BookEditPath: edit/master
  contentDir: docs
  BookSearch: true
```

## Build Commands

### Local Development
```bash
cd hack/doc-site/hugo
./gendoc.sh
```

This script:
1. Clones hugo-book theme
2. Runs Hugo server with:
   - Both config files (hugo.yaml + goswagger.yaml)
   - Draft mode enabled
   - Minification
   - Live reload

### CI/CD Build Process

The `.github/workflows/update-doc.yaml` workflow:

1. **Triggers**:
   - Pushes to master or tags (v*)
   - PRs affecting docs/** or hack/doc-site/**

2. **Build Steps**:
   ```bash
   # Clone theme
   git clone --branch v9 https://github.com/alex-shpak/hugo-book themes/hugo-book

   # Generate dynamic config
   LATEST_RELEASE=$(git tag --list --sort -version:refname -i 'v*'|head -1)
   REQUIRED_GO_VERSION=$(grep "^go\s" go.mod|cut -d" " -f2)

   # Create goswagger.yaml with version info
   printf "params:\n  goswagger:\n    goVersion: '%s'\n    latestRelease: '%s'\n" \
     "${REQUIRED_GO_VERSION}" "${LATEST_RELEASE}" > goswagger.yaml

   # Build with Hugo
   hugo --config hugo.yaml,goswagger.yaml \
        --buildDrafts \
        --cleanDestinationDir \
        --minify \
        --printPathWarnings \
        --ignoreCache \
        --noBuildLock \
        --logLevel info \
        --source hack/doc-site/hugo
   ```

3. **Deployment**:
   - Uploads to GitHub Pages
   - Only on non-PR events

### Hugo Version
- **Pinned to**: v0.123.8 (extended version)
- Pinned to avoid breaking changes from Hugo updates

## Directory Structure

```
.
├── docs/                          # Markdown content source
│   ├── _index.md                  # Homepage
│   ├── about.md
│   ├── features.md
│   ├── contributing/
│   ├── faq/
│   ├── generate/
│   ├── install/
│   ├── reference/
│   ├── tutorial/
│   ├── usage/
│   └── presentations/             # Mounted as static HTML
│
└── hack/doc-site/hugo/            # Hugo site
    ├── hugo.yaml                  # Main config
    ├── goswagger.yaml             # Dynamic config (generated)
    ├── gendoc.sh                  # Local dev script
    ├── content/                   # Empty (docs mounted here)
    ├── layouts/                   # Custom layouts
    │   ├── shortcodes/
    │   ├── partials/
    │   └── _default/_markup/
    └── themes/
        ├── hugo-book/             # Cloned at build time
        ├── go-swagger-assets/     # Custom SCSS
        └── go-swagger-static/     # Logos, favicons
```

## Common Tasks

### Adding New Documentation
1. Create/edit markdown files in `./docs/`
2. Add front matter:
   ```yaml
   ---
   title: Page Title
   date: 2023-01-01T01:01:01-08:00
   draft: true
   weight: 10
   description: Brief description
   ---
   ```
3. For new sections, add `_index.md`
4. Use Hugo shortcodes: `{{< hint >}}`, `{{< forkme >}}`, etc.

### Testing Locally
```bash
cd hack/doc-site/hugo
./gendoc.sh
# Visit http://localhost:1313/go-swagger/
```

### Updating Theme
```bash
cd hack/doc-site/hugo/themes
rm -rf hugo-book
git clone --branch v9 https://github.com/alex-shpak/hugo-book
```

### Customizing Styles
Edit `hack/doc-site/hugo/themes/go-swagger-assets/_custom.scss`

### Adding Shortcodes
Create new files in `hack/doc-site/hugo/layouts/shortcodes/`

## Key Features

1. **Dark Theme**: Default dark mode with hugo-book theme
2. **Search**: Enabled with FlexSearch
3. **Edit Links**: "Edit this page" links to GitHub
4. **Git Info**: Last modified dates from git history
5. **Navigation**: Auto-generated from file structure and weights
6. **Presentations**: Raw HTML slides served as static content
7. **Custom Menus**: Before/after menu items for external links

## Troubleshooting

### Build Fails
- Check Hugo version (must be extended)
- Ensure theme is cloned
- Verify both config files exist

### Content Not Showing
- Check front matter `draft: true` status
- Verify file is in `docs/` directory
- Check mount configuration in `hugo.yaml`

### Styles Not Applied
- Ensure extended Hugo version
- Check SCSS files in go-swagger-assets
- Clear Hugo cache: `--ignoreCache`

## Notes

- Documentation is marked as `draft: true` but built with `--buildDrafts`
- Presentations excluded from content mount, served as static HTML
- Menu items weighted for ordering (lower weight = higher in menu)
- Git tags used to automatically update version info in docs
