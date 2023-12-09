from typing import List, Tuple, Callable

GREEN_TEAM = ("G", "GK")
RED_TEAM = ("R", "RK")


def get_initial_game_board():
    return (
        ("R", " ", "R", " ", "R", " ", "R", " "),
        (" ", "R", " ", "R", " ", "R", " ", "R"),
        ("R", " ", "R", " ", "R", " ", "R", " "),
        (" ", " ", " ", "G", " ", "", " ", " "),
        (" ", " ", " ", " ", " ", " ", " ", " "),
        (" ", "G", " ", "G", " ", "G", " ", "G"),
        ("G", " ", "G", " ", "G", " ", "G", " "),
        (" ", "G", " ", "G", " ", "G", " ", "G"))


def get_possible_move_vectors(game_board: List[List[str]], row: int, column: int, red_team: Callable, is_piece_there: Callable) -> List[Tuple[int, int]]:
    is_king_type = game_board[row][column] == "RK" or game_board[row][column] == "GK"
    if is_piece_there(game_board, row, column):
        vectors_which_piece_can_move = [
            (1, -1), (1, 1)] if red_team(game_board, row, column) else [(-1, -1), (-1, 1)]
        if is_king_type:
            vectors_which_piece_can_move = [(1, -1), (1, 1), (-1, -1), (-1, 1)]

        return vectors_which_piece_can_move


def get_possible_moves_without_smash(game_board: List[List[str]], nodePosition: Tuple[int, int],
                                     vectors_which_piece_can_move: List[Tuple[int, int]],
                                     checker: Callable, piece_there: Callable) -> List[Tuple[int, int]]:
    row, column = nodePosition
    unfiltered_possible_positions = [(row + row_vector, column + column_vector)
                                     for row_vector, column_vector in vectors_which_piece_can_move]
    positions_within_bound = filter(checker, unfiltered_possible_positions)

    bounded_positions_without_piece_on_it = filter(
        lambda position: not piece_there(
            game_board, position[0], position[1]),
        positions_within_bound
    )
    return bounded_positions_without_piece_on_it


def get_possible_smash_moves(game_board: List[List[str]], nodePosition: Tuple[int, int],
                             checker: Callable, red_team: Callable,
                             vectors_which_piece_can_move: List[Tuple[int, int]],
                             is_piece_there: Callable) -> List[Tuple[int, int]]:
    row, column = nodePosition

    def opposite_piece_on_field(
            team_red, position):
        return team_red and position in GREEN_TEAM or not team_red and position in RED_TEAM
    def opposite_pieces_on_position(position): return opposite_piece_on_field(
        team_red=red_team(game_board, row, column), position=position)

    # Erstmal davon ausgehen dass ich überall schmeißen kann
    unfiltered_possible_positions_smash = ((nodePosition[0] + 2*row_vector, nodePosition[1] + 2*column_vector)
                                           for row_vector, column_vector in vectors_which_piece_can_move
                                           if opposite_pieces_on_position(game_board[nodePosition[0] + row_vector][nodePosition[1] + column_vector]) and
                                           not is_piece_there(game_board=game_board, row=nodePosition[0] + 2*row_vector, column=nodePosition[1] + 2*column_vector))

    possible_smash_positions = filter(
        checker, unfiltered_possible_positions_smash)
    return possible_smash_positions


def get_all_possible_moves(game_board: List[List[str]], nodePosition: Tuple[int, int]) -> List[Tuple[int, int]]:
    def red_team(game_board: List[List[str]], row, column) -> bool:
        return game_board[row][column] in RED_TEAM

    def is_piece_there_row_column(
        game_board, row, column): return game_board[row][column] in RED_TEAM or game_board[row][column] in GREEN_TEAM

    possible_moves: List[Tuple[int, int]] = []
    row, column = nodePosition
    vectors_which_piece_can_move = get_possible_move_vectors(
        game_board=game_board, row=row, column=column, red_team=red_team, is_piece_there=is_piece_there_row_column)

    if vectors_which_piece_can_move:
        def check_if_position_out_of_bound(
            position): return 0 <= position[0] < 8 and 0 <= position[1] < 8

        bounded_positions_without_piece_on_it = get_possible_moves_without_smash(
            game_board=game_board, nodePosition=nodePosition,
            vectors_which_piece_can_move=vectors_which_piece_can_move,
            checker=check_if_position_out_of_bound, piece_there=is_piece_there_row_column)

        possible_smash_positions = get_possible_smash_moves(
            game_board=game_board, nodePosition=nodePosition,
            vectors_which_piece_can_move=vectors_which_piece_can_move,
            checker=check_if_position_out_of_bound,
            red_team=red_team, is_piece_there=is_piece_there_row_column)

        possible_moves = list(possible_smash_positions) + \
            list(bounded_positions_without_piece_on_it)
    return possible_moves


print(get_all_possible_moves(get_initial_game_board(), (2, 2)))
