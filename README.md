# ♟️ chess-cli

A minimal chess engine built as a CLI application in Go. This project focuses on correctness, separation of concerns, and
extensibility.

------------------------------------------------------------------------

# 🏗️ Architecture

## High-Level Structure

    Game
     ├── Board
     │    ├── 8x8 grid
     │    └── Piece references
     │
     ├── Piece (interface)
     │    ├── Pawn
     │    ├── Rook
     │    ├── Bishop
     │    ├── Knight
     │    ├── Queen
     │    └── King
     │
     └── Move validation + turn handling

------------------------------------------------------------------------

## Core Components

### 1️⃣ Game

Responsible for:

-   Turn management
-   Input parsing
-   Executing moves
-   Win detection (king capture)
-   Orchestrating board state

------------------------------------------------------------------------

### 2️⃣ Board

Encapsulates:

-   8x8 grid representation
-   Piece placement
-   Piece retrieval
-   Move execution

------------------------------------------------------------------------

### 3️⃣ Piece Abstraction

Each chess piece implements movement logic independently.

------------------------------------------------------------------------

# ♟️ Move Validation Strategy

Validation flow:

1.  Validate turn ownership
2.  Ensure source contains a piece
3.  Delegate movement validation to the piece
4.  Ensure destination is empty or opponent-occupied
5.  Execute move
6.  Check win condition (king captured)

------------------------------------------------------------------------

# 🖥️ CLI Rendering

Board is rendered using ASCII layout:

8 | r  n  b  q  k  b  n  r
7 | p  p  p  p  p  p  p  p
6 | -  +  -  +  -  +  -  +
5 | +  -  +  -  +  -  +  -
4 | -  +  -  +  -  +  -  +
3 | +  -  +  -  +  -  +  -
2 | P  P  P  P  P  P  P  P
1 | R  N  B  Q  K  B  N  R
  ------------------------
    a  b  c  d  e  f  g  h

Uppercase → White\
Lowercase → Black

Input format:

    d2,d4
    d7,d5

------------------------------------------------------------------------

# ⚙️ Limitation

Excluded features:

-   Check detection
-   Checkmate validation
-   Castling
-   En passant
-   Pawn promotion
-   Draw rules
-   Move history

------------------------------------------------------------------------

# 🧪 Run Test

-   Unit testing per piece movement
-   Isolated board state testing
-   Deterministic move validation

Encapsulated movement logic allows modular testing.

Run test:

    go test

------------------------------------------------------------------------

# 🚀 How to Run

Requires Go 1.20+

    go run .

Or:

    go build -o chess
    ./chess

------------------------------------------------------------------------

# 📦 Tech Stack

-   Go

------------------------------------------------------------------------

# 👤 Author

Falahudin Halim Shariski
