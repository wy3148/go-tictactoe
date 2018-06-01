# Tic Tac Toe

As far as I was able to implement tic-tac-toe for an interview that was time
boxed. Although they don't send you a list of every feature that you develop if
you can move extremely fast, at least the following features are still missing:

1. Properly detect end game conditions of WIN / LOSE / TIE
2. Actually implement a functional computer opponet (unbeatable? minimax?).
   Initially it just makes a valid move, not a good one.

Since I was writing this as fast as possible in order to try to get as far into
the problem, the following were the things I expected to change from my initial
implementation over time:

1. Players are basically just glorified `string` representing the letter they
   chose. This is where I'd want board introspection, and would have the
   computer logic live.
2. The board needs to be readable by the Player, so in addition to the stringer,
   it would need to return a `[3][3]string` representation for analyzing.

Beyond that, I wasn't sure where they were going to have me run with it, but
here you are.
