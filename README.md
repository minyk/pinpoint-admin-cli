DC/OS CLI Subcommand for mesos resource reserve/unreserve
==========================================

**CAUTION: This is not battle-hardened yet. USE AT YOUR OWN RISK.** 

# pinpoint-cli

# Examples

### Get
```
$ ./pinpoint-cli -v  --pinpoint="http://pinpoint.example.com" --pinpoint-password="admin" get
Application Name: *
  Application test_contract has 15 agents.
  Service type TOMCAT has 15 agents.

```

### Remove
```
$ ./pinpoint-cli -v  --pinpoint="http://pinpoint.example.com" --pinpoint-password="admin" remove test_contract --inactive-only
Found 10 inactive agents for test_contract.
Agent test_contract.91d6430d is successfully removed.
Agent test_contract.9b0103c2 is successfully removed.
Agent test_contract.a05898f6 is successfully removed.
Agent test_contract.1641dd6e is successfully removed.
Agent test_contract.76b6b141 is successfully removed.
Agent test_contract.b3eedb07 is successfully removed.
Agent test_contract.bed4e916 is successfully removed.
Agent test_contract.f7c908e1 is successfully removed.
Agent test_contract.026c2504 is successfully removed.
Agent test_contract.09e2baf5 is successfully removed.

```

# Build

```
$ make build
```

# Acknowledgement

* The client code is heavily adopted from https://github.com/mesosphere/dcos-commons/tree/master/cli
