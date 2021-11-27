

## What is a REST API ?

A REST API is modeled as collections of individually-addressable resources (the nouns of the API). Resources are referenced with their resource names and manipulated via a small set of methods (also known as verbs or operations).

Standard methods for REST Google APIs (also known as REST methods) are List, Get, Create, Update, and Delete. Custom methods (also known as custom verbs or custom operations) are also available to API designers for functionality that doesn't easily map to one of the standard methods, such as database transactions.


## Design flow


The Design Guide suggests taking the following steps when designing resource- oriented APIs (more details are covered in specific sections below):

- Determine what types of resources an API provides.

- Determine the relationships between resources.

- Decide the resource name schemes based on types and relationships.

- Decide the resource schemas.

- Attach minimum set of methods to resources.

## Resources

A resource-oriented API is generally modeled as a resource hierarchy, where each node is either a simple resource or a collection resource. For convenience, they are often called a resource and a collection, respectively.


* A collection contains a list of resources of the same type. For example, a user has a collection of contacts.

* A resource has some state and zero or more sub-resources. Each sub-resource can be either a simple resource or a collection resource.

## Methods

The key characteristic of a resource-oriented API is that it emphasizes resources (data model) over the methods performed on the resources (functionality). A typical resource-oriented API exposes a large number of resources with a small number of methods. The methods can be either the standard methods or custom methods. For this guide, the standard methods are: List, Get, Create, Update, and Delete.

Where API functionality naturally maps to one of the standard methods, that method should be used in the API design. For functionality that does not naturally map to one of the standard methods, custom methods may be used. Custom methods offer the same design freedom as traditional RPC APIs, which can be used to implement common programming patterns, such as database transactions or data analysis.



