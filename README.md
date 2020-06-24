# DNS load tester

[<img src="ll-logo.png">](https://lablabs.io/)

We help companies build, run, deploy and scale software and infrastructure by embracing the right technologies and principles. Check out our website at https://lablabs.io/

---

## Description

This tool is used to execute many dns lookup queries in parallel. It is useful when debugging DNS related issues.

# Build using docker

```
docker build -t dns-load-test .
```

## Usage

```
docker run --network=host -it dns-load-test google.com cloudflare.com
```

### Optional args
```
  -threads int
        how many threads in parallel to run (default 5)
```

# Docker images

```
docker pull lablabs/dns-load-test:0.0.1
```

# Kubernetes Deployment

This tool can be deployed in Kubernetes with the configuration below. The log output of individual pods can be that collected using various methods (fluentd into ElasticSearch, etc.)

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dns-load-test
  namespace: default
  labels:
    app: dns-load-test
spec:
  replicas: 1
  selector:
    matchLabels:
      app: dns-load-test
  template:
    metadata:
      labels:
        app: dns-load-test
    spec:
      containers:
      - name: dns-load-test
        image: lablabs/dns-load-test:0.0.1
        imagePullPolicy: Always
        args: ["-threads", "50", "google.com", "cloudflare.com", "gitlab-dev-ec.ic2xx5.ng.0001.euw1.cache.amazonaws.com"]
      dnsPolicy: "None"
      dnsConfig:
        nameservers:
          - 100.64.255.31
        options:
          - name: "use-vc"
          - name: "timeout"
            value: "1"
          - name: "attempts"
            value: "1"
```

## Contributing and reporting issues

Feel free to create an issue in this repository if you have questions, suggestions or feature requests.

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

See [LICENSE](LICENSE) for full details.

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
