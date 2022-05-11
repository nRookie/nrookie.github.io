Continuing with the restaurant example used above, let us implement the use case layer on top of the domain layer.



## Example definition



Let’s continue with the example defined in the [Domain Layer lesson](https://www.educative.io/pageeditor/10370001/4616975235416064/6471623718207488#Example-definition) and implement the use case layer. This layer will be in charge of the orchestration of domain objects to fulfill business needs.



### Model definition

The following diagram depicts the use case layer, which defines two artifacts. Additionally, this layer makes use of the domain layer.



<img src="/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510164446848.png" alt="image-20220510164446848" style="zoom:200%;" />



### Model explanation

This layer contains the following artifacts:

- A `CreateOrder`interface: This interface defines a method to carry out an order creation.
- A `CreateOrderImpl`object: This object implements the `CreateOrder` interface. It orchestrates the logic to look up information about a customer and a product.





## Summary

The use case layer contains all of the logic related to business requirements, and information on how use case objects should interact with aggregates to fulfill that which the business experts define.