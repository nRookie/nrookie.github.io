Events are baked into a domain, and communication between them occurs naturally. They are sent and consumed by different subdomains, and these interactions are translated into software in the future. We must understand the events to understand how a domain operates. There are some techniques that people can use to stay updated with the requirements, interactions, and restrictions that occur in a domain. Moreover, the scope of every requirement, interaction, and restriction is capped in a clear way for business and technical experts.



## Why do we require a scoping session?



It is not an easy task to extract information from domain experts to write requirements for a system. Domain experts might explain business processes in one way, and technical experts might interpret them differently. This may lead to the construction of technical models —maybe in UML— which domain experts might not understand. Thus, there should be a session where both sides share their knowledge in the same terms. This will cause the requirements to appear naturally and for both sides, business and technical, to produce the domain model easily.



## What is event storming?

Event storming is a collaborative workshop, created in 2012 by Alberto Brandolini, which allows technical and domain experts to create and share their knowledge of complex business domains and processes. It was widely adopted by the technology industry, and that is why it is one of the best ways through which technical teams can acquire knowledge about business processes from domain experts. Moreover, it is entertaining and easy to use, because sticky notes and pens are used to present information regarding business processes.







### Requirements to carry out an event storming session

Some requirements should be met for us to be successful at event storming:

- Use different-colored sticky notes. Event storming maintains a standard, but it is possible to change colors. The only restriction is in relation to consistency.
- Use marker pens to write on the sticky notes.
- We need access to a surface on which we can pin up our sticky notes. This surface can be a wall, a big whiteboard, or a paper roll.
- If meeting in person, make sure to get a room with enough light and space. g. If the meeting is not in person, we can use a collaborative application, such as [Miro](https://miro.com/).
- Give a brief summary of what the session is about.
- Include the following roles:
  - **Facilitator**: A person who is responsible for the organization of the session and put on duty to ensure that the rules of the event storming are followed throughout the session.
  - **Domain experts**: People from different subdomains. Their role as domain experts depends on the processes and features that are set to be modeled.
  - **Technical experts**: People who are responsible for the design and development of the software. They should be skilled in the extraction of requirements from domain experts.



### Required elements in an event storming session 

Let us begin with an effort to analyse and understand each part of an example. Suppose there is a bank with a subdomain, transfers. A customer wants to transfer money from bank A to bank B. The image below shows what a transfer looks like from an event storming perspective.



![image-20220507214135384](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507214135384.png)



Event storming involves some rules, elements, and considerations that help participants take advantage of this session. Firstly, it is necessary to organize the sticky notes from left to right. It is advisable to draw an arrow in that direction at the bottom of the surface.



The individual colors of the sticky notes each represent something:

- **Orange** sticky notes represent events. **Events** are actions that happen in a domain. They are written in the past tense use of the domain experts’ language.
- **Blue** sticky notes represent commands. **Commands** represent the intent of an actor. They outline what an actor wants to do.
- **Purple** sticky notes represents policies, Reactions, combined with events, represent policies. These reactions are necessary prerequisites to the execution of a command. If this prerequisite is not met, the reactions may even replace a command. They can be understood with the phrase: "whenever this event happens , then carry out this command or action."
- **Yellow** sticky notes represent actors. **Actors** represent the people or systems that execute a command.
- **Pink** sticky notes represent third-party or internal systems. **Third-party** or **internal systems** can emit or receive events.
- **Red** sticky notes represent questions or issues. **Questions** or **issues** reference something that is scheduled to be considered later on in the session. They should be explored in more detail.
- **Green** sticky notes represent the read model. The **read model** highlights the specific data that allows one to make a decision or carry out an action. They come from a SQL query, a specific part of the UI, or any data source. This sticky note can be placed below the command/event pair.
- **Pale yellow** sticky notes represent aggregates. **Aggregates** represent a specific business logic, with a particular responsibility. Aggregates let us discover bounded contexts.



### Steps to carry out an event storming session

A few steps must be followed for a successful event storming session to take place:

1. All the events that occur in the process should be mapped.

2. All the commands that cause those events should also be mapped.

3. Additionally, all the policies or reactions that can control something after an event should be mapped.

   > We are not required to strictly follow the steps given below:
   >
   > 

4. Map external services, questions, and actors.
5. Map the aggregates and the read model.
6. Once the steps shown above are executed, aggregates should be grouped together to find boundaries among functionalities. The idea is to draw a line around the same aggregates. A timeline is not required here and can be deleted.



![image-20220507215347447](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215347447.png)



![image-20220507215426584](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215426584.png)



![image-20220507215445082](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215445082.png)



![image-20220507215520607](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215520607.png)





![image-20220507215543216](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215543216.png)



![image-20220507215616374](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507215616374.png)



Event storming is a powerful technique, which helps technical and domain experts understand how the business operates. It leads to the definition of a model that is coherent for both sides, because everyone only speaks in terms of business processes.

