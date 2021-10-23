An exception occurred during task execution. To see the full traceback, use -vvv. The error was: botocore.exceptions.ClientError: An error occurred (AuthFailure) when calling the DescribeVpcs operation: AWS was not able to validate the provided access credentials
fatal: [localhost]: FAILED! => {"boto3_version": "1.18.63", "botocore_version": "1.21.63", "changed": false, "error": {"code": "AuthFailure", "message": "AWS was not able to validate the provided access credentials"}, "msg": "Failed to describe VPCs: An error occurred (AuthFailure) when calling the DescribeVpcs operation: AWS was not able to validate the provided access credentials", "response_metadata": {"http_headers": {"cache-control": "no-cache, no-store", "content-type": "text/xml;charset=UTF-8", "date": "Sun, 17 Oct 2021 01:42:23 GMT", "server": "AmazonEC2", "strict-transport-security": "max-age=31536000; includeSubDomains", "transfer-encoding": "chunked", "vary": "accept-encoding", "x-amzn-requestid": "94e445f7-5fc0-4b77-ab18-8ca76e23b608"}, "http_status_code": 401, "request_id": "94e445f7-5fc0-4b77-ab18-8ca76e23b608", "retry_attempts": 0}}


## 
should use following command.
``` shell
docker run -it  --volume "$(pwd)":/tmp/ansible --env "AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID" --env "AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY"  kestrel7/ansible-docker /bin/bash 
```