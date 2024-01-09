from logic.classes import GameBoard, Piece
from logic.minimax import terminal

def get_testing_game_board(smash:bool=True, only_red:bool=False, only_green:bool = False) -> GameBoard:

    def get_Piece_with_right_Team(row: int, column: int):
        team = "R" if (row + column) % 2 == 0 and row < 3 and not only_green else \
            "G" if (row + column) % 2 == 0 and row > 4 and not only_red else \
                "G" if row ==3 and column == 3 and smash and not only_red else " "
        return Piece(position=(row, column), team=team, king=False)

    return GameBoard(game_board=tuple(
        tuple(get_Piece_with_right_Team(row, column) for column in range(8))
        for row in range(8)), currPlayer="G")

def test_terminal():
    game_board = get_testing_game_board(only_red=True)
    terminal_state = terminal(game_board)
    assert terminal_state == True
    game_board = get_testing_game_board(only_green=True)
    terminal_state = terminal(game_board)
    assert terminal_state == True
    game_board = get_testing_game_board()
    terminal_state = terminal(game_board)
    assert terminal_state == False  


def test_opposite():
    from logic.minimax import opposite
    assert opposite("R", True) == "R"
    assert opposite("G", True) == "G"
    assert opposite("R", False) == "G"
    assert opposite("G", False) == "R"

def test_value_from():
    from logic.minimax import value_from
    game_board = get_testing_game_board(only_red=True)
    assert value_from(game_board) != 0
    game_board = get_testing_game_board(only_green=True)
    assert value_from(game_board) != 0