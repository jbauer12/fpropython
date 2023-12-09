from typing import List, Tuple, Callable
from possible_moves import get_all_possible_moves, get_initial_game_board, TYPE_GAMEBOARD, red_team


def modify_game_board(game_board):
    def modify_tuple_element(element):
        return element.upper()

    def modify_game_board_row(row):
        return tuple(map(modify_tuple_element, row))
    
    return tuple(map(modify_game_board_row, game_board))


# Example usage:
initial_game_board: TYPE_GAMEBOARD = (
    ("R", " ", "R", " ", "R", " ", "R", " "),
    (" ", "R", " ", "R", " ", "R", " ", "R"),
    ("R", " ", "R", " ", "R", " ", "R", " "),
    (" ", " ", " ", "G", " ", "", " ", " "),
    (" ", " ", " ", " ", " ", " ", " ", " "),
    (" ", "G", " ", "G", " ", "G", " ", "G"),
    ("G", " ", "G", " ", "G", " ", "G", " "),
    (" ", "G", " ", "G", " ", "G", " ", "G")
)



def make_move(game_board: TYPE_GAMEBOARD, newPosition: Tuple[int, int], oldPosition: Tuple[int, int]):
    row,column = newPosition
    oldrow, oldcolumn = oldPosition
    moves = get_all_possible_moves(
        game_board=game_board, nodePosition=(oldrow, oldcolumn ))
    red = red_team(game_board, oldrow, oldcolumn)
    char_to_write = "R" if red else "G"

    if (row, column) in moves:
        new_game_board = tuple(
            tuple(char_to_write if i == row and j == column 
                  else " " if i ==oldrow and j == oldcolumn else col for j, col in enumerate(r))
            for i, r in enumerate(game_board))
        
        return new_game_board
    else:
        return game_board



