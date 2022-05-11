## What exactly is a subdomain? 



A subdomain is nothing but a sub-process of the domain. In other words, a domain is composed of several different activities. These activities are called subdomains. The image below shows how a domain consists of different sub-parts, where every sub-part is a subdomain.

![image-20220506222254657](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506222254657.png)







However, a subdomain becomes a domain when people zero in on a particular subdomain during the modeling process of a domain. This means that the relationship between a domain and subdomain depends on what part of the business is under analysis. For example, if subdomain A was analyzed, it would become the new domain and come to possess its own activities, as is shown in the image below.





![image-20220506222357597](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506222357597.png)



### Types of subdomains



Inside a project, we may find three types of subdomains:



1. **Core domain**: This is where the main business flows and models are understood and defined. In other words, the core domain is where the core activities of a domain reside. For instance, in a veterinary business, the main activities can be appointments, medicine management, surgeries, and so on.
2. **Supporting subdomain**: This contains all of the activities that need to be built, in order to support the core subdomain. With regards to the previous example, a veterinary business should keep a stock of medicine. The veterinary business can manage stocks of different types of medicine, as the main activity of the business is curing the illnesses of animals.
3. **Generic subdomain**: These activities are common across many domains, but they are not part of the core subdomain. They can include activities related to employee management or payrolls, for example.



### Having subdomains is important

As we saw in the previous lesson, it is not easy to understand a domain. Now, imagine that we want to go deeper and model every part of a certain business. Every part is complex in its own way, and the idea is to understand each part. Once all the parts are established, the subdomains will become evident. We should analyze and understand subdomains separately. Let us continue with one of the domain examples that we looked at in the previous lesson and dive deeper into each of its subdomains. Remember that a bank executes several activities, such as loans and account management, ATM operation, payrolls, customer support, and so on. Each activity has its own set of processes, which we should break down to make it easier to understand.



![image-20220506223023980](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506223023980.png)



The blue shapes above represent subdomains. Every subdomain has different subprocesses. For instance, when it comes to ATM management, a bank manages money inside the machine, physical and virtual security, cards, and so on. Similarly, with loan management, a bank should keep in mind the fees and payments.



Similar sub-processes are considered in every subdomain. It is almost impossible for one person to know everything in a domain, because it is too complex. People specialize in only one domain, because a person can learn about different subdomains, but they can only be an expert in one.

## Domain experts 



As we mentioned earlier, it is not easy to understand every activity that a domain carries. This is where the domain experts come in. Domain experts are people who know the most about a subdomain. The image below shows that in the *Account* subdomain, there is someone who knows everything related to money, cards, and savings. Similarly, there are loan experts, ATM experts, and so on.



![image-20220506223301886](/Users/user/playground/share/nrookie.github.io/collections/Domain-driven-design/image-20220506223301886.png)

## Summary 



A subdomain is a subpart of the big domain model, which operates through its own business logic and processes. A domain expert can understand this subpart, in addition to all of the interactions between the subdomains that account for the domain.





