import pygame

WIDTH = 800
ROWS = 8

RED= pygame.image.load('./static/red_new.png')
GREEN= pygame.image.load('./static/green.png')

REDKING = pygame.image.load('./static/red.png')
GREENKING = pygame.image.load('./static/green.png')

WHITE = (255,255,255)
BLACK = (0,0,0)
ORANGE = (235, 168, 52)
BLUE = (76, 252, 241)

pygame.init()
WIN = pygame.display.set_mode((WIDTH,WIDTH))
pygame.display.set_caption('Checkers')

"""
Was brauche ich? 

- Map Funktion von [[R,G," "],[R,G," "], [R,G," "], [R,G," "], [R,G," "] ] --> Grid mit Nodes welche angezeigt werden können 

- Grid mit Nodes zu [[R,G," "],[R,G," "], [R,G," "], [R,G," "], [R,G," "] ] --> Map Funktion

-  ..... 

"""
class Node:
    def __init__(self, row, col, width):
        self.row = row
        self.col = col
        self.x = int(row * width)
        self.y = int(col * width)
        self.colour = WHITE
        self.piece = None

    def draw(self, WIN):
        pygame.draw.rect(WIN, self.colour, (self.x, self.y, WIDTH / ROWS, WIDTH / ROWS))
        if self.piece:
            WIN.blit(self.piece.image, (self.x, self.y))

class Piece:
    def __init__(self, team):
        self.team=team
        self.image= RED if self.team=='R' else GREEN
        self.type=None

    def draw(self, x, y):
        WIN.blit(self.image, (x,y))


def convert_array_to_printable_grid(array):
    grid = []
    for i in range(8):
        grid.append([])
        for j in range(8):
            node = Node(j,i, WIDTH / ROWS)
            node.colour = BLACK if abs(i - j) % 2 == 0 else WHITE
            grid[i].append(node)
            if array[i][j] == 'R':
                grid[i][j].piece = Piece('R')
            elif array[i][j] == 'G':
                grid[i][j].piece = Piece('G')
    return grid

def convert_printable_grid_to_array(grid):
    array = []
    for i in range(8):
        array.append([])
        for j in range(8):
            if grid[i][j].piece:
                array[i].append(grid[i][j].piece.team)
            else:
                array[i].append(" ")
    return array





def update_display(win, grid):
    for row in grid:
        for node in row:
            node.draw(win)
    pygame.display.update()




def make_grid(game_board):
    return convert_array_to_printable_grid(game_board)




def getNode(grid, rows, width):
    gap = width//rows
    RowX,RowY = pygame.mouse.get_pos()
    Row = RowX//gap
    Col = RowY//gap
    return (Col,Row)


def resetColours(grid, node, generatePotentialMoves):
    positions = generatePotentialMoves(node, grid)
    positions.append(node)

    for colouredNodes in positions:
        nodeX, nodeY = colouredNodes
        grid[nodeX][nodeY].colour = BLACK if abs(nodeX - nodeY) % 2 == 0 else WHITE

def HighlightpotentialMoves(piecePosition, grid, generatePotentialMoves):
    positions = generatePotentialMoves(piecePosition, grid)
    for position in positions:
        Column,Row = position
        grid[Column][Row].colour=BLUE

def opposite(team):
    return "R" if team=="G" else "G"