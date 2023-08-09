# Dynamodb Example

## Run
```
docker-compose up -d
```
- Note : POST API will create table

## References
- CLI commands : [SDK for Go v2](https://docs.aws.amazon.com/code-library/latest/ug/go_2_dynamodb_code_examples.html)
- Connecting to dynamodb : [Dynamodb local go](https://davidagood.com/dynamodb-local-go/)
- How to mount volume : Using DATA_DIR click [here](https://stackoverflow.com/questions/70189785/how-to-keep-data-of-localstack-s3-after-docker-compose-down-and-later-up)

## CLI Commands

#### Create Table : 
```
aws dynamodb --endpoint-url=http://localhost:4599 create-table --table-name Music --attribute-definitions AttributeName=Artist,AttributeType=S AttributeName=SongTitle,AttributeType=S --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=5
```

#### Read from Table :
```
aws dynamodb scan --endpoint=http://localhost:4599 --table-name Music
```

#### Put into Table:
```
aws --endpoint-url=http://localhost:4599 dynamodb put-item --table-name Music --item '{"Artist":{"S": "Guns N Roses"},"SongTitle":{"S":"Sweet Child O Mine"}}'
```