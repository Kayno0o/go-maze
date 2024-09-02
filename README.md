# go-maze

A maze generator made in go.

## TODO

- [ ] CLI `<command> <algo:-recursive-backtracking>` `-w|--width=<number:-10>` `-h|--height=<number:-w>` `-bg|--background=<hex:-#fff>` `-wc|--wall-color=<hex:-#000>` `-sc|--start-color=<hex:-00c8ff>` `-ec|--end-color=<hex:-ffc800>` `-s|--block-size=<number:-25>` `-p|--padding=<number:-3>`

### Algorithms
- [x] maze [recursive-backtracking](https://weblog.jamisbuck.org/2010/12/27/maze-generation-recursive-backtracking)
- [x] maze [prim's](https://weblog.jamisbuck.org/2011/1/10/maze-generation-prim-s-algorithm)
- [ ] maze [kruskal](https://weblog.jamisbuck.org/2011/1/3/maze-generation-kruskal-s-algorithm.html)
- [ ] maze [aldous-broder](https://weblog.jamisbuck.org/2011/1/17/maze-generation-aldous-broder-algorithm.html)
- [ ] maze [hunt-and-kill](https://weblog.jamisbuck.org/2011/1/24/maze-generation-hunt-and-kill-algorithm.html)
- [ ] maze [sidewinder](https://weblog.jamisbuck.org/2011/2/3/maze-generation-sidewinder-algorithm.html)

### Improvements
- [ ] only right(2⁰) / bottom(2¹) walls on each cells (bitwise operations using 1 uint8, or use 2 bool)
