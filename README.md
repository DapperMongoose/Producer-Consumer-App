# Producer-Consumer-App

This was a test project for an interview.  Unfortunately I was not selected to go further with that company.  The company provided the prompt/requirements for the app but none of the code or tests were provided and are entirely my own work.

# Running the App

In order to run this super cool program all you will need to do is run go build from the root of the project directory.  Once that is done you may execute the resulting producer-consumer-project binary with the following flags:

-interval \<time\> - Allows you to set how often a new random number is generated. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

It should be noted that setting this to a sub-millisecond time will cause high resource utlization, and that setting it to a value greater than one minute will result in only one value being generated before the program terminates.

-concurrency \<integer\> - The number of workers that will consume random numbers.  Setting this to a very high number may adversely affect performance.

The program will run for one minute, but may be cancelled at any time with a SIGINT or SIGTERM signal (IE: pressing CTRL-C in a Linux terminal)


# Testing
Tests are included for this program.  In order to run them you may simply run go test ./... from the root of the project.  They include one test that runs for ten seconds which may be omitted by running go test with the -short flag.
