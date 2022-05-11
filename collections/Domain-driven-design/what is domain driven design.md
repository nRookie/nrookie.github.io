

When a company acquires big projects, which are complex in terms of analysis and implementation, it can encounter problems. These problems might be caused if the company misunderstands business processes or develops a software solution that will not satisfy business requirements.



## What are the goals of DDD? 



The main goals that DDD seeks to satisfy originate from the book *Domain-Driven Design: Tackling Complexity in the Heart of Software*, written by Eric Evans in 2003. As its name suggests, DDD is designed to tackle complex problems rather than smaller, less complex ones. When it comes to the process of modeling small domains, DDD is not the best approach to use.



One of the book’s main goals is to outline how to develop software based on an evolving business model. This model is a representation of the structure, activities, processes, actors, and interactions that there are in a company. The different parts of a company can be depicted with the use of diagrams. For example, it is possible to use UML. Additionally, domain and technical experts must regularly interact with one another throughout the process of building such a model.



![image-20220506223705152](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506223705152.png)





In simple terms, Domain Driven Design is a set of tools, principles, and patterns that helps us address the challenges we face when a complex domain is modeled. DDD help us not only from a tactical and technical perspective , but also a strategic one.



### What does Domain Driven Design focus on?



DDD seeks to satisfy three main topics, which are very simple:



1. When designing and developing software that will support some business process in a company, the most important way to approach it is by modeling the domain and every subdomain, rather than defining the technological stack. In many cases, technical experts tend to think of it as a technology instead of a business at the beginning of the project. They may think of it as a database name, a server name, or even a programming language. These aspects are not relevant to DDD when the project starts. The most important aspect from the DDD perspective is domain modeling.
2. ![image-20220506223922683](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506223922683.png)
3. s was mentioned before, DDD is appropriate for dealing with complex domains. Thus, when a team analyzes a domain and creates the domain model, the team should break the domain into smaller pieces called subdomains. The team should, then, comprehend these subdomains to build the best model, which will represent the business processes.
4. DDD is oriented to work with agile methodologies, such as **Scrum**. This is why the concept of **Minimum Viable Product (MVP)** fits in perfectly with DDD. The most important step is to create a basic model that everyone understands. Another important consideration is to create a model that everyone can evolve sprint by sprint, until a complete product is formed.



### Building blocks of Domain-Driven Design 

DDD defines some concepts that can be used to meet its restrictions. These concepts will be covered in more detail in the upcoming lessons but, for now, they will only be listed:



- Entity
- Value Object
- Aggregate
- Domain Event
- Command
- Query
- Service
- Repository
- Factory



### Strategic approach

The strategic approach should be the first approach that was tackled during the implementation of DDD. This approach addresses the efforts made in trying to model a domain, and gives us patterns and ways to achieve the goals that DDD sets out. Think of it as the general way to deal with the problem. It applies to every piece of the model and spans across many subdomains. In this approach, there are design patterns, such as bounded context and ubiquitous language.

![image-20220506224538861](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506224538861.png)





The image above shows how the strategic approach involves different components.





### Tactical approach 



The tactical approach involves the translation of the model into specific pieces of software. It only applies to a specific component or part of the model. This approach tries to solve the challenges that arise when the development process of a specific component is underway. It gives us design patterns to develop components, including the **aggregate pattern** or the **repository pattern**. These will be covered in more detail in the upcoming chapters.

![image-20220506224629372](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506224629372.png)





The image above shows how the tactical approach involves only one component.



Domain Driven Design aims to develop any complex software based on understanding a business with its processes and representing such processes in something called a model. This will bring many benefits to a company, but will come with its fair share of drawbacks.



