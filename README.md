# PokedexCLI

A command-line Pokédex application written in Go that lets you explore the Pokémon world, catch Pokémon, and build your collection using the [PokéAPI](https://pokeapi.co/).

## Features

- **🗺️ World Exploration**: Navigate through location areas in the Pokémon world
- **🔍 Area Exploration**: Discover which Pokémon can be found in specific locations
- **⚾ Pokémon Catching**: Attempt to catch Pokémon with realistic difficulty based on their stats
- **📖 Pokédex Management**: View your caught Pokémon collection
- **🔎 Pokémon Inspection**: Get detailed stats and information about your caught Pokémon
- **⚡ Smart Caching**: Built-in HTTP response caching for better performance
- **🎲 Probability-Based Catching**: Legendary Pokémon are harder to catch than common ones

## Installation

### Prerequisites

- Go 1.24.3 or later

### Build from Source

1. Clone the repository:

```bash
git clone https://github.com/karprabha/pokedexcli.git
cd pokedexcli
```

2. Build and run using the provided script:

```bash
chmod +x build_and_run.sh
./build_and_run.sh
```

Or build manually:

```bash
go build -o pokedexcli .
./pokedexcli
```

## Usage

Once the application starts, you'll see the `Pokedex >` prompt. Here are the available commands:

### Basic Commands

- **`help`** - Display all available commands and their descriptions
- **`exit`** - Exit the Pokédex application

### Navigation Commands

- **`map`** - Display the next 20 location areas in the Pokémon world
- **`mapb`** - Display the previous 20 location areas

### Exploration Commands

- **`explore <area_name>`** - Explore a specific location area to see what Pokémon can be found there
  ```
  Pokedex > explore canalave-city-area
  ```

### Pokémon Commands

- **`catch <pokemon_name>`** - Attempt to catch a Pokémon

  ```
  Pokedex > catch pikachu
  ```

- **`inspect <pokemon_name>`** - View detailed information about a Pokémon you've caught

  ```
  Pokedex > inspect pikachu
  ```

- **`pokedex`** - View all Pokémon in your collection

## Example Session

```
Pokedex > help

Welcome to the Pokedex!
Usage:

catch: Attempt to catch a Pokemon and add it to your Pokedex.
exit: Exit the Pokedex
explore: Explore a location area to see all Pokemon found there.
help: Displays a help message
inspect: View details about a caught Pokemon.
map: Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations
mapb: Displays the previous 20 location areas
pokedex: See all the pokemon in your Pokedex.

Pokedex > map
Making API request...
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect command.

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  - electric

Pokedex > pokedex
Your Pokedex:
  - pikachu

Pokedex > exit
```

## Architecture

The project is organized into several packages:

- **`main`** - Entry point that starts the REPL
- **`internal/repl`** - Read-Eval-Print-Loop implementation
- **`internal/commands`** - All CLI commands and their implementations
- **`internal/pokecache`** - HTTP response caching system with automatic cleanup

### Key Features

- **Thread-safe caching**: Implements a concurrent-safe cache with automatic expiration
- **Probability-based catching**: Legendary Pokémon have lower catch rates
- **Graceful error handling**: Handles network issues and invalid inputs
- **Memory efficient**: Caches API responses to reduce network calls

## Development

### Running Tests

```bash
go test -v ./...
```

### Using the Build Script

The `build_and_run.sh` script provides several options:

```bash
./build_and_run.sh build    # Build only
./build_and_run.sh test     # Run tests only
./build_and_run.sh run      # Run the application
./build_and_run.sh all      # Build, test, and run (default)
./build_and_run.sh clean    # Clean up build artifacts
./build_and_run.sh help     # Show help
```

## Future Enhancement Ideas

Here are some ideas to extend the project:

### User Experience Improvements

- **Command History**: Add support for the "up" arrow to cycle through previous commands
- **Auto-completion**: Implement tab completion for Pokémon and location names
- **Colored Output**: Add colors to make the CLI more visually appealing

### Game Mechanics

- **Pokémon Battles**: Simulate battles between Pokémon in your collection
- **Pokémon Party**: Keep Pokémon in an active "party" and allow them to level up
- **Evolution System**: Allow caught Pokémon to evolve after meeting certain conditions
- **Random Encounters**: Add random wild Pokémon encounters while exploring

### Enhanced Catching System

- **Different Poké Balls**: Support for Poké Balls, Great Balls, Ultra Balls with varying catch rates
- **Pokémon Status Effects**: Factor in sleep, paralysis, etc. for catch probability
- **Berries**: Use berries to increase catch rates or calm Pokémon

### Data Persistence

- **Save Progress**: Persist user's Pokédex to disk to save progress between sessions
- **Multiple Profiles**: Support for multiple user profiles
- **Statistics Tracking**: Track catching success rates, exploration history, etc.

### Advanced Features

- **Enhanced Navigation**: Instead of typing area names, provide directional choices ("left", "right", "north", "south")
- **Quest System**: Add simple quests or challenges
- **Trading System**: Simulate trading Pokémon with NPCs
- **Shiny Pokémon**: Rare variants of Pokémon with different colors

### Technical Improvements

- **More Unit Tests**: Increase test coverage across all packages
- **Better Error Messages**: More user-friendly error reporting
- **Configuration File**: Allow users to customize settings
- **Logging**: Add structured logging for debugging
- **Performance Metrics**: Track and display API response times

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the comprehensive Pokémon data
- The Go community for excellent tooling and libraries

## API Usage

This project uses the PokéAPI, which is free and requires no authentication. The application implements caching to be respectful of the API's resources and improve performance.
