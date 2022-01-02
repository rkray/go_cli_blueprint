# go_cli_blueprint
Blueprint for go cli programs

# TODO
- [x] general program structure 
- [x] start with go module
- [x] read comandline parameters
- [x] define global flags
- [ ] define local flags
- [ ] use local parameters
- [x] help page
- [ ] signal handling, what happens when interupting by CTR+C
- [ ] support for json/yaml config files in $HOME/.progname\_(json|yaml).cfg and $HOME/.config/progname/config.(json|yaml)

# Compile
    ./make.sh build && ./go_cli_blueprint help

# Cleanup
    go clean
