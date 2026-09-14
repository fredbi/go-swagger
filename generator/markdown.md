


# Markdown generator demo
A spec to demonstrate the markdown generator.

All swagger fields supported by the doc generator are examplified below.

  
> [More information](https://goswagger.io)

## Informations

### Version

1.0.0

### License

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

fred fred@example.com https://github.com/go-swagger/go-swagger

### Terms Of Service

free to use

## Tags

  ### <span id="tag-greeters"></span>[greeters](https://welcome.example.com/greet)

define several ways to greet someone

  ### <span id="tag-ousters"></span>[ousters](https://welcome.example.com/oust "more about ousters")

define several ways to (tactfully) get rid of someone

## Content negotiation

### URI Schemes
  * http
  * https
  * unix
  * wss

### Consumes
  * application/json
  * text/plain

### Produces
  * application/json
  * text/plain

## Access control

### Security Schemes

#### api_key (header: x-api-key)

Authentication via API key header

> **Type**: apikey

#### password

Authentication with password

> **Type**: basic

#### token

Authentication with oauth2 token

> **Type**: oauth2
>
> **Flow**: accessCode
>
> **Authorization URL**: https://app/auth/authorize
>
> **Token URL**: https://app/auth/token
      

##### Scopes

Name | Description
-----|-------------
read:greeters | Scope to access all greeters
write:greeters | Scope to create or modify greeters

### Security Requirements
  * password
  * api_key

## All endpoints

###  greeters

  
> [Read more](https://welcome.example.com/greet)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /greetings | [get greetings](#get-greetings) | Retrieve all ways to say hello
 |
| POST | /greetings | [post greetings](#post-greetings) | Makes a greeter |
| PUT | /greetings | [put greetings](#put-greetings) | Updates a greeter |
  


###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /more | [delete more](#delete-more) |  |
| GET | /all | [get all](#get-all) | Retrieve all things |
| GET | /more | [get more](#get-more) | Retrieve more things |
| GET | /some | [get some](#get-some) | Retrieve some things |
| POST | /some | [post some](#post-some) | Make some things |
| PUT | /more | [put more](#put-more) |  |
| PUT | /some | [put some](#put-some) | Modify some things |
  


###  ousters

  
> [more about ousters](https://welcome.example.com/oust)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /oustings | [get oustings](#get-oustings) | Retrieve all ways to say goodbye
 |
  


## Paths

### <span id="delete-more"></span> delete more (*DeleteMore*)

```
DELETE /more
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| inlined map body | `body` | map of [Greeter](#greeter) | `map[string]models.Greeter` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-more-204) | No Content | empty |  | [schema](#delete-more-204-schema) |
| [401](#delete-more-401) | Unauthorized | complex object |  | [schema](#delete-more-401-schema) |

#### Responses


##### <span id="delete-more-204"></span> 204 - empty
Status: No Content

###### <span id="delete-more-204-schema"></span> Schema

##### <span id="delete-more-401"></span> 401 - complex object
Status: Unauthorized

###### <span id="delete-more-401-schema"></span> Schema
   
  

[ObjectWithComplexProps](#object-with-complex-props)

### <span id="get-all"></span> Retrieve all things (*GetAll*)

```
GET /all
```

#### Security Requirements
  * token: read:greeters

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-200) | OK | All in one |  | [schema](#get-all-200-schema) |

#### Responses


##### <span id="get-all-200"></span> 200 - All in one
Status: OK

###### <span id="get-all-200-schema"></span> Schema
   
  

[][AnyObject](#any-object)

### <span id="get-greetings"></span> Retrieve all ways to say hello (*GetGreetings*)

```
GET /greetings
```

Retrieve greeters according to language or other optional filter


> [Read more](https://www.greetings.io/greeters "More here")

#### Produces
  * application/json
  * text/plain

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| byID | `query` | uuid (formatted string) | `strfmt.UUID` |  |  |  |  |
| filter | `query` | string | `string` |  |  | `"none"` |  |
| language | `query` | string | `string` |  |  |  | a filter on the language |
| max | `query` | integer | `int64` |  |  | `100` |  |
| multi | `query` | []byte (base64 string) | `[]strfmt.Base64` | `pipes` |  |  | specifies multiple outputs |
| offset | `query` | int32 (formatted integer) | `int32` |  |  |  |  |
| since | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-greetings-200) | OK | A hello world collection | ✓ | [schema](#get-greetings-200-schema) |

#### Responses


##### <span id="get-greetings-200"></span> 200 - A hello world collection
Status: OK

###### <span id="get-greetings-200-schema"></span> Schema
   
  

[][Greeter](#greeter)

###### Examples
    
**application/json**
```json
{
  "id": "xyz",
  "language": "en",
  "name": "hello",
  "strength": 3
}
```
**text/plain**
```json
"'{\"id\": \"xyz\", \"name\": \"hello\", \"language\": \"en\", \"strength\": 3}'\n"
```

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| x-max-items | integer | `int64` |  | `1024` | The maximum number of returned items |
| x-rate-limit | integer | `int64` |  | `50` | The number of allowed requests in the current period |

### <span id="get-more"></span> Retrieve more things (*GetMore*)

```
GET /more
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-more-200) | OK | more items |  | [schema](#get-more-200-schema) |

#### Responses


##### <span id="get-more-200"></span> 200 - more items
Status: OK

###### <span id="get-more-200-schema"></span> Schema
   
  

[]byte (base64 string)

### <span id="get-oustings"></span> Retrieve all ways to say goodbye (*GetOustings*)

```
GET /oustings
```

Retrieve ousters according to language or other optional filter


> [Read more](https://www.greetings.io/ousters "More here")

#### Consumes
  * application/json
  * text/plain

#### Produces
  * application/json
  * text/plain

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| filters | `query` | []string | `[]string` | `pipes` |  |  |  |
| severity | `query` | int64 (formatted integer) | `int64` |  |  |  | a filter on the harshness of the ousting |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-oustings-200) | OK | Farewell messages |  | [schema](#get-oustings-200-schema) |

#### Responses


##### <span id="get-oustings-200"></span> 200 - Farewell messages
Status: OK

###### <span id="get-oustings-200-schema"></span> Schema
   
  

[][Ouster](#ouster)

### <span id="get-some"></span> Retrieve some things (*GetSome*)

```
GET /some
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-some-200) | OK | some items |  | [schema](#get-some-200-schema) |
| [203](#get-some-203) | Request Processed | nested items |  | [schema](#get-some-203-schema) |

#### Responses


##### <span id="get-some-200"></span> 200 - some items
Status: OK

###### <span id="get-some-200-schema"></span> Schema
   
  


 [AliasedArray](#aliased-array)

##### <span id="get-some-203"></span> 203 - nested items
Status: Request Processed

###### <span id="get-some-203-schema"></span> Schema
   
  

[][AliasedArray](#aliased-array)

### <span id="post-greetings"></span> Makes a greeter (*PostGreetings*)

```
POST /greetings
```

Builds a new greeter from a description


#### Consumes
  * application/json
  * text/plain

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| byID | `query` | uuid (formatted string) | `strfmt.UUID` |  |  |  |  |
| filter | `query` | string | `string` |  |  | `"none"` |  |
| max | `query` | integer | `int64` |  |  | `100` |  |
| multi | `query` | []byte (base64 string) | `[]strfmt.Base64` | `pipes` |  |  | specifies multiple outputs |
| offset | `query` | int32 (formatted integer) | `int32` |  |  |  |  |
| since | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  |  |
| greeterContent | `body` | [Greeter](#greeter) | `models.Greeter` | | ✓ | | The description of the new greeter |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-greetings-201) | Created | Created |  | [schema](#post-greetings-201-schema) |

#### Responses


##### <span id="post-greetings-201"></span> 201 - Created
Status: Created

###### <span id="post-greetings-201-schema"></span> Schema

### <span id="post-some"></span> Make some things (*PostSome*)

```
POST /some
```

A more long-winded version on how to make things.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| inlined body | `body` | []uuid (formatted string) | `[]strfmt.UUID` | | ✓ | | an inlined primitive body param |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-some-201) | Created | uuid created |  | [schema](#post-some-201-schema) |

#### Responses


##### <span id="post-some-201"></span> 201 - uuid created
Status: Created

###### <span id="post-some-201-schema"></span> Schema
   
  

any

### <span id="put-greetings"></span> Updates a greeter (*PutGreetings*)

```
PUT /greetings
```

#### URI Schemes
  * https
  * unix

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| byID | `query` | uuid (formatted string) | `strfmt.UUID` |  |  |  |  |
| filter | `query` | string | `string` |  |  | `"none"` |  |
| max | `query` | integer | `int64` |  |  | `100` |  |
| multi | `query` | []byte (base64 string) | `[]strfmt.Base64` | `pipes` |  |  | specifies multiple outputs |
| offset | `query` | int32 (formatted integer) | `int32` |  |  |  |  |
| since | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  |  |
| newGreeterContent | `body` | [Greeter](#greeter) | `models.Greeter` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-greetings-200) | OK | Modified |  | [schema](#put-greetings-200-schema) |
| [204](#put-greetings-204) | No Content | Nothing done |  | [schema](#put-greetings-204-schema) |

#### Responses


##### <span id="put-greetings-200"></span> 200 - Modified
Status: OK

###### <span id="put-greetings-200-schema"></span> Schema
   
> When modifiying things, a report about the actual
updates performed is returned.

  



[][PutGreetingsOKBodyItems0](#put-greetings-o-k-body-items0)

##### <span id="put-greetings-204"></span> 204 - Nothing done
Status: No Content

###### <span id="put-greetings-204-schema"></span> Schema
   

> [Inline reliased primitive test](https://www.greetings.io/realiased)
  



[][Realiased](#realiased)

###### Inlined models

**<span id="put-greetings-o-k-body-items0"></span> PutGreetingsOKBodyItems0**


> Modification report item

> [A modification report](https://www.greetings.io/reports)
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` | ✓ | |  |  |
| modified | boolean| `bool` | ✓ | |  |  |



### <span id="put-more"></span> put more (*PutMore*)

```
PUT /more
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| inlined map body | `body` | map of uuid (formatted string) | `map[string]strfmt.UUID` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#put-more-201) | Created | object with additional properties |  | [schema](#put-more-201-schema) |
| [204](#put-more-204) | No Content | map of defined model |  | [schema](#put-more-204-schema) |
| [401](#put-more-401) | Unauthorized | map with additionalProperties |  | [schema](#put-more-401-schema) |
| [403](#put-more-403) | Forbidden | map with constrained additionalProperties |  | [schema](#put-more-403-schema) |

#### Responses


##### <span id="put-more-201"></span> 201 - object with additional properties
Status: Created

###### <span id="put-more-201-schema"></span> Schema
   
  

[PutMoreCreatedBody](#put-more-created-body)

##### <span id="put-more-204"></span> 204 - map of defined model
Status: No Content

###### <span id="put-more-204-schema"></span> Schema
   
> define a map of a model type
  



map of [Greeter](#greeter)

##### <span id="put-more-401"></span> 401 - map with additionalProperties
Status: Unauthorized

###### <span id="put-more-401-schema"></span> Schema
   
> define a map of objects with additional properties
  



map of [PutMoreUnauthorizedBodyAnon](#put-more-unauthorized-body-anon)

##### <span id="put-more-403"></span> 403 - map with constrained additionalProperties
Status: Forbidden

###### <span id="put-more-403-schema"></span> Schema
   
> define a map of objects with constrained additional properties
  



map of [PutMoreForbiddenBodyAnon](#put-more-forbidden-body-anon)

###### Inlined models

**<span id="put-more-created-body"></span> PutMoreCreatedBody**


> define an inlined object with additionalProperties
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| a | [Greeter](#greeter)| `models.Greeter` | ✓ | |  |  |



**Additional Properties**

any

**<span id="put-more-forbidden-body-anon"></span> PutMoreForbiddenBodyAnon**


> property of constrained map container
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| b | integer| `int64` |  | |  |  |



**Additional Properties**

[][PutMoreForbiddenBodyAnonItems0](#put-more-forbidden-body-anon-items0)

**<span id="put-more-forbidden-body-anon-items0"></span> PutMoreForbiddenBodyAnonItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| c | integer| `int64` |  | |  |  |



**Additional Properties**

any

**<span id="put-more-unauthorized-body-anon"></span> PutMoreUnauthorizedBodyAnon**


> property of map container
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| b | integer| `int64` |  | |  |  |



**Additional Properties**

any

### <span id="put-some"></span> Modify some things (*PutSome*)

```
PUT /some
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| inlined complex body | `body` | [PutSomeBody](#put-some-body) | `PutSomeBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#put-some-201) | Created | additional properties |  | [schema](#put-some-201-schema) |
| [default](#put-some-default) | | default response |  | [schema](#put-some-default-schema) |

#### Responses


##### <span id="put-some-201"></span> 201 - additional properties
Status: Created

###### <span id="put-some-201-schema"></span> Schema
   
  

any

##### <span id="put-some-default"></span> Default Response
default response

###### <span id="put-some-default-schema"></span> Schema

  

[][PutSomeDefaultBodyItems0](#put-some-default-body-items0)

###### Inlined models

**<span id="put-some-body"></span> PutSomeBody**



> [an inlined complex body param](https://www.greetings.io/models/zzz)
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| z | uuid (formatted string)| `strfmt.UUID` | ✓ | |  | `c6bcc97c-a363-4bd0-9ded-8874e6616206` |
| zz | int32 (formatted integer)| `int32` |  | `15`|  | `20` |



**<span id="put-some-default-body-items0"></span> PutSomeDefaultBodyItems0**



> [generic error](https://www.greetings.io/models/errors-generic)
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | integer| `int64` | ✓ | |  |  |
| message | string| `string` |  | |  |  |



**Additional Properties**

| Type | Go type | Default | Description | Example |
|------|---------| ------- |-------------|---------|
| string | `string` ||  |  |

## Models

### <span id="realiased"></span> Realiased


  

[Primitive](#primitive)

#### Inlined models

### <span id="aliased-array"></span> aliasedArray


  

[]uuid (formatted string)

### <span id="anonymous-all-of"></span> anonymousAllOf


  


* inlined member (*AO0*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| a | string| `string` |  | `"A"`| the first letter of the alphabet | `a` |


* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| b | string| `string` | ✓ | |  |  |



### <span id="any-object"></span> anyObject


> a type composition

> [any kind of greeter or ouster](https://www.greetings.io/models/anyObject)
  




* composed type [Greeter](#greeter)
* composed type [Ouster](#ouster)

### <span id="greeter"></span> greeter



> [more about greeters](https://www.greetings.io/models/greeter)
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` | ✓ | |  | `xyz` |
| language | string| `string` |  | |  | `en` |
| name | string| `string` | ✓ | |  | `hello` |
| strength | int32 (formatted integer)| `int32` |  | |  | `3` |



### <span id="int-primitive"></span> intPrimitive


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| intPrimitive | int32 (formatted integer)| int32 | |  |  |



### <span id="object-with-complex-props"></span> objectWithComplexProps


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| prop1 | [Greeter](#greeter)| `Greeter` | ✓ | |  |  |
| prop2 | [Ouster](#ouster)| `Ouster` | ✓ | |  |  |
| prop3 | [Realiased](#realiased)| `Realiased` |  | |  |  |
| prop4 | [ObjectWithComplexPropsProp4](#object-with-complex-props-prop4)| `ObjectWithComplexPropsProp4` |  | |  |  |
| prop5 | [][ObjectWithComplexPropsProp5Items0](#object-with-complex-props-prop5-items0)| `[]*ObjectWithComplexPropsProp5Items0` |  | |  |  |
| prop6 | [ObjectWithComplexPropsProp6Tuple0](#object-with-complex-props-prop6-tuple0)| `ObjectWithComplexPropsProp6Tuple0` |  | |  |  |
| prop7 | [ObjectWithComplexPropsProp7Tuple0](#object-with-complex-props-prop7-tuple0)| `ObjectWithComplexPropsProp7Tuple0` |  | |  |  |
| prop8 | [ObjectWithComplexPropsProp8Tuple0](#object-with-complex-props-prop8-tuple0)| `ObjectWithComplexPropsProp8Tuple0` |  | |  |  |



#### Inlined models

**<span id="object-with-complex-props-prop4"></span> ObjectWithComplexPropsProp4**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| a | string| `string` |  | |  |  |



**Additional Properties**

| Type | Go type | Default | Description | Example |
|------|---------| ------- |-------------|---------|
| integer | `int64` ||  |  |

**<span id="object-with-complex-props-prop5-items0"></span> ObjectWithComplexPropsProp5Items0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| b | uuid (formatted string)| `strfmt.UUID` |  | |  |  |



**Additional Properties**

| Type | Go type | Default | Description | Example |
|------|---------| ------- |-------------|---------|
| integer | `int64` |`28`|  | `14` |

**<span id="object-with-complex-props-prop6-tuple0"></span> ObjectWithComplexPropsProp6Tuple0**


  



**Tuple members**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| P0 | integer| `int64` | ✓ | `15`| first member integer | `29` |
| P1 | string| `string` | ✓ | `"xyz"`| second member string |  |



**Additional Items**

[Greeter](#greeter)

**<span id="object-with-complex-props-prop7-tuple0"></span> ObjectWithComplexPropsProp7Tuple0**


  



**Tuple members**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| P0 | [Greeter](#greeter)| `Greeter` | ✓ | |  |  |
| P1 | [Ouster](#ouster)| `Ouster` | ✓ | |  |  |



**Additional Items**

| Type | Go type | Default | Description | Example |
|------|---------| ------- |-------------|---------|
| string | `string` ||  |  |

**<span id="object-with-complex-props-prop8-tuple0"></span> ObjectWithComplexPropsProp8Tuple0**


  



**Tuple members**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| P0 | [Greeter](#greeter)| `Greeter` | ✓ | |  |  |
| P1 | [Ouster](#ouster)| `Ouster` | ✓ | |  |  |
| P2 | integer| `int64` | ✓ | |  |  |



**Additional Items**

any

### <span id="ouster"></span> ouster


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| uid | uuid (formatted string)| `strfmt.UUID` | ✓ | |  |  |
| uname | string| `string` | ✓ | |  |  |



### <span id="primitive"></span> primitive


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| primitive | string| string | |  |  |



### <span id="primitive-all-of"></span> primitiveAllOf


  


* composed type [PrimitiveAllOfAllOf0](#primitive-all-of-all-of0)
* composed type [PrimitiveAllOfAllOf1](#primitive-all-of-all-of1)

#### Inlined models

**<span id="primitive-all-of-all-of0"></span> PrimitiveAllOfAllOf0**


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| PrimitiveAllOfAllOf0 | string| string | |  |  |



**<span id="primitive-all-of-all-of1"></span> PrimitiveAllOfAllOf1**


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| PrimitiveAllOfAllOf1 | uuid (formatted string)| strfmt.UUID | |  |  |


