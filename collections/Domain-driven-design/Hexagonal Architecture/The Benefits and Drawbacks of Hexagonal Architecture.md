This lesson will cover the benefits and drawbacks that people may come across when they implement hexagonal architecture. These will lead to better trade-offs, which will result in the best outcome for companies and projects



## Benefits of hexagonal architecture



The benefits of hexagonal architecture are as follows:

- Hexagonal architecture has a well-defined dependency structure, which results in a clear domain model implementation.
- It emphasizes the domain logic, which is a good match within a DDD context.
- It clearly defines what to put in the code and where. That definition is important to maintain the code. It helps in faster, more focused, and automated tests for domain logic, mocking databases, and other external services.
- ![image-20220510183625731](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510183625731.png)
- An independent and incremental evolution of concerns is possible with hexagonal architecture. This means that every layer can evolve independently.
- It allows the domain model to evolve to fulfill business requirements, without breaking APIs or migrating a database on every refactoring.





## Drawbacks of hexagonal architecture



- The domain layer tends to become huge, as it can contain many objects if there is no good analysis of bounded contexts.
- With several levels of indirection and isolation, the cost of building and maintaining the application may increase. Therefore, applications with hexagonal architecture can become more complex.
- Complexity may increase, since applications are built with different levels of abstraction.
- Performance may be affected, as a request must pass through different layers and it may imply more classes than the usual.

![image-20220510183857725](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510183857725.png)

- Hexagonal architecture will overcomplicate a project if changing the database regularly or exposing functionality through different protocols is not a requirement.



## Summary



To take advantage of hexagonal architecture, it is important to know that it gives us a list of benefits that come at the cost of some drawbacks.



![image-20220510184007302](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510184007302.png)

![image-20220510184034765](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510184034765.png)





![image-20220510184104169](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510184104169.png)





![image-20220510184142213](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510184142213.png)
