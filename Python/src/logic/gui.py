import pygame
from enum import Enum
from typing import List, Optional
from logic.classes import Piece as logicPiece, GameBoard
from logic.possible_moves import get_all_possible_moves_for_piece_gui

WIDTH = 800
ROWS = 8



WHITE = (255, 255, 255)
BLACK = (0, 0, 0)
ORANGE = (235, 168, 52)
BLUE = (76, 252, 241)
pygame.init()
WIN = pygame.display.set_mode((WIDTH, WIDTH))
pygame.display.set_caption('Checkers')


class Image(Enum):
    RED = pygame.image.load('Python/src/static/red_new.png')
    GREEN = pygame.image.load('Python/src/static/green.png')

    REDKING = pygame.image.load('Python/src/static/red_dame.png')
    GREENKING = pygame.image.load('Python/src/static/green_dame.png')

class Node:
    def __init__(self, row, col, width):
        self.row = row
        self.col = col
        self.x = int(row * width)
        self.y = int(col * width)
        self.colour = WHITE
        self.piece = None

    def draw(self, WIN):
        pygame.draw.rect(WIN, self.colour, (self.x, self.y,
                         WIDTH / ROWS, WIDTH / ROWS))
        if self.piece:
            WIN.blit(self.piece.image, (self.x, self.y))
    def set_piece(self, piece):
        self.piece = piece
        return self


class Piece:
    def __init__(self, team):
        self.team = team
        self.type = None
        self.image = Image.RED.value if team == 'R' else Image.GREEN.value

    def draw(self, x, y):
        WIN.blit(self.image, (x, y))




def convert_array_to_printable_grid(board: GameBoard) -> List[List[Node]]:
    game_board = board.game_board

    def create_node(row, col):
        node = Node(col, row, WIDTH / ROWS)
        node.colour = BLACK if abs(row - col) % 2 == 0 else WHITE
        return node

    def create_piece(team, king):
        piece = Piece(team)
        if king:
            piece.type = 'KING'
            piece.image = Image.REDKING.value if team == 'R' else Image.GREENKING.value
        return piece

    grid = [[create_node(i, j)
            .set_piece(create_piece(game_board[i][j].team, game_board[i][j].king) if game_board[i][j].team != " " else None)
            for j in range(8)]
        for i in range(8)]

    return grid




def convert_printable_grid_to_array(grid: List[List[Node]], currPlayer: str) -> GameBoard:
    def create_logic_piece(node: Optional[Node], row_index: int, col: int) -> logicPiece:
        if node and node.piece:
            if node.piece.type == "KING":
                return logicPiece(team=node.piece.team, position=(row_index, col), king=True)
            else:
                return logicPiece(team=node.piece.team, position=(row_index, col), king=False)
        else:
            return logicPiece(team=" ", position=(row_index, col), king=False)
        
    game_board = tuple(
        tuple(create_logic_piece(node, row_index, col) for col, node in enumerate(row))
        for row_index, row in enumerate(grid)
    )
    return GameBoard(game_board=game_board, currPlayer=currPlayer)



def update_display(grid):
    for row in grid:
        for node in row:
            node.draw(WIN)
    pygame.display.update()


def make_grid(game_board):
    return convert_array_to_printable_grid(game_board)


def getNode(grid, rows, width):
    gap = width//rows
    RowX, RowY = pygame.mouse.get_pos()
    Row = RowX//gap
    Col = RowY//gap
    return (Col, Row)


def resetColours(grid, node, generatePotentialMoves, currPlayer:str):
    computing_grid = convert_printable_grid_to_array(grid, currPlayer=currPlayer)
    positions = generatePotentialMoves(computing_grid, computing_grid.game_board[node[0]][node[1]])
    positions.append(node)

    for colouredNodes in positions:
        nodeX, nodeY = colouredNodes
        grid[nodeX][nodeY].colour = BLACK if abs(
            nodeX - nodeY) % 2 == 0 else WHITE


def HighlightpotentialMoves(piecePosition, grid, generatePotentialMoves, currPlayer:str):
    computing_grid = convert_printable_grid_to_array(grid=grid, currPlayer=currPlayer)
    positions = generatePotentialMoves(computing_grid, computing_grid.game_board[piecePosition[0]][piecePosition[1]])
    for position in positions:
        Column, Row = position
        grid[Column][Row].colour = BLUE


def highlight(ClickedNode, Grid, OldHighlight, currPlayer:str):
    Column, Row = ClickedNode
    Grid[Column][Row].colour = ORANGE
    if OldHighlight:
        resetColours(Grid, OldHighlight,
                     generatePotentialMoves=get_all_possible_moves_for_piece_gui, currPlayer=currPlayer)
    HighlightpotentialMoves(
        ClickedNode, Grid, get_all_possible_moves_for_piece_gui, currPlayer=currPlayer)
    return (Column, Row)