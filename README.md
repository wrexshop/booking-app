# booking-app
We are following on along with this YT tutorial: https://www.youtube.com/watch?v=yyUHQIec83I&t=

We built a simple booking app on the command line.
In the program, we have split up the functions from the main func and created a separate package & file for the validation func
There is also a goroutine (concurrency) that splits out the ticketing email to its own thread.
The main func will wait for the go routine to finish before the program exist.