Learn about the main aggregate in a domain layer, called an aggregate root, and how to implement it





Let us imagine that there are many aggregates in the domain. They have different functionalities, for several use cases. When it is necessary to carry out a particular use case, it will be difficult to choose the access and decide which aggregate should be executed first. Therefore, DDD extends the concept of aggregate and creates a new artifact called an **aggregate root**.



## What is an aggregate root?



Like a normal aggregate, an aggregate root is a cluster of objects with entities, values, or even other aggregates. The main difference between them is that an aggregate root is the main aggregate, while a simple aggregate is not the main aggregate. This means that everything outside of the domain layer boundary must use business logic, through interaction with the aggregate root. The aggregate root should, in turn, orchestrate logic in other aggregates, value objects, and entities.



![image-20220509103528030](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220509103528030.png)





### How can we choose an aggregate root?



The selection of an aggregate root is not a trivial task, but we can make it easier by thinking about the context of our project. To choose the best aggregate root, we should consider the following suggestions:



1. An aggregate root should always be an entity, as it requires an identifier.
2. An aggregate root can be different depending on the context. In **bounded context A**, the aggregate root might be **component A**, but in **bounded context B**, the aggregate root might be **Component B**.
3. ![image-20220509103720255](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220509103720255.png)
4. Objects within an aggregate root can hold references to multiple aggregates.
5. Aggregate roots own a global identity. Artifacts inside the aggregate root show local identities. Nonetheless, the most important identity is that of the root key.
6. It is important to consider the following questions:
   1. If the entity is deleted, is it necessary to delete the other entities?
   2. Does a single transaction span across multiple entities and value objects?
   3. Is the entity involved in multiple functionalities inside the same aggregate root?
   4. Does the entity execute more functionalities than other entities?





### An example of an aggregate root



With regards to the example of a transfer that was analyzed before, let us consider the next business rules:

- The amount must be less than `10000`.
- The transfer date must be the same as the current date. In this case, the aggregate root will be the entity called `Transaction`, and it will contain some entities, value objects, and aggregates. The model will be as follows:



![image-20220509104049974](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220509104049974.png)



The JSON representation is shown below. It shows the main object called `transaction`, which contains some sub representations:



``` json
{
    "eventId": "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
    "eventType": "Transfers",
    "eventName": "moneyTransfered",
    "timestamp": "1628360557000",
    "data": {
        "transaction": {
            "id": "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
            "bankName": "Bank test",
            "customer": {
                "id": "8867g87m",
                "firstName": "DDD",
                "lastName": "test"
            },
            "amount": {
                "value": 76567.78,
                "Currency": "USD"
            },
            "transfer": {
                "id": "1ce5608e-c76c-4412-8bd1-a2c6ed42970d",
                "date": "2021-08-07T17:10:43+00:00",
                "account": {
                    "accountId": "0098976",
                    "balance": 7675.9,
                    "nickname": "test-account"
                }
            }
        }
    }
}
```



There are some objects that are created to execute a transfer. The `main` method creates a `Transaction` object and passes the `id`, a `bankName`, the `amount`, the `currency`, the `current date`, and an `accountId`. After this, the `TransferMoney()` function is called to execute the transfer. The most important point to note here is that `Transaction` is the aggregate root.





## Summary

An aggregate root is a cluster of entities, value objects, and other aggregates, which orchestrates every use case in the domain layer.







An aggregate root is an aggregate as such. It is a cluster of entities and value objects.



A cluster of artifacts, such as entities and values objects, may be considered an aggregate root.



![image-20220509105740944](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220509105740944.png)





Which of the following are the models of Model Driven Architecture? (All)

- CIM (Computation Independent Model),

* PIM (Platform Independent Model),
* and PSM (Platform Specific Model)



![image-20220509105947528](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220509105947528.png)





