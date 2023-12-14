import pytest
from logic.classes import Piece as logicPiece
from logic.classes import GameBoard
from logic.gui import convert_array_to_printable_grid, ROWS, convert_printable_grid_to_array

def get_testing_game_board() -> GameBoard:

    def get_Piece_with_right_Team(row: int, column: int):
        team = "R" if (row + column) % 2 == 0 and row < 3 else \
            "G" if (row + column) % 2 == 0 and row > 4 else " "
        return logicPiece(position=(row, column), team=team, king=False)

    return GameBoard(game_board=tuple(
        tuple(get_Piece_with_right_Team(row, column) for column in range(8))
        for row in range(8)), currPlayer="G")


def test_convert_array_to_printable_grid():
    board = get_testing_game_board()
    result = convert_array_to_printable_grid(board)

    board_pieces = sum(piece.team =="R" or piece.team =="G" for row in board.game_board for piece in row )
    result_pieces = sum(node.piece.team =="R" or node.piece.team =="G" for row in result for node in row if node.piece )
    assert board_pieces == result_pieces
    assert len(board.game_board) == ROWS == len(result)
    for row in board.game_board:
        assert len(row) == ROWS
    for row in result:
        assert len(row) == ROWS
    board_king_pieces = sum(piece.king for row in board.game_board for piece in row )
    result_king_pieces = sum(node.piece.type =="KING" for row in result for node in row if node.piece )
    assert board_king_pieces == result_king_pieces

def test_convert_printable_grid_to_array():
    currPlayer="R"
    board = get_testing_game_board()
    result = convert_array_to_printable_grid(board)
    board = convert_printable_grid_to_array(result, currPlayer=currPlayer)
    board_pieces = sum(piece.team =="R" or piece.team =="G" for row in board.game_board for piece in row )
    result_pieces = sum(node.piece.team =="R" or node.piece.team =="G" for row in result for node in row if node.piece )
    assert board_pieces == result_pieces
    assert len(board.game_board) == ROWS == len(result)
    for row in board.game_board:
        assert len(row) == ROWS
    board_king_pieces = sum(piece.king for row in board.game_board for piece in row )
    result_king_pieces = sum(node.piece.type =="KING" for row in result for node in row if node.piece )
    assert board_king_pieces == result_king_pieces