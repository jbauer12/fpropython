from logic.rules import opposite, make_move, is_opposite_piece_smashed, is_piece_checker_after_move
from logic.classes import Piece, GameBoard


def get_testing_game_board(smash:bool=True) -> GameBoard:

    def get_Piece_with_right_Team(row: int, column: int):
        team = "R" if (row + column) % 2 == 0 and row < 3 else \
            "G" if (row + column) % 2 == 0 and row > 4 else \
                "G" if row ==3 and column == 3 and smash else " "
        return Piece(position=(row, column), team=team, king=False)

    return GameBoard(game_board=tuple(
        tuple(get_Piece_with_right_Team(row, column) for column in range(8))
        for row in range(8)), currPlayer="G")

def test_opposite():
    red =Piece(position=(0,0), team="R", king=False)
    green = Piece(position=(0,0), team="G", king=False)
    assert opposite(red) == "G"
    assert opposite(green) == "R"

def test_make_move():
    game_board = get_testing_game_board()
    #Smash
    new_game_board=make_move(game_board, (4,4), game_board.game_board[2][2])
    assert new_game_board.game_board[2][2].team ==" " and game_board.game_board[2][2].team == "R"
    assert new_game_board.game_board[4][4].team =="R" and game_board.game_board[4][4].team ==" "
    assert game_board.game_board[3][3].team == "G" and new_game_board.game_board[3][3].team ==" "
    new_game_board=make_move(game_board, (3,1), game_board.game_board[2][2])
    assert new_game_board.game_board[3][1].team != "R" 
    assert new_game_board.game_board[2][2].team == "R" 
    #Normal
    game_board = get_testing_game_board(smash=False)
    new_game_board=make_move(game_board, (3,1), game_board.game_board[2][2])
    assert new_game_board.game_board[3][1].team =="R" and new_game_board.game_board[2][2].team == " "
    assert game_board.game_board[2][2].team == "R"

def test_is_opposite_piece_smashed():
    board = get_testing_game_board()
    smashed = is_opposite_piece_smashed(board.game_board[2][2], (4,4))
    assert smashed == True
    board = get_testing_game_board(smash=False)
    smashed = is_opposite_piece_smashed(board.game_board[2][2], (3,1))



    
def test_is_piece_checker_after_move():
    piece_red_true = Piece(position=(6,5), team="R", king=False)
    piece_green_true = Piece(position=(1,1), team ="G", king=False)
    checker_red = is_piece_checker_after_move(piece_red_true, (7,5))
    assert checker_red == True
    checker_green = is_piece_checker_after_move(piece_green_true, (0,1))
    assert checker_green == True
    piece_red_false = Piece(position=(1,5), team="R", king=False)
    piece_green_false = Piece(position=(6,1), team ="G", king=False)
    checker_red = is_piece_checker_after_move(piece_red_false, (0,5))
    assert checker_red == False
    checker_green = is_piece_checker_after_move(piece_green_false, (7,1))
    assert checker_green == False
