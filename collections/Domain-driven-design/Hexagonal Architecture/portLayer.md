We will continue with the example defined in the [Domain Layer lesson](https://www.educative.io/pageeditor/10370001/4616975235416064/6471623718207488#Example-definition) and implement the port layer on top of the use case layer. It is important to mention that this layer will be in charge of exposing an interface, which will define the functionalities that a component offers and how we can invoke them.



### Model definition

The following diagram depicts the port layer, which defines four artifacts. Additionally, this layer makes use of the use case layer.



![image-20220510165911672](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510165911672.png)



### Model explanation

This layer contains the following artifacts:

- A `CreateOrderPort`interface: This interface defines a method to carry out an order creation.
- A `CreateOrderPortImpl`object: This object implements the `CreateOrderPort` interface. It invokes the `CreateOrder` function, located in the use case layer.
- An `InputDTO` object: This object defines the data that is required to carry out an order creation.
- An `OutputDTO` object: This object defines the data that is generated after an order creation is executed.





### Implementation









https://www.c-sharpcorner.com/UploadFile/ff2f08/association-aggregation-and-composition/





![image-20220510173007559](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220510173007559.png)
