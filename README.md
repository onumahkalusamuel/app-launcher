# applauncher

An app launcher for web apps

## Requirements in order of Preference
`Chrome Browser` > `Microsoft Edge` > `Mozilla Firefox`

## Usage
Copy `launch_url.example` into `launch_url` and update it with the URL you want the launcher to load. It could be localhost or online.

```bash
http://localhost:8089
```

Build the app
```bash
go build -o applauncher.exe main.go
```

Copy `applaucher.exe` to desired location and execute.

