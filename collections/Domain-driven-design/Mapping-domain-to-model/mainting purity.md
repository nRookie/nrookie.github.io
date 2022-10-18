After a team defines its bounded contexts, modules, or microservices, the team proceeds to implement interactions between them. Now, it is time for us to think about how these interactions will occur.

Purity is all about maintaining a model away from something that may mess it up. For example, models of external applications, databases, legacy systems, and even models of other bounded contexts.

Let us imagine that there is a component called transfers that needs to obtain information from a third-party application. This external application does not follow the model of transfers that was created by the DDD team. In this case, it is necessary to create a coupling between transfers and the external service. However, it is crucial to avoid any type of coupling.





## What is an anti-corruption layer?





The **anti-corruption layer (ACL)** is a strategic pattern in the DDD world. It is nothing but a translator, which translates a bounded context model into an external model and vice versa. One of the most important things to consider during the implementation of components, based on DDD, is to avoid corrupting business logic. Once the core business logic is implemented, it will need to interact with the external components. It is impossible to implement any external connection from the domain logic. The anti-corruption layer exists to help us implement this external connection. This layer can be implemented by using some design patterns, such as Facade and Adapter, or even by coding a custom translation.



![image-20220507224711067](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507224711067.png)



The image above shows how an anti-corruption layer should be implemented inside the limits of the “Transfers” bounded context. The anti-corruption layer allows us to interact with external and legacy systems without messing up the “Transfers” model. In this way, when a third-party or legacy system changes, only the anti-corruption layer will need to be changed. The “Transfers” model does not need to change in any way.



Not only does the anti-corruption layer help to avoid the messing up of a model, but it also isolates internal and external changes. Let us imagine that there is a service A, which consumes an external service. The external service returns a specific structure of the message, as is shown in the image below.



![image-20220507230639822](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507230639822.png)



Now, the external service changes the message with the removal of a field. This change will cause problems when the external service is consumed. It is necessary to make a change in service A.



![image-20220507230658415](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507230658415.png)



In this case, the core business logic remains in the same state. The anti-corruption layer will be the only layer that needs reimplementation.



## Summary



The anti-corruption layer helps to maintain the purity of our bounded context. It is important to implement some pattern, like a Facade or an Adapter, to translate models between a bounded context and an external application.



