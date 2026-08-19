---
title: Markdown example
weight: 10
description: Example of a generated markdown
hidden: true
---


# Swagger Petstore
This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
  
> [Find out more about Swagger](http://swagger.io)

## Informations

### Version

1.0.0

### License

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

 apiteam@swagger.io 

### Terms Of Service

http://swagger.io/terms/

## Tags

  ### <span id="tag-pet"></span>[pet](http://swagger.io "Find out more")

Everything about your Pets

  ### <span id="tag-store"></span>store

Access to Petstore orders

  ### <span id="tag-user"></span>[user](http://swagger.io "Find out more about our store")

Operations about user

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json
  * multipart/form-data
  * application/x-www-form-urlencoded
  * application/xml

### Produces
  * application/json
  * application/xml

## Access control

### Security Schemes

#### api_key (header: api_key)



> **Type**: apikey

#### petstore_auth



> **Type**: oauth2
>
> **Flow**: implicit
>
> **Authorization URL**: http://petstore.swagger.io/oauth/dialog
      

##### Scopes

Name | Description
-----|-------------
write:pets | modify pets in your account
read:pets | read your pets

## All endpoints

###  pet

  
> [Find out more](http://swagger.io)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/pet | [add pet](#add-pet) | Add a new pet to the store |
| DELETE | /v2/pet/{petId} | [delete pet](#delete-pet) | Deletes a pet |
| GET | /v2/pet/findByStatus | [find pets by status](#find-pets-by-status) | Finds Pets by status |
| GET | /v2/pet/findByTags | [find pets by tags](#find-pets-by-tags) | Finds Pets by tags |
| GET | /v2/pet/{petId} | [get pet by Id](#get-pet-by-id) | Find pet by ID |
| PUT | /v2/pet | [update pet](#update-pet) | Update an existing pet |
| POST | /v2/pet/{petId} | [update pet with form](#update-pet-with-form) | Updates a pet in the store with form data |
| POST | /v2/pet/{petId}/uploadImage | [upload file](#upload-file) | uploads an image |
  


###  store

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /v2/store/order/{orderId} | [delete order](#delete-order) | Delete purchase order by ID |
| GET | /v2/store/inventory | [get inventory](#get-inventory) | Returns pet inventories by status |
| GET | /v2/store/order/{orderId} | [get order by Id](#get-order-by-id) | Find purchase order by ID |
| POST | /v2/store/order | [place order](#place-order) | Place an order for a pet |
  


###  user

  
> [Find out more about our store](http://swagger.io)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v2/user | [create user](#create-user) | Create user |
| POST | /v2/user/createWithArray | [create users with array input](#create-users-with-array-input) | Creates list of users with given input array |
| POST | /v2/user/createWithList | [create users with list input](#create-users-with-list-input) | Creates list of users with given input array |
| DELETE | /v2/user/{username} | [delete user](#delete-user) | Delete user |
| GET | /v2/user/{username} | [get user by name](#get-user-by-name) | Get user by user name |
| GET | /v2/user/login | [login user](#login-user) | Logs user into the system |
| GET | /v2/user/logout | [logout user](#logout-user) | Logs out current logged in user session |
| PUT | /v2/user/{username} | [update user](#update-user) | Updated user |
  


## Paths

### <span id="add-pet"></span> Add a new pet to the store (*addPet*)

```
POST /v2/pet
```

#### Consumes
  * application/json
  * application/xml

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [Pet](#pet) | `models.Pet` | | ✓ | | Pet object that needs to be added to the store |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [405](#add-pet-405) | Method Not Allowed | Invalid input |  | [schema](#add-pet-405-schema) |

#### Responses


##### <span id="add-pet-405"></span> 405 - Invalid input
Status: Method Not Allowed

###### <span id="add-pet-405-schema"></span> Schema

### <span id="create-user"></span> Create user (*createUser*)

```
POST /v2/user
```

This can only be done by the logged in user.

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [User](#user) | `models.User` | | ✓ | | Created user object |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#create-user-default) | | successful operation |  | [schema](#create-user-default-schema) |

#### Responses


##### <span id="create-user-default"></span> Default Response
successful operation

###### <span id="create-user-default-schema"></span> Schema
empty schema

### <span id="create-users-with-array-input"></span> Creates list of users with given input array (*createUsersWithArrayInput*)

```
POST /v2/user/createWithArray
```

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [][User](#user) | `[]*models.User` | | ✓ | | List of user object |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#create-users-with-array-input-default) | | successful operation |  | [schema](#create-users-with-array-input-default-schema) |

#### Responses


##### <span id="create-users-with-array-input-default"></span> Default Response
successful operation

###### <span id="create-users-with-array-input-default-schema"></span> Schema
empty schema

### <span id="create-users-with-list-input"></span> Creates list of users with given input array (*createUsersWithListInput*)

```
POST /v2/user/createWithList
```

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [][User](#user) | `[]*models.User` | | ✓ | | List of user object |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#create-users-with-list-input-default) | | successful operation |  | [schema](#create-users-with-list-input-default-schema) |

#### Responses


##### <span id="create-users-with-list-input-default"></span> Default Response
successful operation

###### <span id="create-users-with-list-input-default-schema"></span> Schema
empty schema

### <span id="delete-order"></span> Delete purchase order by ID (*deleteOrder*)

```
DELETE /v2/store/order/{orderId}
```

For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| orderId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | ID of the order that needs to be deleted |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [400](#delete-order-400) | Bad Request | Invalid ID supplied |  | [schema](#delete-order-400-schema) |
| [404](#delete-order-404) | Not Found | Order not found |  | [schema](#delete-order-404-schema) |

#### Responses


##### <span id="delete-order-400"></span> 400 - Invalid ID supplied
Status: Bad Request

###### <span id="delete-order-400-schema"></span> Schema

##### <span id="delete-order-404"></span> 404 - Order not found
Status: Not Found

###### <span id="delete-order-404-schema"></span> Schema

### <span id="delete-pet"></span> Deletes a pet (*deletePet*)

```
DELETE /v2/pet/{petId}
```

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| petId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | Pet id to delete |
| api_key | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [400](#delete-pet-400) | Bad Request | Invalid ID supplied |  | [schema](#delete-pet-400-schema) |
| [404](#delete-pet-404) | Not Found | Pet not found |  | [schema](#delete-pet-404-schema) |

#### Responses


##### <span id="delete-pet-400"></span> 400 - Invalid ID supplied
Status: Bad Request

###### <span id="delete-pet-400-schema"></span> Schema

##### <span id="delete-pet-404"></span> 404 - Pet not found
Status: Not Found

###### <span id="delete-pet-404-schema"></span> Schema

### <span id="delete-user"></span> Delete user (*deleteUser*)

```
DELETE /v2/user/{username}
```

This can only be done by the logged in user.

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| username | `path` | string | `string` |  | ✓ |  | The name that needs to be deleted |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [400](#delete-user-400) | Bad Request | Invalid username supplied |  | [schema](#delete-user-400-schema) |
| [404](#delete-user-404) | Not Found | User not found |  | [schema](#delete-user-404-schema) |

#### Responses


##### <span id="delete-user-400"></span> 400 - Invalid username supplied
Status: Bad Request

###### <span id="delete-user-400-schema"></span> Schema

##### <span id="delete-user-404"></span> 404 - User not found
Status: Not Found

###### <span id="delete-user-404-schema"></span> Schema

### <span id="find-pets-by-status"></span> Finds Pets by status (*findPetsByStatus*)

```
GET /v2/pet/findByStatus
```

Multiple status values can be provided with comma separated strings

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| status | `query` | []string | `[]string` | `multi` | ✓ |  | Status values that need to be considered for filter |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#find-pets-by-status-200) | OK | successful operation |  | [schema](#find-pets-by-status-200-schema) |
| [400](#find-pets-by-status-400) | Bad Request | Invalid status value |  | [schema](#find-pets-by-status-400-schema) |

#### Responses


##### <span id="find-pets-by-status-200"></span> 200 - successful operation
Status: OK

###### <span id="find-pets-by-status-200-schema"></span> Schema
   
  

[][Pet](#pet)

##### <span id="find-pets-by-status-400"></span> 400 - Invalid status value
Status: Bad Request

###### <span id="find-pets-by-status-400-schema"></span> Schema

### <span id="find-pets-by-tags"></span> Finds Pets by tags (*findPetsByTags*)

```
GET /v2/pet/findByTags
```

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tags | `query` | []string | `[]string` | `multi` | ✓ |  | Tags to filter by |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#find-pets-by-tags-200) | OK | successful operation |  | [schema](#find-pets-by-tags-200-schema) |
| [400](#find-pets-by-tags-400) | Bad Request | Invalid tag value |  | [schema](#find-pets-by-tags-400-schema) |

#### Responses


##### <span id="find-pets-by-tags-200"></span> 200 - successful operation
Status: OK

###### <span id="find-pets-by-tags-200-schema"></span> Schema
   
  

[][Pet](#pet)

##### <span id="find-pets-by-tags-400"></span> 400 - Invalid tag value
Status: Bad Request

###### <span id="find-pets-by-tags-400-schema"></span> Schema

### <span id="get-inventory"></span> Returns pet inventories by status (*getInventory*)

```
GET /v2/store/inventory
```

Returns a map of status codes to quantities

#### Produces
  * application/json

#### Security Requirements
  * api_key

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-inventory-200) | OK | successful operation |  | [schema](#get-inventory-200-schema) |

#### Responses


##### <span id="get-inventory-200"></span> 200 - successful operation
Status: OK

###### <span id="get-inventory-200-schema"></span> Schema
   
  

map of int32 (formatted integer)

### <span id="get-order-by-id"></span> Find purchase order by ID (*getOrderById*)

```
GET /v2/store/order/{orderId}
```

For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| orderId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | ID of pet that needs to be fetched |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-order-by-id-200) | OK | successful operation |  | [schema](#get-order-by-id-200-schema) |
| [400](#get-order-by-id-400) | Bad Request | Invalid ID supplied |  | [schema](#get-order-by-id-400-schema) |
| [404](#get-order-by-id-404) | Not Found | Order not found |  | [schema](#get-order-by-id-404-schema) |

#### Responses


##### <span id="get-order-by-id-200"></span> 200 - successful operation
Status: OK

###### <span id="get-order-by-id-200-schema"></span> Schema
   
  

[Order](#order)

##### <span id="get-order-by-id-400"></span> 400 - Invalid ID supplied
Status: Bad Request

###### <span id="get-order-by-id-400-schema"></span> Schema

##### <span id="get-order-by-id-404"></span> 404 - Order not found
Status: Not Found

###### <span id="get-order-by-id-404-schema"></span> Schema

### <span id="get-pet-by-id"></span> Find pet by ID (*getPetById*)

```
GET /v2/pet/{petId}
```

Returns a single pet

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * api_key

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| petId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | ID of pet to return |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-pet-by-id-200) | OK | successful operation |  | [schema](#get-pet-by-id-200-schema) |
| [400](#get-pet-by-id-400) | Bad Request | Invalid ID supplied |  | [schema](#get-pet-by-id-400-schema) |
| [404](#get-pet-by-id-404) | Not Found | Pet not found |  | [schema](#get-pet-by-id-404-schema) |

#### Responses


##### <span id="get-pet-by-id-200"></span> 200 - successful operation
Status: OK

###### <span id="get-pet-by-id-200-schema"></span> Schema
   
  

[Pet](#pet)

##### <span id="get-pet-by-id-400"></span> 400 - Invalid ID supplied
Status: Bad Request

###### <span id="get-pet-by-id-400-schema"></span> Schema

##### <span id="get-pet-by-id-404"></span> 404 - Pet not found
Status: Not Found

###### <span id="get-pet-by-id-404-schema"></span> Schema

### <span id="get-user-by-name"></span> Get user by user name (*getUserByName*)

```
GET /v2/user/{username}
```

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| username | `path` | string | `string` |  | ✓ |  | The name that needs to be fetched. Use user1 for testing. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-by-name-200) | OK | successful operation |  | [schema](#get-user-by-name-200-schema) |
| [400](#get-user-by-name-400) | Bad Request | Invalid username supplied |  | [schema](#get-user-by-name-400-schema) |
| [404](#get-user-by-name-404) | Not Found | User not found |  | [schema](#get-user-by-name-404-schema) |

#### Responses


##### <span id="get-user-by-name-200"></span> 200 - successful operation
Status: OK

###### <span id="get-user-by-name-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="get-user-by-name-400"></span> 400 - Invalid username supplied
Status: Bad Request

###### <span id="get-user-by-name-400-schema"></span> Schema

##### <span id="get-user-by-name-404"></span> 404 - User not found
Status: Not Found

###### <span id="get-user-by-name-404-schema"></span> Schema

### <span id="login-user"></span> Logs user into the system (*loginUser*)

```
GET /v2/user/login
```

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| password | `query` | string | `string` |  | ✓ |  | The password for login in clear text |
| username | `query` | string | `string` |  | ✓ |  | The user name for login |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-user-200) | OK | successful operation | ✓ | [schema](#login-user-200-schema) |
| [400](#login-user-400) | Bad Request | Invalid username/password supplied |  | [schema](#login-user-400-schema) |

#### Responses


##### <span id="login-user-200"></span> 200 - successful operation
Status: OK

###### <span id="login-user-200-schema"></span> Schema
   
  



###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| X-Expires-After | date-time (formatted string) | `strfmt.DateTime` |  |  | date in UTC when token expires |
| X-Rate-Limit | int32 (formatted integer) | `int32` |  |  | calls per hour allowed by the user |

##### <span id="login-user-400"></span> 400 - Invalid username/password supplied
Status: Bad Request

###### <span id="login-user-400-schema"></span> Schema

### <span id="logout-user"></span> Logs out current logged in user session (*logoutUser*)

```
GET /v2/user/logout
```

#### Produces
  * application/xml
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#logout-user-default) | | successful operation |  | [schema](#logout-user-default-schema) |

#### Responses


##### <span id="logout-user-default"></span> Default Response
successful operation

###### <span id="logout-user-default-schema"></span> Schema
empty schema

### <span id="place-order"></span> Place an order for a pet (*placeOrder*)

```
POST /v2/store/order
```

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [Order](#order) | `models.Order` | | ✓ | | order placed for purchasing the pet |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#place-order-200) | OK | successful operation |  | [schema](#place-order-200-schema) |
| [400](#place-order-400) | Bad Request | Invalid Order |  | [schema](#place-order-400-schema) |

#### Responses


##### <span id="place-order-200"></span> 200 - successful operation
Status: OK

###### <span id="place-order-200-schema"></span> Schema
   
  

[Order](#order)

##### <span id="place-order-400"></span> 400 - Invalid Order
Status: Bad Request

###### <span id="place-order-400-schema"></span> Schema

### <span id="update-pet"></span> Update an existing pet (*updatePet*)

```
PUT /v2/pet
```

#### Consumes
  * application/json
  * application/xml

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [Pet](#pet) | `models.Pet` | | ✓ | | Pet object that needs to be added to the store |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [400](#update-pet-400) | Bad Request | Invalid ID supplied |  | [schema](#update-pet-400-schema) |
| [404](#update-pet-404) | Not Found | Pet not found |  | [schema](#update-pet-404-schema) |
| [405](#update-pet-405) | Method Not Allowed | Validation exception |  | [schema](#update-pet-405-schema) |

#### Responses


##### <span id="update-pet-400"></span> 400 - Invalid ID supplied
Status: Bad Request

###### <span id="update-pet-400-schema"></span> Schema

##### <span id="update-pet-404"></span> 404 - Pet not found
Status: Not Found

###### <span id="update-pet-404-schema"></span> Schema

##### <span id="update-pet-405"></span> 405 - Validation exception
Status: Method Not Allowed

###### <span id="update-pet-405-schema"></span> Schema

### <span id="update-pet-with-form"></span> Updates a pet in the store with form data (*updatePetWithForm*)

```
POST /v2/pet/{petId}
```

#### Consumes
  * application/x-www-form-urlencoded

#### Produces
  * application/xml
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| petId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | ID of pet that needs to be updated |
| name | `formData` | string | `string` |  |  |  | Updated name of the pet |
| status | `formData` | string | `string` |  |  |  | Updated status of the pet |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [405](#update-pet-with-form-405) | Method Not Allowed | Invalid input |  | [schema](#update-pet-with-form-405-schema) |

#### Responses


##### <span id="update-pet-with-form-405"></span> 405 - Invalid input
Status: Method Not Allowed

###### <span id="update-pet-with-form-405-schema"></span> Schema

### <span id="update-user"></span> Updated user (*updateUser*)

```
PUT /v2/user/{username}
```

This can only be done by the logged in user.

#### Produces
  * application/xml
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| username | `path` | string | `string` |  | ✓ |  | name that need to be updated |
| body | `body` | [User](#user) | `models.User` | | ✓ | | Updated user object |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [400](#update-user-400) | Bad Request | Invalid user supplied |  | [schema](#update-user-400-schema) |
| [404](#update-user-404) | Not Found | User not found |  | [schema](#update-user-404-schema) |

#### Responses


##### <span id="update-user-400"></span> 400 - Invalid user supplied
Status: Bad Request

###### <span id="update-user-400-schema"></span> Schema

##### <span id="update-user-404"></span> 404 - User not found
Status: Not Found

###### <span id="update-user-404-schema"></span> Schema

### <span id="upload-file"></span> uploads an image (*uploadFile*)

```
POST /v2/pet/{petId}/uploadImage
```

#### Consumes
  * multipart/form-data

#### Produces
  * application/json

#### Security Requirements
  * petstore_auth: read:pets, write:pets

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| petId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | ID of pet to update |
| additionalMetadata | `formData` | string | `string` |  |  |  | Additional data to pass to server |
| file | `formData` | file | `io.ReadCloser` |  |  |  | file to upload |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#upload-file-200) | OK | successful operation |  | [schema](#upload-file-200-schema) |

#### Responses


##### <span id="upload-file-200"></span> 200 - successful operation
Status: OK

###### <span id="upload-file-200-schema"></span> Schema
   
  

[APIResponse](#api-response)

## Models

### <span id="api-response"></span> ApiResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int32 (formatted integer)| `int32` |  | |  |  |
| message | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="category"></span> Category


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="order"></span> Order


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| complete | boolean| `bool` |  | |  |  |
| id | int64 (formatted integer)| `int64` |  | |  |  |
| petId | int64 (formatted integer)| `int64` |  | |  |  |
| quantity | int32 (formatted integer)| `int32` |  | |  |  |
| shipDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| status | string| `string` |  | | Order Status |  |



### <span id="pet"></span> Pet


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| category | [Category](#category)| `Category` |  | |  |  |
| id | int64 (formatted integer)| `int64` |  | |  |  |
| name | string| `string` | ✓ | |  | `doggie` |
| photoUrls | []string| `[]string` | ✓ | |  |  |
| status | string| `string` |  | | pet status in the store |  |
| tags | [][Tag](#tag)| `[]*Tag` |  | |  |  |



### <span id="tag"></span> Tag


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` |  | |  |  |
| firstName | string| `string` |  | |  |  |
| id | int64 (formatted integer)| `int64` |  | |  |  |
| lastName | string| `string` |  | |  |  |
| password | string| `string` |  | |  |  |
| phone | string| `string` |  | |  |  |
| userStatus | int32 (formatted integer)| `int32` |  | | User Status |  |
| username | string| `string` |  | |  |  |


