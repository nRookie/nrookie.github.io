Learn what a shared kernel is and how it enables bounded contexts to share their functionalities.



Let us assume that a team defines microservices based on the bounded contexts they find. As they start to build these services, they notice that many of them share the same logic.



During the construction of components such as microservices, we deal with concerns of logging, security, communication, and so on. These concerns are the common logic that can be shared between those components.



![image-20220507230824653](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507230824653.png)



he space of this problem is shown in the image above. It depicts how two components that are evolved by two different teams share some logic (red lines). They have their own logic, but these two teams might create redundant implementations at some point, or they could overlap implementations. This is an expensive operation for an organization, because they need a high budget to maintain two or more teams that perform similar tasks.



## What is a shared kernel? 



The shared kernel is a strategic pattern, which tackles the problem that was mentioned previously. It is important to clarify that the functionalities that should be shared between components are the technical ones. The implementation process of components based on DDD, consists of two parts. The first and most important is the core business logic. The second part comprises all the things that are built to support the core business logic, such as database connections, API implementations, and so on.



Let us explore this concept in more detail, with an example. Imagine there are two teams that build two microservices, “Payments” and “Customers”. As was mentioned before, they have their business and technical implementation:



![image-20220507231001256](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507231001256.png)





As the image above shows, “Logs” and “REST” are common implementations in both microservices. What should each microservice team do to avoid redundant implementations?



![image-20220507231020692](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507231020692.png)



They should implement a shared kernel pattern. The image below shows an external library, which should be included in every microservice.



![image-20220507231037313](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507231037313.png)



### Benefits of the shared kernel pattern 



This shared library gives several benefits to teams and organizations. The most noticeable benefits are:



- Reutilization of code, as this library can be included in many components.
- Streamlining the development process, because a lot of common functionalities are developed when a team wants to develop a new component.
- Government over components, because everyone controls changes.



### Drawbacks of the shared kernel pattern



This extra library increases complexity, in terms of management and development:

- There must be regular communication among all of the teams to evolve the shared kernel. Sometimes, this communication is difficult.
- The shared kernel should be as small as possible, because it could slow down the development process if it becomes too big.
- It is important to guarantee backward compatibility. It can cause failures in microservices if it is not guaranteed.





## Summary 

A shared kernel contains functionalities that are common across two or more bounded contexts. It avoids the repetition of code in several places and helps the projects evolve quickly over time.