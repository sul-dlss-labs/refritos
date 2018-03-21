# Refritos

**Note:** based on the [AWS Kinesis Client Library - Ruby example](https://github.com/awslabs/amazon-kinesis-client-ruby/tree/master/samples)
## Dependencies

1. Localstack: This requires that the localstack kinesis and dynamodb services are running.

```shell
SERVICES=dynamodb,kinesis,s3 localstack start
```

1. JAVA_HOME must be set in order to execute the MultiLangDaemon properly

1. Run the rake job to start the KCL service:

```shell
rake run
```