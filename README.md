# Maven Proxy for Node downloads

This project implements a maven repository that proxies node distribution
downloads. The aim is that it can be used by various Java build plugins
that downloads Node and uses Node in build processes.

It's deployed for now on Heroku free instance and can me accessed here:

* https://node-repo-proxy.herokuapp.com/

It will automatically serve metadata (latest node versions). See the
following URL:

* https://node-repo-proxy.herokuapp.com/org/node/node/maven-metadata.xml

To download version 7.6.0 (linux x86, tar.gz) the GAV-coordinate would be:

* `org.node:node:7.6.0:linux-x86@tar.gz`

## What's missing

* Testing.
* SHA and MD5 from distributions.
