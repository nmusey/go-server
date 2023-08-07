# Go Server Template

This is a template repository to get up and running with a dockerized Go server and Postgres database. 

## Features
- Uses Docker so everything is containerized
- Uses Docker Compose to start and stop everything easily
- Uses [Air](https://github.com/cosmtrek/air) for live reloading of the server

## Instructions
1. Create a Github repo for the new project and set this as the template
2. `docker compose up --build` to run everything
3. Wait a few minutes to pull the images and do the first time build
4. Create your packages in `./server` and get to work
    4a. Any top level packages (ie at the same level as `main.go` must be added to the end of the `build.cmd` line in `.air.toml`
