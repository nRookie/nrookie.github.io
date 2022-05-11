### **Index**

When performing queries across multiple indexes, it is sometimes desirable to add query clauses that are associated with documents of only certain indexes. The `_index` field allows matching on the index a document was indexed into. Its value is accessible in certain queries and aggregations, and when sorting or scripting:



| [`_index`](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index-field.html) | The index to which the document belongs. |
| ------------------------------------------------------------ | ---------------------------------------- |
| [`_id`](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-id-field.html) | The document’s ID.                       |



## `_source` field

The `_source` field contains the original JSON document body that was passed at index time. The `_source` field itself is not indexed (and thus is not searchable), but it is stored so that it can be returned when executing *fetch* requests, like [get](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html) or [search](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html).



**[`_source`](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-source-field.html)**

The original JSON representing the body of the document.

**[`_size`](https://www.elastic.co/guide/en/elasticsearch/plugins/8.1/mapper-size.html)**

The size of the `_source` field in bytes, provided by the [`mapper-size` plugin](https://www.elastic.co/guide/en/elasticsearch/plugins/8.1/mapper-size.html).





https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-source-field.html









**Ref**

https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-fields.html



https://www.elastic.co/guide/en/elasticsearch/reference/7.1/search-request-body.html