
text	65,535 
longtext 	4,294,967,295



MaxLen equals to varchar65
``` golang
	return []ent.Field{field.String("id").StorageKey("fsx_id").MaxLen(65),
```

text related to text long text

but 
field.Text("f_sx_access_info").Optional().MaxLen(65535)
cannot set text from long text to text
panic: sql/schema: creating changeset for "t_fsx": changing column type for "desc" is invalid (varchar(65535) != longtext)