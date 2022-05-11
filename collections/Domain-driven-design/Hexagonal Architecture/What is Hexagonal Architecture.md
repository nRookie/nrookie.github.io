Learn to structure a component with the use of an interesting approach called hexagonal architecture.



Once the whole business model is defined, at least for the first MVP, it is time to start coding the project. There are many ways to implement software in the software industry; some implementation methods are good, while others are not. . One point to keep in mind when software is implemented is that there is no silver bullet. This means there is no tool, practice, or framework that can be used in every project. Therefore, it is a good practice to analyze the problem and choose the best software architecture.



## What is software architecture?



This is not an easy term to define, because it can be abstract depending on the perspective. When software for a company is designed, there are three different levels to the process:

- **Strategic level**: This is known as ***enterprise architecture***. It deals with problems related to processes, applications, infrastructure, and actors that are involved across an organization.
- **Solution level**: This is known as ***solution architecture***. It solves everything related to a particular solution of software. The solution may involve different components, actors, and specific processes of an organization.
- **Artifact level**: This is known as ***software architecture***. It deals with problems related to a particular component, in an ecosystem of components, and structures it. It may involve layers, a programming language, design patterns, and so on.



In this case, efforts will be oriented to define how to design and implement a specific component, how we should build parts inside the component, and how those parts should interact amongst themselves. These are the concerns that software architecture deals with. It also gives us good practices, guidelines, and many more things.



### The levels of architectural abstraction

The following image shows the levels of architectural abstraction:



![image-20220509173206778](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220509173206778.png)



At the top level, which is the most abstract representation of software, there is a **system**. A system may consist of a set one or more **subsystems**. Subsystems are typically divided into one or more **layers**. In the same way, layers are often divided into different **components**. Similarly, components contain **classes** that, in turn, contain **data** and **methods**.



## What is hexagonal architecture?



Hexagonal architecture, also known as ports and adapters architecture, is an architectural pattern that separates the business logic from the outside world. It involves APIs exposure or database integration. To understand whether this pattern is well implemented or not,  one could ask themselves questions such as: What would happen if functionality was exposed through another protocol? Or what would happen if the database was changed? If the business logic was required to change too, then this would imply that the hexagonal architecture was poorly implemented. If not, then the pattern is proven to be implemented well. 



![image-20220509173636707](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220509173636707.png)



This pattern was invented by Alistair Cockburn, one of the fathers of the agile movement. As the image above shows, this pattern is represented by a hexagon where the business logic is placed in the center and the adapters are situated on the edges. As Alistair Cockburn says, the term “Hexagonal” comes from this shape. However, this shape does not mean anything in the context of the pattern’s implementation.

Hexagonal architecture can be implemented in DDD projects, because they complement each other. Although it is not required, it is a good practice to combine them because they aim for similar goals.





### Why is it good to invest in hexagonal architecture?



The first reason to invest in hexagonal architecture comes from a cost-benefit argument. Not only is it beneficial from a technical perspective, but also from a business point of view. One of the primary goals of an architect is to minimize the cost of the creation and maintenance of software and, simultaneously, maximize the architecture’s benefits to the business.



The main benefit of hexagonal architecture for the organization is in the form of money. Companies invest a lot of capital every month in software maintenance, which can encompass refactoring or even rebuilding an application. Hexagonal architecture reduces the cost of maintenance by separating different parts of the application. This clear separation of responsibilities and layers allows us to evolve or change functionalities more easily.



For example, when we want to change from REST to gRPC, we only need to focus our efforts on the transport layer while the rest of the code can remain intact.





## Summary



Hexagonal architecture is an architectural pattern, which provides us with guidelines to build components where the business logic is isolated from the technical logic.



