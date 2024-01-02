# Pathfind

Currently only offers a generic astar algorithm.

## CLI Usage

Reading from file:

```console
go run cmd/main.go -filename examples/small.txt
```

Reading from `stdin`:

```console
cat examples/small.txt | go run cmd/main.go
```

Overriding symbols

```console
go run cmd/main.go \
    -filename="examples/emoji.txt" \
    -symbolNonWalkable="🔥" \
    -symbolWalkable="⬜" \
    -symbolStart="🟢" \
    -symbolFinish="🏁" \
    -symbolPath="🚗"
```

See examples:

```console
examples/example.sh <small|emoji>
```

## Library usage

T.B.D

## License

[MIT](./LICENSE)
