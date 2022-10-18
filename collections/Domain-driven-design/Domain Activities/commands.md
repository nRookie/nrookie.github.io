Get an introduction to the concept of commands, learn of their types, and understand how they change the state of a system.



Let us imagine a situation where a customer wants to execute an action in a bounded context, such as buy an object, transfer money, or watch a film.



![image-20220508144114382](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508144114382.png)



In this case, the system will receive a request and execute the action that was requested. The execution of this action can be successful, or it can fail. When it is successful, it will change the system state.



## What are commands?



Commands are the first of the three activities that can occur in a domain, and it is broadly used in a pattern called **Command Query Responsibility Segregation (CQRS)**. Commands represent the actions that are requested by an actor. When an actor requests to do something, this requested action is not executed in that moment. It is executed in the future, and it may or may not be carried out successfully.



### Types of commands



It is common to implement systems that are oriented to ***CRUD\***. These systems execute create, read, update, and delete operations.



In the case of commands, they are oriented to ***CUD\***. This means that they only perform the *create*, *update*, and *delete* operations.



All of these operations change the state of the system. Let us look at an example of such a change in the state of the system. Consider an application that a customer wants to sign up.





1. The system does not have registered people:

2. ![image-20220508144406670](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508144406670.png)

3. A 23 year-old user with the name DDD wants to sign up for the system. In this case, the system executes a **creation** command. The execution is successful and it creates a new record in the database. This is what a change of the state means. It is nothing but an alteration of some part of the system. In this case, the system is altered through the addition of a new record in the database.

4. ![image-20220508144505675](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508144505675.png)

5. Suppose a year passes by, and the user DDD wants to update their age. They will execute an **update** command. Similar to the command of creation, the update will be executed successfully. Hence, the state of the system will be changed. In this case, there will be an alteration of a record in the database.

6. ![image-20220508144554470](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508144554470.png)

7. Ultimately, the user DDD will want to end their subscription. Thus, they will execute a **delete** command. This will also be executed successfully, too. As a result, the system state will change again. In this case, a record will be deleted from the database.

8. ![image-20220508144636700](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508144636700.png)







### Commands and ubiquitous language

When a command is sent to a specific destination —for example, the transfers microservice or the payments microservice— the command’s handler should be given a clear name. The name of the handler should convey what the command wants to do, for instance, whether it wants to transfer money, pay a bill, create an order, and so on. It is important to note that, as per ubiquitous language, command names defined in a model must be translated into code. The following are examples of handler names, as they would appear when translated into code:



- **Transfer money**: the handler’s name should be `transferMoney()`.
- **Pay a bill**: The handler’s name should be `payBill()`.
- **Create an order**: The handler’s name should be `createOrder()`.



In this way, ubiquitous language works together with code.



![image-20220508145024197](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508145024197.png)



## Summary



A command represents what an actor wants to execute in a bounded context. It is possible to create, update, or delete something in a bounded context.

Additionally, the implementation must reflect the ubiquitous language that is defined in the bounded context





