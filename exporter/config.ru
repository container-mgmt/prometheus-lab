require 'awesome_print'
require 'rack'
require 'prometheus/middleware/collector'
require 'prometheus/middleware/exporter'
require 'prometheus/client'

use Rack::Deflater, if: ->(_, _, _, body) { body.any? && body[0].length > 512 }
# use Prometheus::Middleware::Collector
use Prometheus::Middleware::Exporter


prometheus = Prometheus::Client.registry

# create a new counter metric
my_metric = prometheus.gauge(:my_metric, 'A gauge to hack with')
app = lambda do |req|
  s = req["REQUEST_PATH"][1..-1].split(":")
  my_metric.set({:some_label => s[0]}, s[1].to_i)
  [200, { 'Content-Type' => 'text/html' }, ["my_metric{:some_label=\"#{s[0]}\"} #{s[1].to_i}"]]
end

run app
