# guess-it-2

Guess-it-2 is an excersice in the 01-edu curriculum and a program that guesses in what range a next number might fit based on previous numbers and their regression line.

To run the auditing program, extract the provided zip file and copy the student folder to what is by default called guess-it-dockerized. 

Now run the testing program with:
``` bash
docker compose up
```

Guess against at least these three ai models: "big-range", "linear-regr" and "correlation-coef".\
Also available: "average", "huge-range", "median", "mse" and "nic" 

To test "big-range", navigate to this URL:
```
http://localhost:3000/?guesser=big-range
```

### Notes about getting testing to work
To make testing work:
- Abide by the instructed folder structure
- script.sh should include these lines:\
cd student\
go run solution.go  
- Make sure your project go version is 1.23 (or the same the dockerized program uses)
- To run the tester with an updated version of the program in the student folder: close Docker, delete the container and the image from Docker desktop, copy the updated solution.go over the old one and run "docker compose up" again.
- Modifying permissions of script.sh and the executable ai models in the ai/ folder might be necessary

