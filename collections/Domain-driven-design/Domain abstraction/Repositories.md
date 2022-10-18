Learn how to isolate business logic from data-access logic, through the implementation of a repository pattern in the DDD world.



Sometimes, it is necessary to interact with any type of persistence system in an application. Such a persistence system may be a **SQL** or **NoSQL** database, a file, or even an external service. There should be a mechanism that allows us to interact with these systems. The important point here is that the business logic remains separate from the repository-connections logic.



## What are repositories?



A repository does not necessarily involve a database. A **repository** is a pattern, which isolates business logic from data-store-interaction logic. It functions as a collection of objects in memory. It conceals the storage-level details needed for management and query of the information of an aggregate in the underlying data tier. When an aggregate wants to interact with information outside of itself, it invokes functions in a repository. It only knows a group of objects in memory. It is unaware of what lies behind such a repository.



![image-20220509110316937](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509110316937.png)



**One thing that is worth a mention here is that a domain layer is persistence ignorant. This means that business objects do not have any implementation related to how data is retrieved from a data store and how it is stored.**



### Characteristics of a repository



A repository should display the following characteristics to be a good implementation:



- An aggregate should only interact with one repository. Similarly, a repository must only interact with one aggregate. The image below shows how every aggregate is paired with its own repository.
- ![image-20220509110449965](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509110449965.png)
- A repository interface should only expose business-logic behavior. Functions and their attributes should be named in relation to their domain behavior. As the image below shows, the function, `getAccount(Account)`, is named in non-technical terms and receives the artifact account. On the other hand, the wrong definition defines a technical function, `getAccountById(id)`, and receives an ID.
- Persistence operations should be atomic. If the user needs to store a new state of a domain, it should be guaranteed as every stored aggregate acquires a new state. This behavior is depicted in the image shown below.
- ![image-20220509110836834](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509110836834.png)

### Benefits of repositories

The implementation of repositories give us the following benefits:

- It helps in the isolation of the domain model from the storage layer. It enables us to change storage technology, without affecting the business logic.
- It gives us a flexible way to implement unit tests, as it is possible to use a mock-storage system.
- It promotes the separation of concerns among the domain layer and the data-access layer.



### Drawbacks of repositories

It is important to consider the following drawbacks when repositories are implemented:

- It is possible to experience some performance problems when the aggregate is large. Multiple database operations may be required to store or receive objects in a large aggregate.
- Repositories may lead to criteria-based query problems. As was previously mentioned, functions in a repository should be defined with business terms. When a specific inquiry is required, this implementation may be complicated.







### An example of a repository

Let us assume that a customer wants to buy through an e-commerce platform. First of all, the customer will search for the products on the e-commerce website. To return information, an aggregate root called `Sells` will make use of a repository to retrieve all of the relevant products. The repository implementation will be a mock database.



The model will be as follows:



![image-20220509111556626](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509111556626.png)



The `main method` will be created to list the products with the `ListProducts()` method, declared in the `Sells` object. To instantiate the `Sells object`, it is important to create an object that implements `productRepository`. Once `ListProducts()` is executed, it will return the products.

Let’s run the code below, to see this in action.





## Summary

A repository is a tactical pattern, which addresses the isolation between domain logic and data-access logic.







### Glossory

The tactical patterns are **applied within a single bounded context**. In a microservices architecture, we are particularly interested in the entity and aggregate patterns. Applying these patterns will help us to identify natural boundaries for the services in our application (see the next article in this series).





DDD strategic patterns are **used to design abstractions of Business domain models incorporating behavior and data**.
