This simplified implementation of Git, written in Go, to help understand how Git works under the hood. It supports a subset of Git commands with a focus on clarity and correctness.

---

## ✨ Features

The following Git-like commands are currently implemented:

- `init` – Initialize a new repository
- `add` – Add files to the staging area
- `rm` – Remove files from the staging area and working directory
- `status` – Show the working tree status
- `commit` – Record changes to the repository
- `log` – Show commit logs

---

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/a-git.git
cd a-git
```

Build the binary:

```bash
go build -o a-git .
```

## 🚀 Usage
Commands work similarly to Git.

```bash
./a-git init
echo "Hello, world!" > hello.txt
./a-git add hello.txt
./a-git commit -m "Add hello.txt"
./a-git log
```
