# MEGASDK-REST 
MegaSDK downloading functionality as a rest server.

# Documentation
Documentation is divided into two categories.
## server usage
Read output of --help command

## client usage
There are basically five endpoints on server.
```
POST /login
POST /adddl
POST /canceldl
GET /dlinfo/{gid}
GET /ping
```

**Login**: 
Do a POST on /login with payload as
``` 
{"email":email,"password":password}
```
Process the response.

**Add download**:
Do a POST on /adddl with payload as 
```
{"link":link,"dir":directory}
```
The directory in the context is the directory on the machine where the server is running, the response will return a random string called gid which you are supposed to store for later use.

**Cancel download**:
Do a POST on /canceldl with payload as 
```
{"gid":gidToCancel}
```
process the response.

**Get current info of the download by gid**:
Do a GET on 
```
/dlinfo/{gid}
``` 
the gid here is variable, if the server have that dl in its storage then it will return details of the download, if not then empty details with friendly message about what went wrong will be returned.

**Ping**: 
Just for testing if the server is up or not.