https://stackoverflow.com/questions/330793/how-to-initialize-a-struct-in-accordance-with-c-programming-language-standards


``` shell

#include<stdio.h>
#include"http_parser.h"


typedef struct {
        uint16_t http_major; /* HTTP major version */
        uint16_t http_minor; /* HTTP minor version */
        const char *body; /* body to send or NULL */
        size_t body_length; /* bytes in body to send or 0 */
        const char *body_encoding; /* body encoding type or NULL */
} send_http_response_args_t;



int main() {

        send_http_response_args_t args = {
                .http_major = 1,
                .http_minor = 1,
                .body_length = 0,
        };

        printf("%d\n", args.http_major);
        printf("%d\n", args.http_minor);
        printf("%d\n",10);
        return 0;
}         
```

