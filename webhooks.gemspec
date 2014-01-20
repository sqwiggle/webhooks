# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'webhooks/version'

Gem::Specification.new do |spec|
  spec.name          = 'webhooks'
  spec.version       = Webhooks::VERSION
  spec.authors       = ['Luke Roberts']
  spec.email         = ['email@luke-roberts.co.uk']
  spec.description   = %q{A webhook engine}
  spec.summary       = %q{Webhook Engine}
  spec.homepage      = ''
  spec.license       = 'MIT'

  spec.files         = `git ls-files`.split($/)
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ['lib']

  spec.executables << 'webhooks_server'

  spec.add_dependency 'goliath'

  spec.add_development_dependency 'bundler', '~> 1.3'
  spec.add_development_dependency 'rake'
  spec.add_development_dependency 'rspec'
  spec.add_development_dependency 'em-http-request'
end
