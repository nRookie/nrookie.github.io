It might be necessary to create a huge number of entities and value objects and facilitate interactions between them, when we deal with a complex bounded context. These conditions may lead to the creation of a big ball of mud. As was explained before, this is an anti-pattern where interactions among objects are unmanageable. It can be described as a spaghetti code.



A big ball of mud implementation tends to grow out of control and is difficult to maintain over time. DDD understands this problem and solves it through the definition of a tactical pattern called aggregate.



## What is an aggregate?



An **aggregate object** is a cluster of entities and value objects that are treated as a single unit from the domain and data perspective. This object acts as a load balancer for accessing a set of nodes. This means that an aggregate is the only access point for external objects. As the image below shows, an aggregate has clear bounds and the artifacts inside it cannot interact with the outside world:

![image-20220508172634658](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508172634658.png)



### Accessing functionalities to internal artifacts

An aggregate orchestrates the logic among all of the value objects and entities that are within its limits. When an object outside of the aggregate bounds wants to execute the business logic of entities or value objects inside the aggregate bounds, it should do this through that aggregate. Therefore, an aggregate should expose a type of interface that allows external objects to reach the functionalities inside it. To understand this behavior, let us look at an image where an artifact wants to execute a method located in an entity. The artifact must pass through an aggregate:



![image-20220508175557438](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508175557438.png)



### Changing information with aggregates

The set of an aggregate, its entities, and value objects should be treated as a single unit. Information related to these artifacts must travel back and forth between the bounded context and its repository, as a whole. It should be impossible to process a piece of information related to an entity or value objects, separately. An aggregate related to the aggregate example item is as follows. It comprises a main object called `transaction` and two main attributes. These attributes are the `amount` value object and the `transfer` entity:



``` json
{
    "eventId" : "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
    "eventType" : "Transfers",
    "eventName": "moneyTransfered",
    "timestamp" : "1628360557000",
    "data" : {
        "transaction":{
            "id": "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
            "bankName": "Bank test",
            "amount":{
                "value": 76567.78,
                "Currency": "USD"
            },
            "transfer":{
                "id": "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
                "date": "2021-08-07T17:10:43+00:00",
                "accountId": "09876543"
            }
        }
    }
  }
```



### An example of an aggregate



In continuity with the transfer example, let us recall the following business rules:

- The amount must be less than `10000`.
- The transfer date must be the same as the current date.



In this case, an aggregate called `Transaction` is introduced to put objects together.



![image-20220508175737359](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508175737359.png)





Three objects are created: an entity called `Transfer`, a value object called `Amount`, and an Aggregate called `Transactions`. The last one is in charge of the execution of a transfer. The `main` method passes an `id`, a `bankName`, an `amount`, the chosen `currency`, the `current date`, and an `accountId`. After this, the `TransferMoney()` is called to execute the transfer. Let’s run the code below to see this in action.



## Summary

Aggregates are a collection of domain objects that are treated as a single object.
