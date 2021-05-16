# Use a map

Use the code in the "starting code" folder.

Remove mongodb from the code.

Instead of using mongodb, store all the data in a map.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

# Now Start the server and Test

1. create user
```shell
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```
- output:
```shell
{"id":"527b3c29-88c3-4d21-b8a3-2e39e073cc81","name":"Miss Moneypenny","gender":"female","age":27}
```
2. query the user back
```shell
curl http://localhost:8080/user/527b3c29-88c3-4d21-b8a3-2e39e073cc81
```
- output:
```shell
{"id":"527b3c29-88c3-4d21-b8a3-2e39e073cc81","name":"Miss Moneypenny","gender":"female","age":27}
```
3. delete user
```shell
curl -X DELETE http://localhost:8080/user/527b3c29-88c3-4d21-b8a3-2e39e073cc81
```
- output:
```shell
Deleted user:527b3c29-88c3-4d21-b8a3-2e39e073cc81
```