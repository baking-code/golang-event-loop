# GoDash

Popular (and by no means exhaustive!) methods from https://lodash.com/docs implemented in Go.

Started as a learning exercise, some of these utilities are instrinsic to the JS ecosystem (many have been made 1st class language features) so this could be a good resource for new engineers learning Golang.

## Structure

- Each function in it's own folder and package
- Each function accompanied by a test file
- Any internal functions/helpers not part of lodash are in the `internal` directory

## Contributing

If you'd like to contrbute, just create a PR targeting `main` which follows the same structure to get it included. As this project evolves more stringent CI will be introduced.
