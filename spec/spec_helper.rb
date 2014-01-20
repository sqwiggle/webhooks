require 'bundler'
require 'bundler/setup'

require 'webhooks'

require 'goliath/test_helper'

Goliath.env = :test

RSpec.configure do |config|
  config.include Goliath::TestHelper,
    example_group: { file_path: /spec\/integration/ }

  config.color = true
  config.order = :random
end
