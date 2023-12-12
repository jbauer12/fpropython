from pydantic import BaseModel
from typing import Literal, List, Tuple, Union


class Piece(BaseModel):
    position: Tuple[int, int]
    #Literal -> entweder Team R oder Team G
    team: Literal["R", "G", " "]
    #Literal -> entweder King oder Normal
    king: bool = False

    def __str__(self):
        return f"|{self.team}"
    #frozen -> Nach Initialisieren keine Änderungen mehr möglich
    class Config:
        frozen = True

class PieceWithPositions(Piece):
    possible_positions: List[Tuple[int, int]] = []

class GameBoard(BaseModel):
    game_board: Tuple[Tuple[Piece, ...], ...]
    currPlayer: str = "R"
    def __str__(self):
        return "\n".join(" ".join(str(piece) for piece in row) for row in self.game_board)

    class Config:
        frozen = True

def print_game_board(board: GameBoard) -> None:
    game_board= board.game_board
    for row in game_board:
        for piece in row:
            print(piece, end=" ")
        print()