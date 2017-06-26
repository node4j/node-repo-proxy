# Maven Proxy for Node downloads

This project implements a maven repository that proxies node distribution
downloads. The aim is that it can be used by various Java build plugins
that downloads Node and uses Node in build processes.

It's deployed for now on a Heroku instance and can be accessed here:

* https://repo.node4j.org/

## Node download

It will automatically serve metadata (latest node versions).

* https://repo.node4j.org/org/node/node/maven-metadata.xml

To download version 7.6.0 (linux x86, tar.gz) the GAV-coordinate would be:

* `org.node:node:7.6.0:linux-x86@tar.gz`

And the generated URL would then be:

* https://repo.node4j.org/org/node/node/7.6.0/node-7.6.0-linux-x86.tar.gz

## NPM download

It will also serve tgz bundles and meta data from NPM registry. For example,
to download yarn (`org.npmjs:yarn:0.25.3@tgz`):

* https://repo.node4j.org/org/npmjs/yarn/0.25.3/yarn-0.25.3.tgz

The meta-data for NPM packages are also available:

* https://repo.node4j.org/org/npmjs/yarn/maven-metadata.xml

## What's missing

* Testing.
* SHA and MD5 from distributions.
