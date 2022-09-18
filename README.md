# GreedyGameAssignment

## Tech Stack
- Go
- Docker

## How To Run
- Clone the repository and `cd` into it
- Run the command `docker compose -f .\docker-compose-pull.yml up`
- The app should've started on Port 3000, 3001

## How To Use
- Open `:3000/swagger/` for getting the endpoints and models
- Use the `:3001/addBidder`end point to add a Bidder
- Use the `:3001/removeBidder`end point to add a Bidder
- Finally, use `:3000/startAuction` to start a new auction. This will return the max value under 200 ms.
- Check the logs on the terminal for a more detailed view.
- OPTIONAL: Use the `postman.json` file to import the workspace to Postman for testing

## How To Build
- Clone the repository and `cd` into it
- Run the command `docker compose -f .\docker-compose-build.yml build`
- The app should've started on Port 3000, 3001
