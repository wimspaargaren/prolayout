# ProLayout

Pro(ject) layout is a static analysis tool that allow you to lint the project
structure of your go project.

## Development

Either use `make` or `task`:

### Task

```zsh
TASK_X_REMOTE_TASKFILES=1 task test-all --yes
```

## Why

Since Go does not enforce any real project structure, we wanted to have a
static analysis tool, to help us to ensure projects are structured in a similar
fashion.

## Example configuration file

```YAML
module: "github.com/wimspaargaren/prolayout"
root:
  - name: "cmd"
    dirs:
      - name: ".*"
        files:
          - "main.go"
  - name: "internal"
  - name: "pkg"
  - name: "tests"
    files:
      - name: ".*_test.go"
```
