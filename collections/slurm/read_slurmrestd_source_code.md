


## 

_find_plugins (listf=0x0, plugin_list=0x0) at data.c:266
 if ((rc = plugrack_read_dir(rack, slurm_conf.plugindir)))

(gdb) p slurm_conf.plugindir
$48 = 0x629d50 "/opt/slurm/21.08.2//lib64/slurm"



plugin_load_from_file



data.c : _load_plugins()





plugrack_create("rest_auth");


 



## init openapi

value of oas_specs specs. "openapi/v0.0.37,dbv0.0.37"

if ((rc = data_init(MIME_TYPE_JSON_PLUGIN, NULL)))
is used to load the JSON PLUGIN.


(gdb) p *t->rack
$135 = {entries = 0x6306d0, major_type = 0x630340 "openapi"}


plugrack_read_dir


char *pbuf = xstrdup(plugins);

```
		while (type) {
			xstrtrim(type);

			/* Permit both prefix and no-prefix for plugin names. */
			if (xstrncmp(type, "openapi/", 8) == 0)
				type += 8;
			type = xstrdup_printf("openapi/%s", type);
			xstrtrim(type);

			_oas_plugrack_foreach(type, NULL, PLUGIN_INVALID_HANDLE,
					      t);

			xfree(type);
			type = strtok_r(NULL, ",", &last);
		}


```

type is openapi/v0.0.37



``` c
int
plugin_get_syms( plugin_handle_t plug,
		 int n_syms,
		 const char *names[],
		 void *ptrs[] )
{
	int i, count;

	count = 0;
	for ( i = 0; i < n_syms; ++i ) {
		ptrs[ i ] = dlsym( plug, names[ i ] );
		if ( ptrs[ i ] )
			++count;
		else
			debug3("Couldn't find sym '%s' in the plugin",
			       names[ i ]);
	}

	return count;
}
```

mainly used for getting three function.

names are  slurm_openapi_p_init, slurm_openapi_p_fini,slurm_openapi_p_get_specification


corresponding to the 

``` c
static const char *syms[] = {
	"slurm_openapi_p_init",
	"slurm_openapi_p_fini",
	"slurm_openapi_p_get_specification",
};
```


call slurm_openapi_p_init in api.c (273)

``` c
extern void slurm_openapi_p_init(void)
{
	/* Check to see if we are running a supported accounting plugin */
	if (!slurm_with_slurmdbd()) {
		fatal("%s: slurm not configured with slurmdbd", __func__);
	}

	init_op_accounts();
	init_op_associations();
	init_op_config();
	init_op_cluster();
	init_op_diag();
	init_op_job();
	init_op_qos();
	init_op_tres();
	init_op_users();
	init_op_wckeys();
}
```

init_op_accounts(), bind_operation_handler("/slurmdb/v0.0.37/accounts/", op_handler_accounts)


## how does the slurmrestd read the token



find where is the entry point, or router.


``` 
#1  0x00007ffff75ef4bc in http_parser_execute (parser=parser@entry=0x7fffec000900, settings=settings@entry=0x40e340 <settings.14012>, data=<optimized out>, len=<optimized out>)
    at /usr/src/debug/http-parser-2.7.1/http_parser.c:1130
#2  0x0000000000409819 in parse_http (con=0x7fffec0009b0, x=<optimized out>) at http.c:776
#3  0x0000000000405544 in _wrap_on_data (x=0x7fffec0009b0) at conmgr.c:822
#4  0x0000000000404748 in _wrap_work (x=<optimized out>) at conmgr.c:579
#5  0x00007ffff7b6bbed in _worker (arg=0x630fe0) at workq.c:305
#6  0x00007ffff6ebbea5 in start_thread (arg=0x7ffff57a1700) at pthread_create.c:307
#7  0x00007ffff6be4b0d in clone () at ../sysdeps/unix/sysv/linux/x86_64/clone.S:111

```