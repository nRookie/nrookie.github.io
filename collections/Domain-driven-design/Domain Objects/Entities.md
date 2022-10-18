# Entities



It can be difficult to model a domain. Once the whole model is done, the next step is to move on to the code. The best approach to translate the domain model into code is to use **Object-Oriented Programming (OOP)**. To implement this approach, DDD defines a number of artifacts that allow technical experts to represent a domain model in the code.



## What is Model Driven Design?



**Model Driven Design (MDD)** says that a software solution should be implemented based on the analysis of models. Remember that a model is an abstraction and representation of something from the real world that makes it easy for us to understand it. Once a domain is modeled, it is possible to begin the implementation of software.



Beyond the MDD methodology, the **Object Management Group (OMG)** created a conceptual framework called **Model Driven Architecture (MDA)**. The MDA approach provides guidelines for how to structure software specifications that are expressed as models. It covers the complete development lifecycle in the analysis, design, programming, testing, deployments, and maintenance stages. MDA uses three models:





1. **Computation Independent Model (CIM)**: This model focuses on the domain model.
2. **Platform Independent Model (PIM)**: This model focuses on a general platform, where a solution is deployed.
3. **Platform Specific Model (PSM)**: This model focuses on a specific platform, where a solution is deployed.





## What is an entity in the DDD world?



According to Model Driven Design, DDD uses tactical patterns to represent a model. There are three possible domain objects for the implementation of business logic: **entities**, **value object**s, and **aggregates**.



In the context of DDD, an **entity** represents something that is involved in a business process. For example, in the context of a bank, some entities may be Account, Credit Card, Customer, or Transaction. An entity shows well-defined attributes and domain behavior. Additionally, it is identified by a form of identification, such as an ID or a key, which implies that attribute values may change when required. This is why this domain object is mutable. The identification of an object is not affected when attribute values change.



![image-20220508163743878](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508163743878.png)





### The characteristics of an entity



Entities must meet the following characteristics to be well-defined:

- Each needs a unique identifier. An entity should be defined within a specific bounded context.
- They should be meaningful only in one bounded context.
- They should not use **setter** methods.
- They may use **getter** methods, when necessary.
- If some logic is required, value objects can acquire it.
- A constructor should be the only way to create an entity.



### An example of an entity

Let us assume that there is a need to make a transfer between two banks. In this case, there are some business rules that should be followed:



- The amount must be less than `10000`.
- The transfer date should be the present-day date.



The model is as follows:

![image-20220508164523129](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508164523129.png)

A `Transfer` entity is created to encapsulate the business logic. To test it, a `main` method is created. According to the `Transfer` constructor, the `main` method creates a `Transfer` object when it passes an `id`, the present-day `date`, an `amount`, and an `accountId`. After this, it executes the following methods:

- `IsAmountLimitAllowed()`: This will print `Amount is not allowed`, because the amount will be greater than `10000`.
- `IsDateAllowed()`: This will print `Date is allowed`, because the date that is passed will be the same as the current one.

Let’s run the code below to see this process in action.



