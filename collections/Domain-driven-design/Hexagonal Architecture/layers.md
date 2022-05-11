As with any other architectural pattern, hexagonal architecture defines some layers that should be followed. We will learn about these layers in this lesson.



## The layers of hexagonal architecture

Hexagonal architecture defines four layers to structure a component. As was mentioned before, these layers are focused on all of the parts that an application consists of, such as frameworks, database connections, business flows, and so on. Let’s look at them one by one.



### Domain objects layer

Located in the center of the hexagon, a domain contains all of the objects related to the business. In DDD terms, business objects are entities, value objects, and aggregates. Additionally, this layer contains all of the abstractions that are defined to carry out a business process. The abstractions are factories and the definition of repositories. Services can result in two possibilities. If the services interact with external dependencies, the domain layer will only have access to the services’ definition. If they do not interact with external dependencies, then the domain layer will have access to the full implementation.



![image-20220510154233883](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510154233883.png)





### Use cases layer



The use cases layer exposes functionalities to fulfill business requirements. This layer is located on top of the domain object layer. It can orchestrate logic in an aggregate root, and between two or more of them if necessary. Suppose that there is a requirement to create an order in a restaurant. However, the customer cannot place another order, while the previous one is still in progress. To fulfill this requirement, we will need to search for customer-opened orders. It will be possible to execute the creation if the customer shows no record of opened orders. Otherwise, the request will be rejected. A use case layer can implement the previous scenario.



![image-20220510155201993](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510155201993.png)



### Ports layer

A **port** is an interface, which defines the functions and the data required to carry them out. Ports are placed between the domain logic and the outside world that is designed for a particular purpose or protocol. They are used to allow outside clients to access the domain logic, and topermit the business logic to access the external systems. To compare a port with an example from the real world, we can say that it is a plug that exposes a particular interface. Ultimately, this layer is located after the use case layer.



![image-20220510155433060](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510155433060.png)

### Adapters layer

The adapters layer is the outermost layer of the hexagonal architecture. It interacts with the outside world directly. An adapter uses a specific technology. Therefore, this layer has a high coupling with a specific technological stack. Such a stack may be REST or gRPC, which are used to expose functionalities, or it may be **JDBC** or **Cache**, which are used to store information.



The implementation of an adapter allows the external systems to execute the domain logic. An adapter must comply with the interface definition and send all of the required data to achieve that interaction.

![image-20220510155625380](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510155625380.png)





## Summary



Hexagonal architecture defines four layers for the placement of logic. These layers are: the domain layer, the use case layer, the ports layer, and the adapters layer.





