from typing import List, Tuple, Callable
from possible_moves import  get_all_possible_moves_for_piece_gui
from classes import Piece, GameBoard


def opposite(piece: Piece):
    return "R" if piece.team == "G" else "G"


def make_move(game_board: GameBoard, newPosition: Tuple[int, int], piece:Piece) -> GameBoard:
    row, column = newPosition
    moves = get_all_possible_moves_for_piece_gui(
        game_board=game_board, piece=game_board.game_board[piece.position[0]][piece.position[1]])



    def nextPlayer(smash: bool, piece: Piece):
        team = "R" if piece.team =="R" else "G"
        return team if smash else opposite(
            game_board.game_board[piece.position[0]][piece.position[1]])

    if (row, column) in moves:
        smash = is_opposite_piece_smashed(piece=piece, newPosition=newPosition)
        game_board_after_move = place_piece_on_new_position(game_board=game_board, piece=piece,
                                                            newPosition=newPosition, smash=smash)
        return GameBoard(game_board=game_board_after_move, currPlayer=nextPlayer(smash, piece))
    else:
        return GameBoard(game_board=game_board.game_board, currPlayer=nextPlayer(False, piece))


def place_piece_on_new_position(game_board: GameBoard, piece: Piece,
                                newPosition: Tuple[int, int],  smash: bool ):
    row, column = newPosition
    oldrow, oldcolumn = piece.position

    def delete_piece_team(row_index, col_index, oldrow, oldcolumn):
        because_piece_moved = True if (
            row_index == oldrow and col_index == oldcolumn) else False
        because_piece_smashed = True if smash and (row_index == int(
            (oldrow+row)/2) and col_index == int((oldcolumn+column)/2)) else False
        return because_piece_moved or because_piece_smashed

    return tuple(
        tuple(Piece(position=newPosition, team=piece.team, king=is_piece_checker_after_move(piece, newPosition=newPosition)) if i == row and j == column
              else Piece(position=newPosition, team=" ", king=False) if (delete_piece_team(i, j, oldrow, oldcolumn)) else col for j, col in enumerate(r))
        for i, r in enumerate(game_board.game_board))


def is_opposite_piece_smashed(piece:Piece, newPosition: Tuple[int, int]) -> bool:
    row, column = newPosition
    oldrow, oldcolumn = piece.position
    if abs(row-oldrow) == 2 or abs(column-oldcolumn) == 2:
        return True
    return False


def is_piece_checker_after_move(piece: Piece, newPosition: Tuple[int, int]) -> bool:
    row, column = newPosition
    if (piece.team == "R" and row == 7) or (piece.team == "G" and row == 0) or piece.king:
        return True
    else: return False



def is_game_over(game_board: GameBoard, piece: Piece):
    pass
    #return True if len(get_all_possible_moves(game_board=game_board.game_board, piece=piece )) == 0 else False