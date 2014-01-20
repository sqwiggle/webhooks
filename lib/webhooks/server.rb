module Webhooks
  class Server < Goliath::API
    def response(env)
      [200, {}, 'Hello World']
    end
  end
end
