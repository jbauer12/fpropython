from classes import Piece, print_game_board, GameBoard
from possible_moves import get_all_possible_moves_for_team, get_initial_game_board


def terminal(game_board: GameBoard) -> bool:
    """Endzustand ja oder nein?"""
    # TODO Keine Züge mehr möglich
    # --> eine Farbe oder
    # zwei Farben --> dann Remis
    number_red_pieces = sum(1 for row in game_board.game_board for piece in row if piece.team == "R") == 0
    number_green_pieces = sum(1 for row in game_board.game_board for piece in row if piece.team == "G") == 0

    return number_red_pieces or number_green_pieces


def value(game_board: GameBoard):
    """Bewertungsfunktion --> Wert des Endzustandes
    Max Player wants to maximize the value --> Value +1
    Min Player wants to minimize the value --> Value -1"""
    pass


def player(game_board: GameBoard):
    """Wer ist am Zug?
    return max or min player"""
    pass


def actions(game_board: GameBoard, piece: Piece):
    """Welche Züge sind möglich?
    return list of possible moves"""
    return get_all_possible_moves_for_team(game_board=game_board, team=piece.team)
    


def result(game_board: GameBoard, action):
    """Ergebnis eines Zuges
    return new game board"""
    pass


def minimax(game_board: GameBoard, depth: int = 10):

    if depth == 0 or terminal(game_board):
        # game over
        return value(game_board)
    elif player(game_board) == "max":
        # max player
        return max(minimax(result(game_board, action), depth=depth-1) for action in actions(game_board))
    elif player(game_board) == "min":
        # min player
        return min(minimax(result(game_board, action), depth=depth-1) for action in actions(game_board))


minimax(get_initial_game_board())