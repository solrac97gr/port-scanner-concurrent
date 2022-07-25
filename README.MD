# Usage
We need to pass the next parameters:
- Host: Represent the URL of the website we want scan[Default:scanme.nmap.org]
- From: From what port we want start the scan [Default:0]
- Limit: Until what port we want Scan. [Default:65535]

```bash
go run main.go --host="example.com" --from=100 --limit=500
```