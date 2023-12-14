from logic.possible_moves import get_all_possible_moves_for_team, get_possible_moves_without_smash, get_possible_smash_moves, filter_with_smash, check_if_position_out_of_bound, is_piece_there_on_row_column, get_possible_move_vectors
from test_rules import get_testing_game_board
from logic.classes import Piece, PieceWithPositions


def test_check_if_position_out_of_bound():
    in_bound = check_if_position_out_of_bound((8, 8))
    assert in_bound == False
    in_bound = check_if_position_out_of_bound((7, 7))
    assert in_bound == True
    in_bound = check_if_position_out_of_bound((0, 0))
    assert in_bound == True


def test_get_possible_move_vectors():
    king_piece = Piece(position=(2, 2), team="R", king=True)
    vectors = get_possible_move_vectors(king_piece, True)
    assert vectors == ((1, -1), (1, 1), (-1, -1), (-1, 1))
    normal_piece_red = Piece(position=(2, 2), team="R", king=False)
    vectors = get_possible_move_vectors(normal_piece_red, True)
    assert vectors == ((1, -1), (1, 1))
    normal_piece_green = Piece(position=(2, 2), team="G", king=False)
    vectors = get_possible_move_vectors(normal_piece_green, True)
    assert vectors == ((-1, -1), (-1, 1))


def test_filter_with_smash():
    game_board = get_testing_game_board(smash=True)
    print(game_board)
    piece = Piece(position=(2, 2), team="R", king=False)
    smash = filter_with_smash(game_board=game_board,
                              piece=piece, vector=(1, 1))
    assert smash == True
    smash = filter_with_smash(game_board=game_board,
                              piece=piece, vector=(1, -1))
    assert smash == False


def test_possible_smash_moves():
    game_board = get_testing_game_board(smash=True)
    piece = Piece(position=(2, 2), team="R", king=False)
    vectors = get_possible_move_vectors(piece, True)
    moves = get_possible_smash_moves(
        game_board=game_board, piece=piece, vectors_which_piece_can_move=vectors)
    assert moves == ((4, 4),)
    game_board = get_testing_game_board(smash=False)
    moves = get_possible_smash_moves(
        game_board=game_board, piece=piece, vectors_which_piece_can_move=vectors)
    assert moves == ()


def test_get_possible_moves_without_smash():
    game_board = get_testing_game_board(smash=True)
    piece = Piece(position=(2, 2), team="R", king=False)
    vectors = get_possible_move_vectors(piece, True)
    moves = get_possible_moves_without_smash(
        game_board=game_board, piece=piece, vectors_which_piece_can_move=vectors)
    assert moves == ((3, 1),)


def test_get_all_possible_moves_for_team():
    game_board = get_testing_game_board(smash=True)
    moves = get_all_possible_moves_for_team(game_board=game_board, team="R")
    piece = PieceWithPositions(
        position=(2, 2), team='R', king=False, possible_positions=[(4, 4)])
    assert piece.model_dump() == moves[0].model_dump()
    game_board = get_testing_game_board(smash=False)
    moves_without_smashes = get_all_possible_moves_for_team(game_board=game_board, team="R")
    moves_without_smashes= [move.model_dump() for move in moves_without_smashes]
    assert piece.model_dump() not in moves_without_smashes
