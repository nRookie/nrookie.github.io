Learn what a factory is and how we can implement it to create complex entities in DDD.



As Eric Evans mentions in his book, there may be occasions when an object’s creation and assembly are complex. Such a process may need to orchestrate a lot of logic to create the complex object. With regards to this, he mentions an example related to the opening of a bank account. This process requires many objects in order to be executed properly. Oftentimes, the logic required to create those objects is not related to business logic, which can mess up the business logic. Refering to the Eric Evan’s example, let us imagine that the aggregate root of the bank account process is `Account`. Its attributes are a `customer` entity, which in turn has an `address` value object and a `bankOffice` entity, which in turn has an `address` value object too. Offices addresses exist in a repository, for this reason, it is required to use a repository called `officeAddress`. To build an `Account` entity is necessary to load information from external sources, validate the customer information. Interacting with those objects is what adds complexity to objects’ creation. To deal with this situation, DDD suggests that we implement a factory.

## What is a factory?



A **factory** is a tactical pattern used in the DDD world. It helps us create complex objects. It is important to keep in mind that we should only implement this pattern when the instantiation of an object is complex. The use of a factory may add unnecessary complexity.



Despite the help that it offers in the creation of aggregates, a factory does not need to be present inside the domain logic in every aggregate. A factory should only be present inside the aggregate root, as this is the object that orchestrates all of the business logic.





![image-20220509114300210](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509114300210.png)



In the context of DDD, the use of a factory pattern is not required to implement the factory design pattern. It is possible to implement any design pattern, such as a Builder or even a custom implementation. The only thing necessary for the factory design pattern to isolate business logic from the creation of complex objects.



### When is it appropriate to implement a factory?



As was mentioned before a factory is not always the best solution. However, it is a suitable approach when:

- Business logic requires the instantiation of many objects. Perhaps, in this context, a factory pattern is required to look for information in some external systems.
- The aggregate root input data is extensive and variable. We can use the data in different use cases. In this scenario, a builder pattern would make perfect sense.
- The aggregates, entities, or value objects that need to be instantiated vary in accordance with the use case.
- It is necessary to translate one bounded context into another.



### An example of a factory



Let us continue with the example that we looked at in the lesson on repositories. It was about listing all of the products on an e-commerce website. Now, the use case is to generate a bill. We will assume that something was already sold. The method that generates the bill requires a custom function, which instantiates `Customer` and `Product` to be executed.



![image-20220509115116454](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain abstraction/image-20220509115116454.png)



When the `Sells` object is instantiated, it creates all of the dependencies it needs to execute its work. To test it, a `main` method creates a `Sell` object and executes the `GenerateBill()` method. This will print all of the information related to a bill.



Let’s run the code below to see this in action.