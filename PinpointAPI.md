### /removeApplicationName
* Parameter
 * `applicationName`

### /removeAgentId
* Parameter
 * `applicationName`
 * `agentId`

### /removeInactiveAgents
* Parameter
 * `durationDays`

### /agentIdMap

Return all agents map.

* Parameter: None
* Result:
```json
{
"ajgw_batserver.035f2bf7": [
{
"applicationName": "ajgw_batserver",
"serviceType": "TOMCAT",
"code": 1010
}
],
"ajgw_batserver.07d37f68": [
{
"applicationName": "ajgw_batserver",
"serviceType": "TOMCAT",
"code": 1010
}
],
"ajgw_batserver.0813bcf9": [
{
"applicationName": "ajgw_batserver",
"serviceType": "TOMCAT",
"code": 1010
}
]
}
```

### /duplicateAgentIdMap
Return agents with same ID
* Parameter: None

### /getInactiveAgents
Return currently inactive agents map.

* Parameter
 * `applicationName`
 * `durationDays`