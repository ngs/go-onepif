workflow "New workflow" {
  on = "push"
  resolves = ["docker://golang:1.11-alpine"]
}

action "docker://golang:1.11-alpine" {
  uses = "docker://golang:latest"
  runs = "go test"
}
