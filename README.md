# Go REST API microservice sample with Redis

### WORK IN PROGRESS!!!

This is a deployable REST API microservice in Go that receives (or fetches) an XML document and processes it against an XSLT by receiving parameters to produce a final XML document output. This is a work in progress that has pending its tests and additional features such as Redis caching.

## Getting Started

### Prerequisites

- Install Docker and/or Docker-compose on system.
- Install Go on system (this is optional in case you want to run it locally).

### Installing

#### Running with Docker

- Clone this repo on local machine and navigate into directory.
- Build and run application: ```docker-compose up```
- App will be served at localhost:3030

### TODO:
- Implement redis route-caching system ?
- Implement route for "Accept: application/json" headers
- Unit tests

## WORK IN PROGRESS!!!
![Image of Gopher](https://cdn-images-1.medium.com/max/500/1*vHUiXvBE0p0fLRwFHZuAYw.gif)
