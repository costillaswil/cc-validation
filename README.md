# cc-validation



## Installation

1. install `Go` (version 1.13+ is required)

2. Extract zip file to Go directory

3. Install Postman ( for API testing )

## Quick Start

1. open your CLI ( command line interface )

2. navigate to your extracted zip and locate `main.go` in the root folder

3. Run this command to initialize a Local connection for the API

```
    $ run main.go 
    Now listening on: http://localhost:8006
    Application started. Press CTRL+C to shut down.
```
4. Open any API testing tools ( e.g. Postman )

5. Using "POST" method key-in this url in your api testing tool address bar `http://localhost:8000/creditcard`.

6. Fill the request body, the content-type is `json`. 
   Card number is required,if left blank it will throw a status 400 `bad request` 

```
{
	"cardNumber": "4012888888881881"
}
```

7. If successfull response should be like this

```
  {
      "cardNumber": "4012888888881881",
      "cardType": "Visa",
      "status": "Valid"
  }
```
