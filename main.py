import sys
from itertools import combinations
from gui import *
import possible_moves as possible_moves
from rules import make_move


# passt
def highlight(ClickedNode, Grid, OldHighlight):
    Column, Row = ClickedNode
    Grid[Column][Row].colour = ORANGE
    if OldHighlight:
        resetColours(Grid, OldHighlight,
                     generatePotentialMoves=possible_moves.get_all_possible_moves)
    HighlightpotentialMoves(
        ClickedNode, Grid, possible_moves.get_all_possible_moves)
    return (Column, Row)


def move(grid, piecePosition, newPosition):
    resetColours(grid, piecePosition, possible_moves.get_all_possible_moves)
    new_game_board = make_move(convert_printable_grid_to_array(
        grid), newPosition, piecePosition)
    grid = make_grid(new_game_board)

    return grid, opposite(grid[newPosition[0]][newPosition[1]].piece.team)


def main(WIDTH, ROWS):
    grid = make_grid(possible_moves.get_initial_game_board())
    highlightedPiece = None
    currMove = 'G'

    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                print('EXIT SUCCESSFUL')
                pygame.quit()
                sys.exit()

            if event.type == pygame.MOUSEBUTTONDOWN:
                clickedNode = getNode(grid, ROWS, WIDTH)
                ClickedPositionColumn, ClickedPositionRow = clickedNode
                if grid[ClickedPositionColumn][ClickedPositionRow].colour == BLUE:
                    if highlightedPiece:
                        pieceColumn, pieceRow = highlightedPiece
                    if currMove == grid[pieceColumn][pieceRow].piece.team:
                        resetColours(grid, highlightedPiece,
                                     possible_moves.get_all_possible_moves)
                        grid, currMove = move(
                            grid, highlightedPiece, clickedNode)
                elif highlightedPiece == clickedNode:
                    pass
                else:
                    if grid[ClickedPositionColumn][ClickedPositionRow].piece:
                        if currMove == grid[ClickedPositionColumn][ClickedPositionRow].piece.team:
                            highlightedPiece = highlight(
                                clickedNode, grid, highlightedPiece)

        update_display(grid)


main(WIDTH, ROWS)
