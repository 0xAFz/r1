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
- Create a vm resource on germany data center
```bash
kumo iaas create --region eu-west1-a --name kumo --flavor-id eco-1-1-0 --image-id "514508bd-0a60-4c88-ae72-3e7b7dcc3968" --network-ids "30a8d5e8-4752-4974-bccc-9e49f5ccc506" --security-group-id "71cf34ab-f0a7-4663-ba98-a2db7d0a1972" --key-name kumo --ssh-key --disk-size 25 --count 1
```
- Create a vm resource on tehran simin data center
```bash
kumo iaas create --region ir-thr-si1 --name kumo --flavor-id eco-1-1-0 --image-id "fb7b732b-5d1f-43d9-9377-8418d7ad303f" --key-name kumo --ssh-key --disk-size 25 --count 1 --network-ids "bab96191-dad5-46bb-96fd-3e29086aa504" --security-group-id "3d2df2b7-2ed1-4998-a87d-974b0ad4bd4e"
```
- Get resource details
```bash
kumo iaas status
```
- Destroy all resources
```bash
kumo iaas destroy
```
- Show local state `.state.json`
```bash
kumo state
```
```json
{
  "7dc6a41b-03fb-482b": {
    "status": "ACTIVE",
    "ip": [
      "194.5.193.0"
    ],
    "region": "eu-west1-a"
  }
}
```
## Todo
- [x] Add show state
- [x] Destroy all resources
- [x] Get resources info with flags
- [x] Support various region resource creation
- [ ] Support other providers
- [ ] Support other resources (CDN, Storage, Container, ...)

