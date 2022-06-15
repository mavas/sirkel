# sirkel

The circle (drinking) game.

Circle is a game most optimally played with 5 or 6 people, all sitting in a circle, facing each other.  Players all do a simple 1-2-clap beat, and on the clap, one player throws their unique sign using their hands; they then on the immediate next clap, throw/display another player's sign.  The game continues until someone makes a mistake interpretting or throwing signs, and you have to take a drink.

sirkel is best-efforts attempt to simulate playing this game on a computer over a network.

## Protocol

The protocol, for now, uses text over TCP.  The client sends "START" upon connection, and the server sets up a "game", and starts playing the game with the participants.

The server continuously keeps track of what's called a "pace", for each game, and this gets updated frequently, according to the perceived pace of the participants.  If the participants play fast, the server expects things to get faster, and tells the participants.
