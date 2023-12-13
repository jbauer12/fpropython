from typing import List, Sequence, Tuple, Callable, Generator
from classes import Piece, PieceWithPositions, GameBoard
from functools import partial


def get_initial_game_board() -> GameBoard:

    def get_Piece_with_right_Team(row: int, column: int):
        team = "R" if (row + column) % 2 == 0 and row < 3 else \
            "G" if (row + column) % 2 == 0 and row > 4 else " "
        return Piece(position=(row, column), team=team, king=False)

    return GameBoard(game_board=tuple(
        tuple(get_Piece_with_right_Team(row, column) for column in range(8))
        for row in range(8)), currPlayer="G")


def get_possible_move_vectors(game_board: GameBoard, piece: Piece, is_piece_there: bool) -> Tuple[Tuple[int, int], ...]:
    if is_piece_there:
        vectors_which_piece_can_move = (
            (1, -1), (1, 1)) if piece.team == "R" else ((-1, -1), (-1, 1))
        if piece.king == True:
            vectors_which_piece_can_move = ((1, -1), (1, 1), (-1, -1), (-1, 1))
        return vectors_which_piece_can_move
    return tuple()


def filter_without_smash(game_board: GameBoard, position: Tuple[int, int]):
    in_bound = check_if_position_out_of_bound(position)
    return in_bound and not is_piece_there_on_row_column(
        game_board=game_board, piece=game_board.game_board[position[0]][position[1]]
    )


def filter_with_smash(game_board: GameBoard, piece: Piece, vector: Tuple[int, int]):
    def opposite_pieces_on_position(own_piece: Piece, newPositionPiece: Piece):
        return own_piece.team != " " and newPositionPiece.team != " " and own_piece.team != newPositionPiece.team
    row, column = piece.position
    row_vector, column_vector = vector

    in_bound_intermediate = check_if_position_out_of_bound(
        (row + row_vector, column + column_vector))
    in_bound_destination = check_if_position_out_of_bound(
        (row + 2*row_vector, column + 2*column_vector))
    if in_bound_intermediate and in_bound_destination:
        opposite = opposite_pieces_on_position(
            piece, game_board.game_board[row + row_vector][column + column_vector])
        without_piece_on_position = not is_piece_there_on_row_column(
            game_board=game_board, piece=game_board.game_board[row+2*row_vector][column+2*column_vector])
        return opposite and without_piece_on_position
    else:
        return False


def get_possible_moves(game_board: GameBoard, piece: Piece,
                       vectors_which_piece_can_move: List[Tuple[int, int]],
                       smash: bool, filter_function: Callable) -> Tuple[Tuple[int, int], ...]:
    row, column = piece.position
    if smash:
        possible_positions = tuple((row + 2*row_vector, column + 2*column_vector)
                                   for row_vector, column_vector in vectors_which_piece_can_move
                                   if filter_function(game_board=game_board, piece=piece, vector=(row_vector, column_vector)))
    else:
        possible_positions = tuple((row + row_vector, column + column_vector)
                                   for row_vector, column_vector in vectors_which_piece_can_move if filter_function(game_board=game_board, position=(row + row_vector, column + column_vector)))
    return possible_positions


def get_possible_smash_moves(game_board: GameBoard, piece: Piece,
                             vectors_which_piece_can_move: List[Tuple[int, int]]
                             ) -> Tuple[Tuple[int, int], ...]:
    return get_possible_moves(game_board=game_board, piece=piece, vectors_which_piece_can_move=vectors_which_piece_can_move, smash=True, filter_function=filter_with_smash)


def get_possible_moves_without_smash(game_board: GameBoard, piece: Piece,
                                     vectors_which_piece_can_move: List[Tuple[int, int]]):
    return get_possible_moves(game_board=game_board, piece=piece, vectors_which_piece_can_move=vectors_which_piece_can_move, smash=False, filter_function=filter_without_smash)


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
            vectors_which_piece_can_move=vectors_which_piece_can_move)
        if possible_postions:
            return PieceWithPositions(**piece.model_dump(), possible_positions=possible_postions)
        else:
            return piece
    else:
        return piece


def get_all_possible_moves_for_piece_gui(game_board: GameBoard, piece) -> List[Tuple[int, int]]:
    all_moves = get_all_possible_moves_for_team(
        game_board=game_board, team=piece.team)
    potential_move = next(filter(lambda piece_iter: isinstance(piece_iter, PieceWithPositions)
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
