# Services

Learn what services are in DDD and how they allow us to place logic that does not fit into any entity or value object.





There may be instances where it is required to execute a particular logic. As was previously mentioned, entities and value objects can contain business logic. However, what happens if the particular logic does not fit into any entity or value object? The answer can be found in the functionality of a service.



## What are services?

First of all, these objects are called **domain services** in the context of DDD. **Services** are objects that implement domain functionality, which cannot be modeled naturally in any entity or value object as part of the domain logic. Domain services are usually implemented **when the logic of two or more domain objects is invoked**. If service logic needs to interact with an external system, it should be built with the definition of an interface in the domain layer and the implementation of it in another layer.



### Characteristics of a domain service



Domain services should meet the following characteristics to be well-defined and isolated from the domain model:



- They should be stateless. They should not maintain information between calls. Once an execution is completed, all of the information related to the execution should disappear.
- They can produce domain events. This may, in turn, cause the execution of other bounded contexts.
- They must be highly cohesive. These objects must execute only one specific task. Since the domain logic contained within domain services does not fit elsewhere, they need to accomplish domain logic through interactions with other business objects. This means that domain objects are aware of value objects and entities.
- They can interact with other domain services and repositories, as required.



### Domain service



Recall the e-commerce example that we worked on previously, where a customer buys some product. Suppose there is a use case where we are required to send an email to a customer after their bill is created. To fulfill this requirement, we need to create a domain service that combines customer information with product information. This definition shows how the use case needs to deal with the information of two business objects, which is why this logic does apply to any of them.



The model will be as follows:



![image-20220509164330944](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509164330944.png)

When the `Sells` object is instantiated, it creates all of the dependencies that it needs to execute its work. To test it, a `main` method creates a `Sell` object and executes its `GenerateBill()` method. Once this method is executed, it prints a message simulating that an email is sent.





## Summary

Domain services implement business logic that cannot be placed in other domain objects, because this logic may need to interact with more than one of those objects.









![image-20220509171345285](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509171345285.png)





![image-20220509171403835](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509171403835.png)





![image-20220509172656280](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509172656280.png)





