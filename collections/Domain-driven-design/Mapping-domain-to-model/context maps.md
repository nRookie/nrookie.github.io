Learn why it is important to understand the relationships between different bounded contexts, what types of bounded contexts exist, and how they can be mapped onto a context map.



Bounded contexts are independent, but they do not work in isolation. They interact with each other to fulfill business requirements. Since a team usually works on a big domain, the team can end up with a huge number of bounded contexts. This is worth mentioning because every bounded context is translated into one or more modules or microservices in the future. It is difficult to understand the relationship between different microservices. If incorrectly implemented, our efficient architecture can devolve into a big ball of mud anti-pattern.



## Big ball of mud anti-pattern



When there are many bounded contexts in an organization, it is important to manage the relationships between them. Unmanaged bounded context relationships lead to the formation of a big ball of mud. This is when a software solution possesses multiple haphazard models that do not work with explicit boundaries.A big ball of mud architecture can cause potential problems, such as the creation of a spaghetti code, whereby a team will need to connect components improperly. The evolution of components and a model will be difficult and expensive.



> **Spaghetti code:** A slang term for unstructured and difficult-to-maintain source code. Spaghetti code can be caused by several factors, such as volatile project requirements, lack of programming style rules, and software engineers with insufficient ability or experience.



![image-20220507223235885](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507223235885.png)



## What are context maps?

The aim of mapping the relationships that exist among bounded contexts is to understand the dependencies between them and how they communicate with each other. **Context maps** help to do just that.



They are a visual representation of the relationships that are present between different bounded contexts. They show the interactions and the types of connections that bounded contexts share. Context maps help to maintain clear communication between bounded contexts, as it is clear what types of interactions occur during a business process. This is why the addition or removal of components is less complicated.



### Types of relationships among bounded contexts 



The four main types of interactions that can appear in relationships among bounded contexts are:



1. **Separate ways**: Consider a situation where there is no need to connect two bounded contexts, because that interaction will not give significant payoffs to a team. In this situation, it is possible for these bounded contexts to implement their own logic and operate in a separate way.
2. **Symmetric relationship**: This is known as bidirectional dependency. This relationship occurs when there is communication between two bounded contexts in both directions. This scenario is common in orchestrations where there is a component A that invokes component B. After an execution of something, component B invokes component A to convey the result.
3. **Asymmetric relationship**: This is known as unidirectional dependency. This relationship occurs when there is communication between two bounded contexts in only one direction. It is common to encounter this type of communication in connections by queues, where component A sends a message to component B through a queue and does not wait for a response.
4. **One-to-many relationship**: This relationship occurs when there is communication between two or more bounded contexts. This type of relationship means that a bounded context can connect and send messages to one or more bounded contexts. It is common to come across this type of relationship on implementations based on a publish-subscribe pattern, where component A sends a message to a broker, and the broker sends a copy of that message to all of the subscribers.





![image-20220507224051504](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507224051504.png)



The image above shows the relationships that exist between these four bounded contexts:



- `Customer` is connected with `Credit Card` and `Movements`: This is a **one-to-many** relationship.
- `Credit Card` is connected with `Customer` and vice versa: This is a **symmetric** relationship.
- `Notifications` are connected with `Customer`: This is an **asymmetric** relationship.
- There is no connection between `Credit Card` and `Notifications`: This is a **separate-ways** relationship.





It is possible to label relationships to give information about them. This information can be with regards to the nature of a relationship, for example.



### Benefits of contexts maps 





Mapping the interactions between bounded contexts brings us the following benefits:



- It makes it easier for us to understand the bigger picture of the components.
- It helps us determine the level of collaboration that is required between teams.
- It helps us understand the interactions between components.
- It helps us clarify bounded contexts and models.



### Summary

Context maps help us visualize the relationships and types of communication that take place between bounded contexts.



