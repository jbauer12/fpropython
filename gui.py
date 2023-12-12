import pygame
from typing import List, Tuple
from possible_moves import get_initial_game_board
from classes import Piece as logicPiece, GameBoard

WIDTH = 800
ROWS = 8

RED = pygame.image.load('./static/red_new.png')
GREEN = pygame.image.load('./static/green.png')

REDKING = pygame.image.load('./static/red_dame.png')
GREENKING = pygame.image.load('./static/green_dame.png')

WHITE = (255, 255, 255)
BLACK = (0, 0, 0)
ORANGE = (235, 168, 52)
BLUE = (76, 252, 241)
pygame.init()
WIN = pygame.display.set_mode((WIDTH, WIDTH))
pygame.display.set_caption('Checkers')




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


class Piece:
    def __init__(self, team):
        self.team = team
        self.type = None
        self.image = RED if team == 'R' else GREEN

    def draw(self, x, y):
        WIN.blit(self.image, (x, y))



def convert_array_to_printable_grid(board: GameBoard) -> List[List[Node]]:
    grid: List[List[Node]] = []
    game_board = board.game_board
    for i in range(8):
        grid.append([])
        for j in range(8):
            node = Node(j, i, WIDTH / ROWS)
            node.colour = BLACK if abs(i - j) % 2 == 0 else WHITE
            grid[i].append(node)
            if game_board[i][j].team == 'R' and not game_board[i][j].king:
                grid[i][j].piece = Piece('R')
            elif game_board[i][j].team == 'G' and not game_board[i][j].king:
                grid[i][j].piece = Piece('G')
            elif game_board[i][j].team == 'R' and game_board[i][j].king:
                grid[i][j].piece = Piece('R')
                grid[i][j].piece.type = 'KING'
                grid[i][j].piece.image = REDKING
            elif game_board[i][j].team== 'G' and game_board[i][j].king:
                grid[i][j].piece = Piece('G')
                grid[i][j].piece.type = 'KING'
                grid[i][j].piece.image = GREENKING
    return grid


def convert_printable_grid_to_array(grid: List[List[Node]], currPlayer:str) -> GameBoard:
    game_board = tuple(
        tuple(logicPiece(team=node.piece.team,position=(row_index,col), king=True) if node.piece and node.piece.type == "KING" \
              else logicPiece(team=node.piece.team, position=(row_index,col), king=False) if node.piece and not node.piece.type =="KING" \
                else logicPiece(team=" ", position=(row_index,col), king=False)
              for col,node in enumerate(row))
        for row_index,row in enumerate(grid)
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


