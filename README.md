###### README

follow the next instructions tp run the fizzbuzz service:
==================


1. Build go binary
> make build

2. Run HTTP server
> make app-up

#Note:
- The default host is 127.0.0.1 and if you want try the fizzbuzz app with your browser you can use this url:
> http://127.0.0.1:1337/list?int1=2&int2=3&string1=fizz&string2=buzz&limit=1000

- If you want change the default port used by the app, you can update the JSON file __fizzbuzz.json__ in the folder __/config/fizzbuzz__


- Run tests with command: 
> make test


