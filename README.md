# PAWNS-TOUR
It is a variation of the Knight's Tour problem (https://en.wikipedia.org/wiki/Knight's_tour).
## Pawn Movement Rules :
- Jumps 3 squares horizontally and vertically.
- Jumps 2 squares diagonally.
- Jumps to square already visited or outside the checkerboard are NOT allowed.

## Build and Run

```sh
$ go vet ./...
$ go run ./...
```

Output :
```sh
----------47126th try----------
01↓ 29↓ 40→ 06→ 28← 41↘ 07↓ 27←
38↘ 16↓ 51→ 37← 15← 52↓ 36← 14←
49↘ 44↓ 62→ 20→ 43← 63↓ 21↓ 42←
02↓ 30↓ 39↑ 05↑ 23↓ 56↓ 08↓ 26↑
59↓ 17↓ 50↑ 32↓ 10↓ 53↓ 35↑ 13↑
48↑ 45→ 61↑ 19↑ 46↙ 64* 22↖ 55↖
03→ 31↗ 58↖ 04↑ 24→ 57← 09↖ 25↑
60↗ 18↗ 47↖ 33→ 11→ 54↗ 34↑ 12↑ 
```
* This output is calculated for 8x8 board size and starting position 0,0 after 47126 tries.