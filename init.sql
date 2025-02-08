-- Drop the tables if they exist
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS team_players;
DROP TABLE IF EXISTS tournaments;
DROP TABLE IF EXISTS tournament_teams;
DROP TABLE IF EXISTS tournament_players;


-- Create the users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL UNIQUE
);

-- Create the players table
CREATE TABLE players (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Create the teams table
CREATE TABLE teams (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

-- Create the team_players table
CREATE TABLE team_players (
    id INTEGER PRIMARY KEY,
    team_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (player_id) REFERENCES players(id)
);

-- Create the tournaments table
CREATE TABLE tournaments (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL
);

-- Create the tournament_teams table
CREATE TABLE tournament_teams (
    id INTEGER PRIMARY KEY,
    tournament_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id),
    FOREIGN KEY (team_id) REFERENCES teams(id)
);

-- Create the tournament_players table
CREATE TABLE tournament_players (
    id INTEGER PRIMARY KEY,
    tournament_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id),
    FOREIGN KEY (player_id) REFERENCES players(id)
);
