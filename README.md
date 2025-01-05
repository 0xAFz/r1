# Kumo

Cli tool for working with Cloud providers

## Supported providers:
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
1. Create a vm resource on germany datacenter
```bash
kumo iaas --region eu-west1-a create --name kumo --flavor-id eco-1-1-0
 --image-id "514508bd-0a60-4c88-ae72-3e7b7dcc3968" --network-ids "30a8d5e8-4752-4974
-bccc-9e49f5ccc506" --security-group-id "71cf34ab-f0a7-4663-ba98-a2db7d0a1972" --key
-name kumo --ssh-key --disk-size 25 --count 1
```
2. Get resource details
```bash
kumo iaas status
```
3. Destroy all resources
```bash
kumo iaas destroy
```
4. Show local state `.state.json`
```bash
kumo state
```
```json
{
  "2bac392d": {
    "status": "ACTIVE",
    "ip": [
      "194.5.193.0"
    ]
  },
  "588f57f7": {
    "status": "BUILD",
    "ip": []
  }
}
```
## Todo
- [x] Add show state
- [x] Destroy all resources
- [x] Get resources info with flags
- [ ] Support other providers
- [ ] Support other resources (CDN, Storage, Container, ...)

