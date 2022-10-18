As was shown in the [previous lesson](https://www.educative.io/collection/page/4927079282376704/4708195672522752/5718229709750272), DDD uses tactical patterns to represent a model. Consider the situation in which we need to create an object that does not need identification, such as an amount with value and currency. Value objects are useful in such cases.



Before going further with the definition of value objects, it is worth looking at a common situation of the real world. This example, taken from Eric Evans book on DDD, describes what a value object is from the real-world perspective:



*“When a child is drawing, he cares about the color of the marker he chooses, and he may care about the sharpness of the tip. But if there are two markers of the same color and shape, he probably won’t care which one he uses. If a marker is lost and replaced by another of the same color from a new pack, he can resume his work unconcerned about the switch.”*



## What is a value object?

A **value object** is a representation of something that is involved in a business process. Unlike entities, value objects are defined through their attributes rather than an identifier. This means that value objects do not have an ID, and they are identifiable only by their set of attributes. Hence, these objects are immutable. Some common examples of value objects are `Address`, `emailAddress`, `Amount`, and so on.



![image-20220508170608729](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508170608729.png)



### Characteristics of a value object

Value objects should meet the following characteristics to be well-defined:



- They should be defined within a specific bounded context.
- A value object in one bounded context may be an entity.
- They should be meaningful only in one bounded context.
- They should have any logic that is required for their use.
- They should not use **setter** methods.
- They may use **getter** methods, when necessary.
- A constructor should be the only way to create a value object.





## Example of a value object



Let us continue with the [example](https://www.educative.io/collection/page/4927079282376704/4708195672522752/5718229709750272#entity-example) that was given in the previous lesson, in which we needed to make a transfer between two banks. There are some business rules that must be adhered to during this transaction:



- The amount must be less than `10000`.
- The transfer date should be the same as the current date.





It is possible to create a value object, called `Amount`, that deals with the logic related to `amount`. Hence, the new model could turn out like this:



![image-20220508170857670](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Domain Objects/image-20220508170857670.png)



A `Transfer` entity is created to encapsulate the business logic. Additionally, an `Amount` value object is created to control the logic related to amount. To test it, a `main` method is created. Functionality is the same as in the [previous lesson](https://www.educative.io/pageeditor/10370001/4616975235416064/5125176678678528) but in this case, the amount is validated in the `Amount` object. The `main` method creates a `Transfer` object when it passes an `id`, the present-day `date`, an `amount`, and an `accountId`. After this, it executes the following methods:





- `IsDateAllowed()`: This method prints `Date is allowed`, because the date that was passed is the same as the current one.
- `IsAmountLimitAllowed()`: This method prints `Amount is not allowed`, because the amount is greater than `10000`.





## Summary

Value objects in DDD are domain objects that do not have an identifier. They may contain attributes and business logic. These attributes, together, become their source of identification.
