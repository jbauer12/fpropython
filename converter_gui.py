from typing import List, Tuple
from gui import Node, Piece, WHITE, BLACK, WIDTH, ROWS

def convert_array_to_printable_grid(array: List[List[str]]) -> List[List[Node]]:
    grid: List[List[Node]] = []
    for i in range(8):
        grid.append([])
        for j in range(8):
            node = Node(j,i, WIDTH / ROWS)
            node.colour = BLACK if abs(i - j) % 2 == 0 else WHITE
            grid[i].append(node)
            if array[i][j] == 'R':
                grid[i][j].piece = Piece('R')
            elif array[i][j] == 'G':
                grid[i][j].piece = Piece('G')
    return grid

def convert_printable_grid_to_array(grid: List[List[Node]]) -> List[List[str]]:
    array: List[List[str]] = []
    for i in range(8):
        array.append([])
        for j in range(8):
            if grid[i][j].piece:
                array[i].append(grid[i][j].piece.team)
            else:
                array[i].append(" ")
    return array