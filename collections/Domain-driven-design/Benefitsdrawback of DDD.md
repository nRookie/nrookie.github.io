## Benefits of DDD



### Flexibility



When a project is based on DDD, it is easier to evolve and change things like business processes, implementations, or technological stacks, which gives us a better time to market.





### Code is organized



Domain Driven Design can work together with hexagonal architecture. **Hexagonal architecture** aims to organize the code of an application into a specific structure. It helps developers to understand and modify code more easily.



### Business logic lives in one place

When a code is organized on the basis of hexagonal architecture, there is a layer where all of the business objects are placed. This layer allows for a team to change business logic more easily, because it only needs to concentrate on one layer. The yellow hexagon in the image below represents the logic that is defined through the analysis of a business process. It is easy to evolve code in this layer whenever there is a change in the business rules. The business logic code must not spread out over to other layers.



![image-20220506225136993](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/image-20220506225136993.png)





### Relationship between business people and the development team



The process of modeling a domain requires regular communication between two worlds. The code used in this case must be written based on business language.



### The domain is more important than user interface (UI) and user experience (UX)



As the domain is the main concern, a technical team will not build software that focuses on the visual layer. Software-based on DDD always has its backbone in the domain.



## Drawbacks of DDD



### The learning curve is large



DDD includes many principles, patterns, and processes, which make it difficult to comprehend. This is why any team that wants to build software based on DDD should possess good knowledge and expertise in this practice. Otherwise, this task will be extremely difficult for them.



### Time and effort



Most of the time on the project can be spent in conversation with domain experts to understand and model business logic.



### DDD cannot be used on every project



As was mentioned previously, DDD only works on complex domains. Not only is it complex in technical terms but, more importantly, it is complicated in business terms. DDD is not the best approach for simple or small domains and projects.



![image-20220506225500263](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/image-20220506225500263.png)



The image above shows how a big domain may be divided into several subdomains and that there might be numerous actors involved in its processes. Here is when Domain Driven Design makes sense:



![image-20220506225522225](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/image-20220506225522225.png)



When there is a small domain with no complexity, Domain Driven Design will overcomplicate the development process.







Domain Driven Design gives technical and business experts many interesting benefits when software is implemented. Nonetheless, it is important for teams to keep in mind some of the drawbacks that DDD might present.



