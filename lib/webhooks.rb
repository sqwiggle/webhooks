require 'goliath'

require 'webhooks/version'
require 'webhooks/server'

module Webhooks
  def self.run!
    runner = Goliath::Runner.new(ARGV, nil)
    runner.api = Server.new
    runner.app = Goliath::Rack::Builder.build(Server, runner.api)
    runner.run
  end
end
