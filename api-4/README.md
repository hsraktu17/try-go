## ⚙️ Setup: CompileDaemon for Live Reloading

### 1. Install CompileDaemon

```bash
go install github.com/githubnemo/CompileDaemon@latest
```

Make sure `$GOPATH/bin` is added to your system `PATH`.

---

### 2. Run with Auto Reload (Windows)

```bash
CompileDaemon --build="go build -o server.exe main.go" --command=".\\server.exe"
```

- ✅ Automatically rebuilds and restarts your app on any `.go` file change
- 🪄 Works just like `nodemon` in Node.js

