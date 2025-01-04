# Kumo

Cli tool for working with Cloud providers

Supported providers:
- ArvanCloud

## Build
```bash
make
```
## Usage
0. Setup `.env`
```bash
cp .env.example .env
<your-editor> .env
```
1. Create a iaas resource
```bash
kumo iaas create
```
2. Get resource details
```bash
kumo iaas status
```
3. Destroy all resources
```bash
kumo iaas remove
```
## Todo
- [ ] Add show state
- [ ] Support other providers
- [ ] Support other resources

