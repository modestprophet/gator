# GATOR - RSS Feed Aggregator CLI

# PREREQUISITES
PostgreSQL 13+ installed and running
Go 1.16+ installed
Basic command line experience

# INSTALLATION
1. Install the CLI:
```bash
$ go install github.com/modestprophet/gator@latest
```
Ensure your $GOPATH/bin is in your PATH:
```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

# CONFIGURATION
Create a config.yaml file in your working directory with:
```bash
DB_URL: "postgres://username:password@localhost:5432/gator?sslmode=disable"
```
Create the database (if not exists):
```bash
$ createdb gator
```

# RUNNING
Run the program:
```bash
$ ./gator <command>
```

# MAIN COMMANDS
```
Register new user:
> register <username>

Login as user:
> login <username>

Add new RSS feed:
> addfeed <feed-name> <feed-url>

Follow a feed:
> follow <feed-url>

Browse posts:
> browse [limit]

Start aggregator:
> agg <interval> (e.g.: agg 1m30s)
```

# OTHER COMMANDS
```
users - List all users
feeds - List all feeds
following - Show followed feeds
unfollow - Stop following a feed
reset - Clear all users (admin)
```

# TIPS
The CLI maintains state between sessions using your local database
First run: register a user, then login to start adding feeds
Use 'agg' command with a duration to start automatic feed collection
All feed operations require being logged in first

# TROUBLESHOOTING
Ensure PostgreSQL is running and config.yaml has correct credentials
If getting database errors, try recreating the database:

Start by registering a user and adding your first RSS feed! Try adding a popular tech blog RSS feed URL to get started.