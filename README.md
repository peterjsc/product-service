# Product Item calculate packs service



## Getting Started

The following are instructions on how to copy the project and run it locally on your machine.

### Prerequisites

To successfully run the project you will need to make sure you have the following software installed:

* **GitHub** - for cloning the repo of the project
* **Visual Studio Code** or other faviourite IDE - for viewing, running or building source code 
* **Docker** - for running services in containers
* **Postman** - for sending requests to service or use Invoke-WebRequest command in powershell

### Setup

Clone repo using Git and open code in Visual Studio Code. Make sure Docker is enabled and running on your local machine. Use the Terminal window in Visual Studio Code to navigate to to `deploy` folder.

`cd deploy`

Once in deploy folder, to build the services use the following docker command to execute the `docker-compose.yaml` file.

`docker-compose up --build`

### What gets built

Once docker compose have finished running, the following containers get created:

* **productservice** - container application of the Product service 
* **mongo** - container of the mongoDB database 
* **mong-seed** - container that is run against the mongoDB database to do a bulk import of data of for the Product services to consume

## Service API

### Product Service

Allows for a user to specify a 'Product' along with the 'ordered' amount they want to purchased of that product. The service will returns a json object that contains the product name, and item-packs that will consists of the item-pack and quantity of that item pack to fullfil that order number specified by the user. The service is on Port: 8080

----

**Get Product item pack endpoint**

* Endpoint used to recieve and calcuate item packs for a given product and item ordered amount
* **URL:**
    /products/{product-name}
    * Example url `http://localhost:8080/products/socks?ordered=501`

* **Method:**
    GET

* **URL Parms**
  **Required**
  `ordered=[integer]`

* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{ "product": { "product-name": "socks", "item-packs": [{ "item-pack": "500 items",  "quantity": 1 } ]} }`

* **Error Response:**

  * If no product-name is given in the Url or if the product does not exist in the database the service will give an error of 404 Not found, below is an example if there the product supplied doesnt exist.
  * **Code:** 404 Not Found <br />
    **Content:** `{ "error": "Product doesn't exist"}`

  * Error occurs when a string is supplied for the 'ordered' url param 
  * **Code:** 400 Bad Request <br />
    **Content:** `{"error": "Error converting string to int"}`

  * Error occurs when 0 is supplied for the 'ordered' url param 
  * **Code:** 400 Bad Request <br />
    **Content:** `{"error": "Please supply an order amount greater than 0"}`
  
  * Error occurs if no parameter called 'ordered' is supplied or if it has no value 
  * **Code:** 400 Bad Request <br />
    **Content:** ` { "error": "Invalid Request no Item orders supplied, please supply value for 'ordered' parameter" }`

----

**Add Product endpoint**

* Endpoint used to create a new product with item-packs, post url needs a header of content-type 'application/json' and a body for item-packs see example below of a post
* **URL:**
    /products/{product-name}
    * Example url `http://localhost:8080/products/running-top`
    * Header: `"Content-Type"="application/json; charset=utf-8"`
    * Body: `{ "item-packs": [ 250, 500, 1000, 2000, 5000 ] }`

* **Method:**
    POST

* **Success Response:**
  * **Code:** 201 <br />
    **Content:** `{ Product is now created }`

* **Error Response:**

  * Error if trying to add product that already exists
  * **Code:** 409 Conflict <br />
    **Content:** `{"error": "Product already exists"}`
  
  * Error if no content type is supplied in header or a request body
  * **Code:** 415 Unsupported Media Type <br />
    **Content:** `{ "error": "Content-Type header is not application/json" }`

  * Error if content type is supplied in header but not body is supplied in the request or supplies an incorrect json body
  * **Code:** 400 Bad Request <br />
    **Content:** `{ "error": "Please supply a body in json for request or in correct format" }`


## Useful Tips

For accessing and querying the mongodb database here are a few useful commands 

For connecting to the mongodb container using the following command:
`docker exec -it mongodb bash`

Then connect to the database using the credentials 
`mongo admin -u root -p example`

This shows the databases once connected `show dbs` and `use products` picks the database to work with.

The `show collections` shows the collections withing that database, there will be a  collections `product` for product data.

The following to commands can be used to bring back all the data in the product collections, `db.product.find()`