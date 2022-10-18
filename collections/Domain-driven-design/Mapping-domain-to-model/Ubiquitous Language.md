# Ubiquitous Language



Learn what ubiquitous language is, and how it helps domain and technical experts speak in the same language.



Imagine a situation in which it is necessary to translate a word into two different languages. For instance, in Italian, pasta is a type of food, but in Spanish, it can be a type of medicine. Another example is the word “gift”. In Scandinavia, it means to get married, whereas it means “poison” in German. These comparisons are mentioned because similar differences in meaning can come up when we are designing a model. These differences in meaning tend to cause problems and misunderstandings between technical and domain experts.



## What is ubiquitous language?

When a domain is modeled, for instance, during an event storming session, business experts may use business jargon and technical experts may use technical jargon. This will require both sides to translate the terms between these two worlds, and may cause mistranslations and misinterpretations of certain terms.



In the DDD world, domain and technical experts must speak the same language. They must use the same terminology, because this will guarantee that the software evolves quickly over time.



To solve this problem of a difference in technical and business language, DDD offers a strategic pattern called ubiquitous language. Ubiquitous language aims to avoid translations between the business and technical world to create a better understanding of the domain. Software that is built on top of ubiquitous language is easy to understand, because it reflects business terminology in its code. Such software needs access to the exact names and terms that are defined in a model.



![image-20220507220835027](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507220835027.png)



### Challenges in communication among experts



DDD suggests that domain and technical experts must be in constant communication during the construction process of any software. Nonetheless, this leads to further difficulties.





1. Technical experts need to learn multiple business terms that they are not familiar with. A company may use several different terms, across all of its teams. To memorize and understand those terms is a laborious task. The image below shows an example of banking services with different business processes or subdomains. Each subdomain uses its own terminology. Now, the technical experts must also be able to comprehend these terms.
2. ![image-20220507220955677](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507220955677.png)
3. A term may appear in different subdomains. This may lead to a misunderstanding of the meaning of certain words, because the meaning for one word in one subdomain is not the same as the meaning of that word in another subdomain. As the image below shows, the term “credit” appears in the subdomains: “Credit cards” and “Saving accounts”. However, the meaning of this term is different in each of them. Credit in the “Credit cards” subdomain means to reduce the limit of that credit card. Whereas, in the “Saving accounts” subdomain, credit means to add money to an account.
4. ![image-20220507221218911](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507221218911.png)
5. Technical experts tend to translate business concepts into technical terms. This is not a good practice, because it can cause technical experts to misunderstand the domain through the translation.
6. ![image-20220507221306934](/Users/kestrel/developer/nrookie.github.io/collections/Domain-driven-design/Mapping-domain-to-model/image-20220507221306934.png)

### How to tackle communication challenges

Ubiquitous language solves this problem. In each business context, ubiquitous language is used by all of the participants in the model-definition process.



The model has common terms and acronyms, which are all used within a defined context. For example, if someone works in a specific context, it is implied that there is a context of particular and understandable terms. However, consider the case where a company uses multiple ubiquitous languages, due to the presence of multiple subdomains or contexts.



Ubiquitous language should evolve over time. It should be well-documented in a repository, such as in a Confluence, for a team to manage their common terminology. This language consists of terms that are commonly used by business and technical experts. In other words, ubiquitous language is made up of a mixture of terminologies from both the business and the technical worlds.



Ubiquitous language must be used across source code, testing code, daily conversations between domain and technical experts, and documentation. It guarantees a correct understanding of the domain and avoids the need for translations, because it ensures that everyone speaks the same language.



## Summary

Ubiquitous language allows for domain and technical experts to understand a business in the same way. Regardless of the domain and solution complexity, it guarantees that everyone will speak in the same terminologies, all across the modeling process and coding stage.





