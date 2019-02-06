workflow "Test" {
  on = "push"
  resolves = ["Run tests"]
}

action "Run tests" {
  uses = "docker://golang:1.11-alpine"
  runs = "go test"
}
