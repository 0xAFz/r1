# Kumo

Cli tool for working with Cloud providers

## Supported providers:
- ArvanCloud

## Build
```bash
make
```
## Usage
At first setup your `.env`
```bash
cp .env.example .env
<your-editor> .env
```

Create desired state file
```bash
touch kumo.json
```

Write resources in this strcuture
```json
[
    {
        "region": "ir-thr-si1",
        "data": {
            "name": "kumo",
            "backup_name": null,
            "count": 1,
            "enable_backup": false,
            "network_ids": [
                "8bcaf216-055b-4bc8-92f1-d18b8285bea5"
            ],
            "flavor_id": "eco-1-1-0",
            "security_groups": [
                {
                    "name": "e699ac81-538d-49a4-80c4-5ebe3b79b22d"
                }
            ],
            "ssh_key": true,
            "key_name": "kumo",
            "disk_size": 25,
            "init_script": "",
            "ha_enabled": true,
            "server_volumes": [],
            "enable_ipv4": true,
            "enable_ipv6": false,
            "image_id": "fb7b732b-5d1f-43d9-9377-8418d7ad303f"
        }
    }
]
```

Apply resources
```bash
kumo iaas apply
```

Show current state
```bash
kumo state
```

Update current state
```bash
kumo iaas status
```

Delete All resources
```bash
kumo iaas destroy
```

## Todo
- [x] Support various region resource creation
- [x] Read desired state from json instead of CLI args
- [ ] Support other resources (CDN, Storage, Container, ...)
