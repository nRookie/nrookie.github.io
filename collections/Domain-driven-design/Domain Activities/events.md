A domain can generate some results when executing commands. Those results are known as events. When an event is produced as a result of executing a command, it can be processed in different ways. For example, it can be stored in an append-only system or trigger events in other domains.



## What are events?

An event is something that has happened in the past. They usually cause a reaction elsewhere. It is important to clarify that events are different from each other. It means that although two events have been created based on the same information, they are completely different. In other words, events cannot be duplicated and cannot be modified.



### Events are immutable

When something happens in the real world, it is impossible for us to change that action, because it immediately becomes a part of the past. This principle of the real world is also applicable to our model behaviour in a software application.



Domain events are immutable because they have already occurred. We can save everything that happens, but we cannot change it. As per the real-world, nothing can be modified and these events produce reactions elsewhere.



Domain events are immutable because they have already occurred. We can save everything that happens, but we cannot change it. As per the real-world, nothing can be modified and these events produce reactions elsewhere.



## What are domain events?

Domain events are the second of the three activities that can occur in a domain. Domain events describe the actions or events that occur in the domain. They are important to domain experts. Therefore, they are built based on ubiquitous language. A domain event is represented by any notation that is useful to create it. Such a notation can be **JSON** or **YAML**.



Let us assume that a customer of a bank wants to transfer money. Once the command is sent and the bank’s system processes it, a domain event is produced and saved in the system.



![image-20220508153123007](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508153123007.png)



As this image shows, after the execution of the request to transfer money, a domain event called moneyTransferred is created and saved in an event log. An event log is a system, which stores the events that happen in a system and guarantees their immutability. A saved domain event will look like this:





### How can we define domain events?



When defining domain events that happen in a domain, it is important to consider some important guidelines:



1. Events must be written in the past tense. For example, `billPaid`, `moneyTransferred`, or `orderCreated`.
2. Ubiquitous language must be used as the basis when domain events are named.
3. Domain events should be defined and built, because they are required to execute something else. According to the **YAGNI** principle, which stands for “You Aren’t Gonna Need It”, someone should always implement things when they are needed, never when it is presumed that they will be needed. In terms of DDD, this means that domain events should be defined and implemented only when they will support some business process.
4. Domain events can be emitted to or received from a system, from either a logging system, a billing system, or an external system. A consideration of these functionalities permits us to predict whether or not the construction of a domain event is worth it to a business. Interactions by domain events can be detected in an event storming session. A very basic interaction between some bounded contexts is shown in the image below —where *“bounded context A”* shares the same domain event with *“bounded context B”* and *“bounded context C”*.
5. ![image-20220508154007220](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Activities/image-20220508154007220.png)



### How can we detect a domain while talking to domain experts?

As we saw in a previous lesson, [event storming](https://www.educative.io/collection/page/4927079282376704/4708195672522752/4589670530285568) is relevant when extracting information from business experts. In this session, it is possible to come across phrases like *“notify the user when”*, *“inform the user when”*, *“whenever this happens, then”*. These phrases help us discover which events happen in our domain. They inform us when events are relevant to a system, through avoidance of the YAGNI principle.



## Summary



 Domain events are another type of activity that can happen in a domain. They represent actions that were already executed in a domain. Therefore, they cannot be modified.



