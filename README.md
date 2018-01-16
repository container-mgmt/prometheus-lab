# prometheus-lab
Run a prometheus stack locally for research purposes. The stack includes:

- A smart server exporting a /metrics endpoint
- Prometheus built from the release-2.0 branch  
- Alertmanager built from the release-0.12 branch  
- FakeSMTP build from master

## Questions

What happens to firing alerts in prometheus after restart[1]?

There are no alerts after restart. Method:
- trigger alert
- restart prometheus
- view alerts screen

What happens to firing alerts in prometheus after configuration reload?

They disappear from the firing alerts list
- remove definition in rules.yml file
- reload configuration using `curl -X POST http://localhost:9090/-/reload`
- view alerts screen

What happens to firing alerts in prometheus after removing their definition and running configuration reload?
They are removed from prometheus after a while. Controlled by a setting resolve_timeout

What happens to active alerts in alermanager after restart?
Alerts are not persisted

http_requests_total{code="200",instance="localhost:9093",job="alertmanager"}
