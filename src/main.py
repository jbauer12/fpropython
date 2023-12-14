import sys
from logic.gui import *
import logic.possible_moves as possible_moves
import logic.rules as rules
from logic.minimax import minimax_for_gui


def move(grid, piecePosition, newPosition, currPlayer):
    resetColours(grid, piecePosition, possible_moves.get_all_possible_moves_for_piece_gui, currPlayer=currPlayer)
    game_board = convert_printable_grid_to_array(grid, currPlayer)
    new_game_board = rules.make_move(game_board=game_board, newPosition=newPosition, piece=game_board.game_board[piecePosition[0]][piecePosition[1]])
    grid = make_grid(new_game_board)
    return grid, new_game_board.currPlayer

def handle_quit():
    print('EXIT SUCCESSFUL')
    pygame.quit()
    sys.exit()

def handle_mouse_click(grid, highlighted_piece, curr_move):
    clicked_node = getNode(grid, ROWS, WIDTH)
    clicked_position_column, clicked_position_row = clicked_node

    if grid[clicked_position_column][clicked_position_row].colour == BLUE:
        if highlighted_piece:
            piece_column, piece_row = highlighted_piece

        if curr_move == grid[piece_column][piece_row].piece.team:
            resetColours(grid, highlighted_piece, possible_moves.get_all_possible_moves_for_piece_gui, currPlayer=curr_move)
            grid, curr_move = move(grid, highlighted_piece, clicked_node, curr_move)
            update_display(grid)

    elif highlighted_piece == clicked_node:
        pass
    else:
        if grid[clicked_position_column][clicked_position_row].piece:
            if curr_move == grid[clicked_position_column][clicked_position_row].piece.team:
                highlighted_piece = highlight(clicked_node, grid, highlighted_piece, curr_move)

    return grid, highlighted_piece, curr_move

def handle_computer_move(grid, highlighted_piece, curr_move):
    resetColours(grid, highlighted_piece, possible_moves.get_all_possible_moves_for_piece_gui, currPlayer=curr_move)
    game_board = minimax_for_gui(state=convert_printable_grid_to_array(grid, curr_move), player=curr_move)
    return make_grid(game_board), game_board.currPlayer

def main(WIDTH, ROWS):
    grid = make_grid(rules.get_initial_game_board())
    highlightedPiece = None
    currMove = 'G'

    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                handle_quit()

            if event.type == pygame.MOUSEBUTTONDOWN and currMove == 'G':
                grid, highlightedPiece, currMove = handle_mouse_click(grid, highlightedPiece, currMove)

            if currMove == 'R':
                grid, currMove = handle_computer_move(grid, highlightedPiece, currMove)
                
        update_display(grid)

if __name__ == "__main__":
    main(WIDTH, ROWS)
