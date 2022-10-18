# Adapter Layer

Learn more about the adapter layer, through the implementation of an example.





To finish off our work with the example of a restaurant, let us implement the adapter layer on top of the port layer.







### Example definition

We will continue with the example that was defined in the [Domain Layer lesson](https://www.educative.io/pageeditor/10370001/4616975235416064/6471623718207488#Example-definition) and finish the implementation. The last layer is the adapter layer. This layer is in charge of communication with the outside world, for example, communication with database connections, HTTP, and so on.



### Model definition

The following diagram depicts the adapter layer, which defines six artifacts. Additionally, this layer makes use of the port layer. Due to its large size, this model is divided into two parts. Part one shows the domain, use case, and port layers, whereas part two shows the interaction between the port and abstract layers.



![image-20220510173243125](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510173243125.png)



The image above shows what was implemented previously. Now, it is time to understand how the adapter layer interacts only with the port layer. This interaction can be seen in the image below.

![image-20220510173414330](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510173414330.png)



### Model explanation

This layer contains the following artifacts:



- A `Rest` handler: This handler exposes a REST API.
- A `Request` object: This object defines the data that is required to carry out an order creation.
- A `Response` object: This object defines the data that is generated after an order creation is executed.
- An `InMemoryCustomerRepository` object: This object implements the `CustomerRepository`, which is located in the domain layer.
- An `InMemoryProductRepository` object: This object implements the `ProductRepository`, which is located in the domain layer.
- An `InMemoryOrderRepository` object: This implements the `OrderRepository`, which is located in the domain layer.





## Summary

The adapter layer contains all of the components that supply the interfaces defined in the port layer. Additionally, all the objects in this layer interact with the external world, for instance, REST API implementations, database connections, and so on.





