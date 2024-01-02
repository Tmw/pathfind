function emoji() {
    go run cmd/main.go \
        -filename="examples/emoji.txt" \
        -symbolNonWalkable="🔥" \
        -symbolWalkable="⬜" \
        -symbolStart="🟢" \
        -symbolFinish="🏁" \
        -symbolPath="🚗"
}

function small() {
    go run cmd/main.go \
        -filename="examples/small.txt"
}

function help() {
    echo "Missing agrument.\n\nUsage: examples/example.sh <emoji|small>"
}

case "$1" in
    "emoji") emoji ;;
    "small") small ;;

    *) help ;;
esac
