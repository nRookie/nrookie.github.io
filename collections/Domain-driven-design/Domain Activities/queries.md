A user may want to know what happens in the system when something is executed in a domain and an event is produced. Queries enable the user to receive information from a system, with regards to actions and events that occur in a domain.



![image-20220508154331886](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508154331886.png)

## What are queries?



Queries are third in order of the three activities that can occur in a domain, and it is part of the **Command Query Responsibility Segregation (CQRS) pattern**. They represent the requests that a user makes to receive information about a domain. Typically, a system exposes some interface, which enables us to look for information. A query should always return a `h` response with information when it exists.



It is common to implement systems oriented to **CRUD**. They execute the create, read, update, and delete operations. In the case of queries, they are oriented to Read. The system state should not change when a query is executed.

### Queries are idempotent

Queries should be idempotent. This means that queries should not change the state of the system. In other words, no matter how many times a user requests information, the system should always return the same information. The responses must be the same, as long as there are no changes in between the queries. When a command is executed between two queries, it will cause the second query to return a different response. Let us understand this behavior with an example.



The following image shows how two queries are exactly the same when there are no changes in the state between them. Additionally, they do not change the state either:



![image-20220508154707958](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508154707958.png)

On the other hand, when there is a change of state between two queries, the result is different:

![image-20220508154754588](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508154754588.png)





### Queries and ubiquitous language 

Ultimately, queries are defined based on the ubiquitous language pattern. This implies that they should be named in accordance with business terms, and the code should reflect these terms. Some examples of queries are: *checked if payment was executed*, *get user information*, or *get movements*.



![image-20220508154843365](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508154843365.png)





## Summary

Queries allow an actor to receive information from a bounded context. Once a query is executed, it should not change the state of the bounded context.



