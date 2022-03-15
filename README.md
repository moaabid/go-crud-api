# Simple CRUD API using GO and Gorilla mux

## Summary

This project servers an simple CRUD movie Api using GO and Gorilla mux

Note: This project is not connected to any database. Here I manage data with struct and slices.

## Usage

### API

- GET <http://localhost:8080/movies>
- GET <http://localhost:8080/movies/{id}>
- POST <http://localhost:8080/movies/>  {
    "isbn": "3272812",
    "title": "The Adam Project",
    "director": {
      "firstname": "Shawn",
      "lastname": "Levy"
    }
  }
- PUT <http://localhost:8080/movies/1>  {
    "isbn": "3272811",
    "title": "The Adam Project",
    "director": {
      "firstname": "Shawn",
      "lastname": "Levy"
    }
  }
- DELETE <http://localhost:8080/movies/1>

### TODO

- [ ] Delete All movies
- [ ] Add Dockerfile
