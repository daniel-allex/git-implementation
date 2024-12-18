Git implementation in Go
- Supports reading / writing Git objects (blobs, trees, commits)
- Supports Git plumbing commands (internal commands, used by user-facing Git commands)
- Compressed versions of files stored by sha1 in git objects folder, accessible by relevant commit
