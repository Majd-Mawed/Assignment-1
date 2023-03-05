The PDF document includes screenshots
Overview
. In the following READ ME, you will find a guide on how to use the RESTFULL web application that I have developed in Golang that makes it possible for the client to retrieve information about universities using the following websites: 
1.http://universities.hipolabs.com/ - used to get the information about the given university 
2.https://restcountries.com/ - used to get the information about the country and the borders of the country, we use this to get the information regarding the neighbour universities

How to use it?
First you can run the program by typing go run main.go in the terminal
Now in order to use this program, you need to type the following in either a browser or postman.
http://localhost:8080//unisearcher/v1/uniinfo/ 
http://localhost:8080//unisearcher/v1/neighbourunis/
http://localhost:8080//unisearcher/v1/diag/
Be aware that it might takes some time, because using the root link will give you all of the information, in order to fix that, we have to specify what we are looking for. We can start with a guide on how to use uniinfo path.
By typing http://localhost:8080//unisearcher/v1/uniinfo/, we will get the following result:
 
Now we have to simply write the university name like for example:
 
Be Aware that you have to type + instead of %20 for space. We can also type the country. 
 
Now for the second endpoint which is neighbour countries, by typing the following http://localhost:8080//unisearcher/v1/neighbourunis/ we should get the following result
 
The result only contains the country name with its borders.  And by writing the name of a country, we will get the following result.
 
We can get more information by typing the name of the university. 
 
Here we see that we have more information like, the name of the university, country name, the alpha two code of the country, the languages used in that country, and finally it is location
Now for our final endpoint, the diagnostic which we get access to by simply typing http://localhost:8080//unisearcher/v1/diag/, and we should get the following result:
 
The first two lines are the status code for these websites:
http://universities.hipolabs.com/ 
https://restcountries.com/  
While the third line is hardcoded, and it is the second part of the link 
The last line is the uptime.
