Hexagonal architecture is based on a few principles that should be followed to build well-structured components. If we overlook any of them, we will experience problems in the future. Additionally, it can be quite tortuous to change bad implementations, especially in a system that runs on production.







## SOLID principles



To implement components based on hexagonal architecture, it is necessary to understand the **SOLID** principles, which are as follows:

1. **The single-responsibility principle**: This principle states that a class should only change due to one reason. The class should only execute one functionality.
2. **The open-close principle**: This principle states that a class should be opened to extension, but closed to modification.
3. **Liskov’s substitution principle**: This principle states that the objects of a superclass should be replaceable with the objects of its subclasses, and that this replacement will not cause a break in the system.
4. **The interface-segregation principle**: This principle states that the use of one interface with many defined methods, which may not be required in some implementers, is not the best approach.
5. **The dependency-inversion principle**: This principle states that high-level functions should be reusable and unaffected by changes in low-level functions. To follow this principle, it is necessary to introduce abstractions to decouple both the levels from one another. This is the most important principle that is used to implement hexagonal architecture.





## Principles of hexagonal architecture

The pattern of hexagonal architecture is based on three principles which are, in turn, supported by the SOLID principles.

### Clear-separation principle

There must be a clear separation of responsibilities. In every component, there is a way to identify these responsibilities through analysis. For example, we can identify these responsibilities if we analyze a specific action that is executed or a way that is used to save information. To comply with this principle, it is important to separate these functionalities from one another. The following image shows the basic form of the separation of responsibilities.



![image-20220509213354711](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220509213354711.png)





### Layer boundaries well-defined

Layer boundaries must be clearly defined. To comply with this principle, it is good practice to define interfaces that expose functionalities and the data that is required to execute them. Dependency between the layers should be based on the interfaces.

![image-20220509213630884](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220509213630884.png)





### Communication from the outside to the inside

Communication must always come from the outermost layer to the innermost one. This means that every execution starts in the API adapter. Business logic is executed afterward and, ultimately, the storage logic is also executed. Internal layers cannot know anything about the external layers.



![image-20220509213719922](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Hexagonal Architecture/image-20220509213719922.png)



## Summary



Hexagonal architecture works with three principles. These principles call for the separation of responsibilities, well-defined boundaries, and communication to take place from the outside to the inside.





