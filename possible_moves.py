from typing import List, Tuple, Callable, Generator
from classes import Piece, TYPE_GAMEBOARD, print_game_board


def get_initial_game_board() -> TYPE_GAMEBOARD:
    return tuple(
        tuple(Piece(position=(row, column), team="R", king=False) if (row + column) % 2 == 0 and row < 3 else
              Piece(position=(row, column), team="G", king=False) if (row + column) % 2 == 0 and row > 4 else
              # Testen Piece(position=(row, column), team="R", king=True) if row ==4 and column == 4  else
              Piece(position=(row, column), team=" ", king=False) for column in range(8))
        for row in range(8))



def get_possible_move_vectors(game_board: TYPE_GAMEBOARD, piece: Piece, is_piece_there: Callable) -> List[Tuple[int, int]]:
    if is_piece_there(game_board, piece):
        vectors_which_piece_can_move = [
            (1, -1), (1, 1)] if piece.team == "R" else [(-1, -1), (-1, 1)]
        if piece.king:
            vectors_which_piece_can_move = [(1, -1), (1, 1), (-1, -1), (-1, 1)]
        return vectors_which_piece_can_move
    else:
        return []


def get_possible_moves_without_smash(game_board: TYPE_GAMEBOARD, piece: Piece,
                                     vectors_which_piece_can_move: List[Tuple[int, int]],
                                     checker: Callable, piece_there: Callable) -> List[Tuple[int, int]]:
    row, column = piece.position
    unfiltered_possible_positions = [(row + row_vector, column + column_vector)
                                     for row_vector, column_vector in vectors_which_piece_can_move]
    positions_within_bound = filter(lambda position: checker((position[0],position[1])), unfiltered_possible_positions)

    bounded_positions_without_piece_on_it = filter(
        lambda position: not piece_there(
            game_board, game_board[position[0]][position[1]]),
        positions_within_bound
    )
    return list(bounded_positions_without_piece_on_it)


def get_possible_smash_moves(game_board: TYPE_GAMEBOARD, piece: Piece,
                             checker: Callable,
                             vectors_which_piece_can_move: List[Tuple[int, int]],
                             is_piece_there: Callable) -> List[Tuple[int, int]]:
    row, column = piece.position

    def opposite_pieces_on_position(own_piece: Piece, newPositionPiece: Piece):
        return own_piece.team != " " and newPositionPiece.team != " " and own_piece.team != newPositionPiece.team

    # Erstmal davon ausgehen dass ich überall schmeißen kann
    possible_smash_positions =  ((row + 2*row_vector, column + 2*column_vector)
                                           for row_vector, column_vector in vectors_which_piece_can_move
                                           if checker((row +row_vector,column+column_vector)) and checker((row +2*row_vector,column + 2*column_vector))\
                                              and opposite_pieces_on_position(own_piece=game_board[row][column], newPositionPiece=game_board[row + row_vector][column + column_vector])\
                                                and not is_piece_there(game_board=game_board, piece=game_board[row+2*row_vector][column+2*column_vector]))

    return list(possible_smash_positions)

def check_if_position_out_of_bound(
    position: Tuple[int,int]): return 0 <= position[0] < 8 and 0 <= position[1] < 8

def is_piece_there_on_row_column(game_board: TYPE_GAMEBOARD, piece: Piece): return check_if_position_out_of_bound(
    piece.position) and (game_board[piece.position[0]][piece.position[1]].team != " ")

def get_all_possible_moves_for_gui(game_board: TYPE_GAMEBOARD, piece: Piece) -> List[Tuple[int, int]]:


    possible_moves: List[Tuple[int, int]] = []
    vectors_which_piece_can_move = get_possible_move_vectors(
        game_board=game_board, piece=piece, is_piece_there=is_piece_there_on_row_column)

    if vectors_which_piece_can_move:

        bounded_positions_without_piece_on_it = get_possible_moves_without_smash(
            game_board=game_board, piece=piece,
            vectors_which_piece_can_move=vectors_which_piece_can_move,
            checker=check_if_position_out_of_bound, piece_there=is_piece_there_on_row_column)

        possible_smash_positions = get_possible_smash_moves(
            game_board=game_board, piece=piece,
            vectors_which_piece_can_move=vectors_which_piece_can_move,
            checker=check_if_position_out_of_bound,
            is_piece_there=is_piece_there_on_row_column)
    
        possible_moves = list(possible_smash_positions) + \
            list(bounded_positions_without_piece_on_it)
    return possible_moves

def get_all_possible_moves_for_team(game_board: TYPE_GAMEBOARD, team:str):
    vectors_which_piece_can_move = get_possible_move_vectors(
        game_board=game_board, piece=piece, is_piece_there=is_piece_there_on_row_column)
    filtered_pieces = filter(lambda piece: piece.team == team, (piece for row in game_board for piece in row))
    filtered_pieces = map(lambda piece: get_possible_smash_moves(game_board, piece, vectors_which_piece_can_move=get_possible_move_vectors), filtered_pieces)

    return print(list(filtered_pieces))

if __name__ == "__main__":
    game_board = get_initial_game_board()
    print_game_board(game_board)
    get_all_possible_moves_for_team(game_board, "R")    
