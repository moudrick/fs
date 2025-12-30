# fs

```
go test ./...
```

```sh
go run ./cmd/demo -mem -root /
go run ./cmd/demo -mem -root /d3
```

```sh
vscode ➜ /workspaces/codelab/fs (main) $ go run ./cmd/demo -mem=false -root="$PWD"
```

```log
Group 1:
  /workspaces/codelab/fs/.git/logs/HEAD
  /workspaces/codelab/fs/.git/logs/refs/heads/main
Group 2:
  /workspaces/codelab/fs/.git/refs/heads/main
  /workspaces/codelab/fs/.git/refs/remotes/origin/main
```

```sh
mkdir -p /tmp/finddup-big
cd /tmp/finddup-big

dd if=/dev/urandom of=big.bin bs=1M count=20 status=none
cp big.bin big-copy.bin

cd -
```

```sh
go run /workspaces/codelab/fs/cmd/demo -mem=false -root="/tmp/finddup-big"
```

```sh
Group 1:
  /tmp/finddup-big/big-copy.bin
  /tmp/finddup-big/big.bin
```