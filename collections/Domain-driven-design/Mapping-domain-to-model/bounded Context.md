It is almost impossible to create only one model that represents the complete operation of an organization when the processes of an organization are modeled.   The model becomes difficult to understand and translate into code.   Different groups of people can use subtly different terms for different subdomains in a large company. This leads to confusion and misunderstanding of vocabulary. Domain Driven Design proposes that we explore each subdomain separately. This will allow a team to get the best definition of vocabulary, used in a particular subdomain. Therefore, the team will be able to create a more accurate model for a subdomain.



## What is a bounded context?

DDD deals with the problems that the common domain models experience, the domain into independent parts called bounded contexts.   A **bounded context** is a strategic pattern, which helps to maintain consistency between subdomains and their models. There must be a clear definition of limits between all of the bounded contexts that exist in a domain. These limits can be set during scoping sessions, such as an event storming.



![image-20220507222000545](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507222000545.png)



The image above shows two contexts, each with their own model. There is a common model between these two contexts. That common model is the “Credits” model.  However, it is completely different because it uses the concepts we learned in the [Ubiquitous Language](https://www.educative.io/collection/page/4927079282376704/4708195672522752/6417035921195008) lesson. there is a clear separation of concerns and models between the two bounded contexts. That is what bounded context solves.



We should avoid an abstraction of generic things when a process or software is modeled. This means that things, such as actors and entities, are different between subdomains. For example, data that is needed for a customer in subdomain A may not be the same as the data that is needed in subdomain B. Such an instance is perfectly fine and normal.



### Characteristics of bounded contexts

When bounded contexts are defined, a team should take the following characteristics into account to confirm that they are designing bounded context well:

- Each bounded context should possess its own domain model. As shown in the image above, each subdomain represents its structure and the way that every part interacts with another.
- Each bounded context uses its own ubiquitous language. There is a clear separation of vocabulary that exists inside each subdomain in a large organization. It guarantees a clear definition of responsibilities and concerns that an actor may encounter.
- A domain model built for a bounded context is only suitable within its boundaries.





### Difference between subdomains and bounded contexts

The confusion between these two terms stems from the fact that they correspond to each other. However, they are not the same concept. On the one hand, a subdomain is the context of the problem. On the other hand, a bounded context is the context of the solution. In other words, subdomains describe how our business is broken down, whereas bounded contexts describe how the software is broken down.



![image-20220507222611822](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507222611822.png)

The image above depicts how an organization is broken up into small parts called subdomains, where each subdomain has its own processes. This represents the **problem space**, where we can represent the problems in the domain that we want to solve. It is an abstract definition of the parts that a company is composed of. The image below shows how software components are defined to meet domain problems. This is the **solution space**. Solution space models are built based on abstract problem space models, like the image above





![image-20220507222737712](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507222737712.png)





## Summary

Bounded contexts is an important strategic pattern in the DDD world. They work to establish a clear separation of models and concerns across functionalities by using ubiquitous language.





