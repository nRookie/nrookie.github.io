


# annotation

## entity
The @Entity annotation specifies that the class is an entity and is mapped to a database table. The @Table annotation specifies the name of the database table to be used for mapping

## id
This is the City entity. Each entity must have at least two annotations defined: @Entity and @Id. The @Entity annotation specifies that the class is an entity and is mapped to a database table. The @Table annotation specifies the name of the database table to be used for mapping. The @Id annotation specifies the primary key of an entity and the @GeneratedValue provides for the specification of generation strategies for the values of primary keys.

## Spring Boot basic annotations
In the example application, we have these Spring Boot annotations:

@Bean - indicates that a method produces a bean to be managed by Spring.
@Service - indicates that an annotated class is a service class.
@Repository - indicates that an annotated class is a repository, which is an abstraction of data access and storage.
@Configuration - indicates that a class is a configuration class that may contain bean definitions.
@Controller - marks the class as web controller, capable of handling the requests.
@RequestMapping - maps HTTP request with a path to a controller method.
@Autowired - marks a constructor, field, or setter method to be autowired by Spring dependency injection.
@SpringBootApplication - enables Spring Boot autoconfiguration and component scanning.


@Component is a generic stereotype for a Spring managed component. It turns the class into a Spring bean at the auto-scan time. Classes decorated with this annotation are considered as candidates for auto-detection when using annotation-based configuration and classpath scanning. @Repository, @Service, and @Controller are specializations of @Component for more specific use cases.


## Spring Data JPA

Spring Data JPA, part of the larger Spring Data family, makes it easy to easily implement JPA based repositories. This module deals with enhanced support for JPA based data access layers. It makes it easier to build Spring-powered applications that use data access technologies.

Implementing a data access layer of an application has been cumbersome for quite a while. Too much boilerplate code has to be written to execute simple queries as well as perform pagination, and auditing. Spring Data JPA aims to significantly improve the implementation of data access layers by reducing the effort to the amount that’s actually needed. As a developer you write your repository interfaces, including custom finder methods, and Spring will provide the implementation automatically.


- Sophisticated support to build repositories based on Spring and JPA

- Support for Querydsl predicates and thus type-safe JPA queries

- Transparent auditing of domain class

- Pagination support, dynamic query execution, ability to integrate custom data access code

- Validation of @Query annotated queries at bootstrap time

- Support for XML based entity mapping

- JavaConfig based repository configuration by introducing @EnableJpaRepositories.

https://spring.io/projects/spring-data-jpa


## Component

public @interface Component

Indicates that an annotated class is a "component". Such classes are considered as candidates for auto-detection when using annotation-based configuration and classpath scanning.
Other class-level annotations may be considered as identifying a component as well, typically a special kind of component: e.g. the @Repository annotation or AspectJ's @Aspect annotation.



@PostConstruct - Spring calls methods annotated with @PostConstruct only once, just after the initialization of bean properties. Keep in mind that these methods will run even if there is nothing to initialize.

The method annotated with @PostConstruct can have any access level but it can't be static.


## bean 
A bean is an object that is instantiated, assembled, and otherwise managed by a Spring IoC container. Otherwise, a bean is simply one of many objects in your application. Beans, and the dependencies among them, are reflected in the configuration metadata used by a container.


## RepositoryRestResource

@Target(value=TYPE)
 @Retention(value=RUNTIME)
 @Inherited
public @interface RepositoryRestResource
Annotate a Repository with this to customize export mapping and rels.


@Indexed
public interface Repository<T,ID>
Central repository marker interface. Captures the domain type to manage as well as the domain type's id type. General purpose is to hold type information as well as being able to discover interfaces that extend this one during classpath scanning for easy Spring bean creation.
Domain repositories extending this interface can selectively expose CRUD methods by simply declaring methods of the same signature as those declared in CrudRepository.





https://spring.io/projects/spring-boot

3. @Embeddable
JPA provides the @Embeddable annotation to declare that a class will be embedded by other entities.

Let's define a class to abstract out the contact person details:

4. @Embedded
The JPA annotation @Embedded is used to embed a type into another entity.

Let's next modify our Company class. We'll add the JPA annotations and we'll also change to use ContactPerson instead of separate fields: