from typing import Tuple
from classes import Piece, GameBoard
from possible_moves import get_all_possible_moves_for_team
from rules import make_move

def opposite(team:str, smash:bool):
    if smash: return team
    return "R" if team == "G" else "G"

def terminal(game_board: GameBoard) -> bool:
    """Endzustand ja oder nein?"""
    number_red_pieces = all(piece.team =="R" for row in game_board.game_board for piece in row) 
    number_green_pieces = all(piece.team =="G" for row in game_board.game_board for piece in row) 

    return number_red_pieces or number_green_pieces


def value_from(game_board: GameBoard):
    def value_from_piece(piece, positional_weight):
        piece_count_weight = 5
        kinged_piece_weight = 3
        return piece_count_weight + kinged_piece_weight * piece.king + piece.position[0] * positional_weight

    player_score = sum(
        value_from_piece(piece, positional_weight=0.5)
        for row in game_board.game_board
        for piece in row
        if piece.team == game_board.currPlayer
    )

    opponent_score = sum(
        value_from_piece(piece, positional_weight=0.5)
        for row in game_board.game_board
        for piece in row
        if piece.team != " " and piece.team != game_board.currPlayer
    )

    return opponent_score - player_score


def player(game_board: GameBoard):
    """Wer ist am Zug?
    return max or min player"""
    return game_board.currPlayer


def actions(game_board: GameBoard, team: str):
    """Welche Züge sind möglich?
    return list of possible moves"""
    pieces_with_all_possible_moves = get_all_possible_moves_for_team(game_board=game_board, team=team)
    actions = [(piece.position, move) for piece in pieces_with_all_possible_moves for move in piece.possible_positions]
    return actions
    


def result(game_board: GameBoard, action: Tuple[Tuple[int, int], Tuple[int, int]]):
    if not terminal(game_board=game_board):
        oldPosition, newPosition = action
        piece = game_board.game_board[oldPosition[0]][oldPosition[1]]
        return make_move(game_board=game_board, newPosition=newPosition, piece=piece)
    else: return game_board





def minimax(state:GameBoard, depth, player):

    best_score = [None,  -1000000] if player == "R" else [None, +1000000]

    if depth == 0 or terminal(state):
        score = value_from(state)
        return [None, score]

    for action in actions(state, player):
        result1 = result(state, action)
        value = minimax(result1, depth - 1, opposite(player, result1.currPlayer == player))

        if player == "R":
            if value[1] > best_score[1]:
                best_score = [action, value[1]]
        else:
            if value[1] < best_score[1]:
                best_score = [action, value[1]]

    return best_score

def minimax_for_gui(state:GameBoard, player):
    best = minimax(state, 4, player)
    return result(state, best[0])
