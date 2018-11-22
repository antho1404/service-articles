## service-articles [![CircleCI](https://img.shields.io/circleci/project/github/ilgooz/service-articles.svg)](https://github.com/ilgooz/service-articles) [![codecov](https://codecov.io/gh/ilgooz/service-articles/branch/master/graph/badge.svg)](https://codecov.io/gh/ilgooz/service-articles)
A MESG service to create and manage articles.


```bash
mesg-core service deploy https://github.com/ilgooz/service-articles
```



# Tasks

## create

Task key: `create`

Create an article

### Inputs

| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **article** | `Object` | An article to create. |


### Outputs

##### error

Output key: `error`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **message** | `String` |  |

##### success

Output key: `success`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **article** | `Object` | An article. |




## get

Task key: `get`

Create an article

### Inputs

| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **id** | `String` | ID or human readable id of article. |


### Outputs

##### error

Output key: `error`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **message** | `String` |  |

##### success

Output key: `success`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **article** | `Object` | Requested Article |




## list

Task key: `list`

Create an article



### Outputs

##### error

Output key: `error`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **message** | `String` |  |

##### success

Output key: `success`



| **Key** | **Type** | **Description** |
| --- | --- | --- |
| **articles** | `Object` | List of articles. |




