from typing import List, Sequence, Tuple, Callable, Generator
from classes import Piece, PieceWithPositions, print_game_board, GameBoard


def get_initial_game_board() -> GameBoard:
    return GameBoard(game_board=tuple(
        tuple(Piece(position=(row, column), team="R", king=False,) if (row + column) % 2 == 0 and row < 3 else
              Piece(position=(row, column), team="G", king=False) if (row + column) % 2 == 0 and row > 4 else
              # Testen
             # Piece(position=(row, column), team="G", king=False) if row == 3 and column == 3 else
              #Piece(position=(row, column), team="G", king=True) if row == 5 and column == 1 else

              Piece(position=(row, column), team=" ", king=False) for column in range(8))
        for row in range(8)), currPlayer="G")


def get_possible_move_vectors(game_board: GameBoard, piece: Piece, is_piece_there: bool) -> List[Tuple[int, int]]:
    if is_piece_there:
        vectors_which_piece_can_move = [
            (1, -1), (1, 1)] if piece.team == "R" else [(-1, -1), (-1, 1)]
        if piece.king == True:
            vectors_which_piece_can_move = [(1, -1), (1, 1), (-1, -1), (-1, 1)]
        return vectors_which_piece_can_move
    else:
        return []


def get_possible_moves_without_smash(game_board: GameBoard, piece: Piece,
                                     vectors_which_piece_can_move: List[Tuple[int, int]],
                                     checker: Callable, is_piece_there: Callable) -> List[Tuple[int, int]]:
    row, column = piece.position
    unfiltered_possible_positions = [(row + row_vector, column + column_vector)
                                     for row_vector, column_vector in vectors_which_piece_can_move]
    positions_within_bound = filter(lambda position: checker(
        (position[0], position[1])), unfiltered_possible_positions)

    bounded_positions_without_piece_on_it = filter(
        lambda position: not is_piece_there(
            game_board, game_board.game_board[position[0]][position[1]]),
        positions_within_bound
    )
    return list(bounded_positions_without_piece_on_it)


def get_possible_smash_moves(game_board: GameBoard, piece: Piece,
                             checker: Callable,
                             vectors_which_piece_can_move: List[Tuple[int, int]],
                             is_piece_there: Callable) -> Tuple[Tuple[int, int], ...]:
    row, column = piece.position

    def opposite_pieces_on_position(own_piece: Piece, newPositionPiece: Piece):
        return own_piece.team != " " and newPositionPiece.team != " " and own_piece.team != newPositionPiece.team

    # Erstmal davon ausgehen dass ich überall schmeißen kann
    possible_smash_positions = ((row + 2*row_vector, column + 2*column_vector)
                                for row_vector, column_vector in vectors_which_piece_can_move
                                if checker((row + row_vector, column+column_vector)) and checker((row + 2*row_vector, column + 2*column_vector))
                                and opposite_pieces_on_position(own_piece=game_board.game_board[row][column], newPositionPiece=game_board.game_board[row + row_vector][column + column_vector])
                                and not is_piece_there(game_board=game_board, piece=game_board.game_board[row+2*row_vector][column+2*column_vector]))

    return tuple(possible_smash_positions)


def check_if_position_out_of_bound(
    position: Tuple[int, int]): return 0 <= position[0] < 8 and 0 <= position[1] < 8


def is_piece_there_on_row_column(game_board: GameBoard, piece: Piece): return check_if_position_out_of_bound(
    piece.position) and (game_board.game_board[piece.position[0]][piece.position[1]].team != " ")




def get_all_possible_moves_for_piece(game_board: GameBoard, piece: Piece, possible_move_function: Callable) \
        -> Piece:

    vectors_which_piece_can_move = get_possible_move_vectors(
        game_board=game_board, piece=piece, is_piece_there=is_piece_there_on_row_column(game_board=game_board, piece=piece))
    if vectors_which_piece_can_move:
        possible_postions = possible_move_function(
            game_board=game_board, piece=piece,
            vectors_which_piece_can_move=vectors_which_piece_can_move,
            checker=check_if_position_out_of_bound, is_piece_there=is_piece_there_on_row_column)
        if possible_postions:
            return PieceWithPositions(**piece.model_dump(), possible_positions=possible_postions)
        else:
            return piece
    else:
        return piece


def get_all_possible_moves_for_piece_gui(game_board: GameBoard, piece) -> List[Tuple[int, int]]:
    all_moves = get_all_possible_moves_for_team(
        game_board=game_board, team=piece.team)
    potential_move = next(filter(lambda piece_iter: isinstance(piece_iter, PieceWithPositions)\
                        and piece_iter.position == piece.position, all_moves), None)
    if all_moves and potential_move:
        return potential_move.possible_positions
    else:
        return []

def get_all_possible_moves_for_team(game_board: GameBoard, team: str) -> List[Piece]:
    
    team_pieces = list(filter(lambda piece: piece.team ==
                              team, (piece for row in game_board.game_board for piece in row)))
    pieces_with_possible_smash = map(lambda piece: get_all_possible_moves_for_piece(
        game_board=game_board, piece=piece,
        possible_move_function=get_possible_smash_moves), team_pieces.copy())

    pieces_with_possible_smash_moves = list(filter(lambda piece: isinstance(
        piece, PieceWithPositions), pieces_with_possible_smash))

    if not pieces_with_possible_smash_moves:
        pieces_with_possible_moves_without_smash = map(lambda piece: get_all_possible_moves_for_piece(
            game_board=game_board, piece=piece,
            possible_move_function=get_possible_moves_without_smash), team_pieces)
        return list(filter(lambda piece: isinstance(piece, PieceWithPositions), pieces_with_possible_moves_without_smash))
    else:
        return pieces_with_possible_smash_moves


if __name__ == "__main__":
    game_board = get_initial_game_board()
    print_game_board(game_board)
    #print(get_all_possible_moves_for_team(game_board, "G"))
    print(get_all_possible_moves_for_piece_gui(game_board, game_board.game_board[2][2]))