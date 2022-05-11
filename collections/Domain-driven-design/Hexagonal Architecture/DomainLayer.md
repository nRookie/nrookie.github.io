To comprehend how software may be implemented, based on hexagonal architecture, let us review an example related to the processes that occur in a restaurant.



## Example definition



This restaurant wants to allow a customer to order something through its website. The website consumes **REST APIs** to execute the business logic. Therefore, we have to implement a microservice that exposes the functionality to create an order. This use case displays the following restrictions:

- If a customer has no previously created orders, it is possible to create a new order.
- If a customer has already created an order, it is not possible to create a new order.



In the upcoming lessons, each layer of this microservice will be implemented in order to accomplish the business need.





### Model definition

The following diagram depicts the domain layer and how the domain objects should be implemented to fulfill the business rules.



![image-20220510160330874](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510160330874.png)



### Model explanation



This layer contains the following domain objects:

- An `Order` object: This is the aggregate root. Hence, it will orchestrate the logic across other objects.
- A `Customer` object: This is an entity that holds the number of created orders and the logic used to validate them.
- An `Address` object: This is a value object that holds information about a customer’s address.
- A `Product` object: This is an entity that represents the product a client wants to order.
- A `CustomerRepository`: This is an interface, which defines the operations that are possible to carry out in the customer repository.
- A `ProductRepository`: This is an interface, which defines the operations that are possible to carry out in the product repository.
- An `OrderRepository`: This is an interface, which defines the operations that are possible to carry out in the order repository.



### Implementation



The code above shows how we can translate a model into code. In this case, the code implements the business functionalities that were modeled in the diagram above.





## Summary

The domain layer contains all of the logic related to the business model. Aggregates, entities, value objects, repositories, and services are all placed in this layer.





