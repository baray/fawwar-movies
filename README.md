fawwar-movies
Movies library that gives the users ability to add movies to the library and save their watched movies
===================
- - - - 
# Movies #

	List of movies [GET] /movie-list
		- Retrieves all movies from the database and sorts them by date or rating if specified in the query
  parameters. Returns a list of movies with their ID, name, description, date, and cover.
	
	Create a movie [post] /movies
		- Takes the movie's name, description, and date, saves the information in the 
  database and only the user with the token can add a movie.

	Get movie info [GET] /movies/[id]
		- Takes the movie id as a parameter , then returns the description of that movie with the rating
  of that movie which is calculated by dividing the sum of all ratings by the number of raters.

	Update the cover of movie [POST] /movies/[id]/update-cover
		- Edit the cover of the movie by sending the cover photo.
	
	Update the movie [PUT] /movies/[id]
		- Takes the movie's name, description and date, and only the user with the token 
  can actually edit the movie in the database.

	Delete the movie [DELETE] movies/[id]
  	- Takes the movie ID as a parameter, and checks that the user who wants to delte the movie
  is the user with the token, then the movie gets deleted from the database.

- - - - 
# User / Authentication #

	Create an account [POST] /auth/register
		- takes name, age, email and password of the user and saves it in the database.
	
	Login [POST] /auth/login
		- takes email and password of the user and returns a token for the user .
the token is encrypted.

- - - - 
# Watch and Review #

	Get movies you've watched [GET] /watches
		- Retrieve all movies you watched that you can review and rate. in addition to sorting them 
		by date and average rating.

	Add a movie to the wateched list [POST] /watches/[id]/watch
		- takes name, age, email and password of the user and saves it in the database.

	Create an account [POST] /watches/[id]/rating
		- takes the rating as a query parameter from the user, the user must have watched the movie and 
  must have the token.
