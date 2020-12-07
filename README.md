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
