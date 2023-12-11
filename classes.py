from pydantic import BaseModel
from typing import Literal, Tuple, Union


class Piece(BaseModel):
    position: Tuple[int, int]
    #Literal -> entweder Team R oder Team G
    team: Literal["R", "G", " "]
    #Literal -> entweder King oder Normal
    king: bool = False
    #frozen -> Nach Initialisieren keine Änderungen mehr möglich
    def __str__(self):
        return f"|{self.team}"

    class Config:
        frozen = True

        
TYPE_GAMEBOARD = Tuple[Tuple[Piece, ...], ...]
def print_game_board(game_board: TYPE_GAMEBOARD) -> None:
    for row in game_board:
        for piece in row:
            print(piece, end=" ")
        print()