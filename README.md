# httpHammer
Allows stress testing of http Web server. 
<pre>
This makes use of Go's concurrency, and when run will open a user specified number of threads against the webserver being tested. 

]$ ./httpHammer -h 10 -u http://www.google.com
Loop: 1
200 OK, downloaded 17k in 338.119966ms
200 OK, downloaded 17k in 339.670717ms
200 OK, downloaded 17k in 341.477184ms
200 OK, downloaded 17k in 339.933475ms
200 OK, downloaded 17k in 340.62257ms
200 OK, downloaded 17k in 340.583227ms
200 OK, downloaded 17k in 344.300447ms
200 OK, downloaded 17k in 342.579672ms
200 OK, downloaded 17k in 342.89689ms
200 OK, downloaded 17k in 344.470692ms
Total Time: 3.41465484s
Average Time: 341.465484ms

Program also allows you to specify the number of loops to run through, or run continuously.

]$ ./httpHammer
Usage: httpHammer -u <url> -h <hits> (optional) -l <times to run> -r <regexp>
  -c=false: Run Continuously
  -h=1: Number of hits
  -l=1: Number of loops to run...
  -r="": Regular Expression
  -u="": URL

This will also use regular expressions to look for text within the page, and will also print out an etag if there is one in the page.
</pre>
