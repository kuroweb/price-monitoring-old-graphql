module GraphqlSchema
  module Mutations
    module Products
      class CreateYahooAuctionCrawlSettingExcludeKeyword < Base
        input_object_class InputObjects::Products::CreateYahooAuctionCrawlSettingExcludeKeywordInput

        field :result, Unions::Products::CreateYahooAuctionCrawlSettingExcludeKeywordResult, null: false

        type Unions::Products::CreateYahooAuctionCrawlSettingExcludeKeywordResult

        def resolve(input)
          product = Product.find(input[:product_id])
          yahoo_auction_crawl_setting_exclude_keyword =
            product
            .yahoo_auction_crawl_setting
            .yahoo_auction_crawl_setting_exclude_keywords
            .create!(keyword: input[:keyword])

          {
            __typename: "CreateYahooAuctionCrawlSettingExcludeKeywordResultSuccess",
            yahoo_auction_crawl_setting_exclude_keyword:,
            ok: true
          }
        rescue StandardError => e
          handle_error(e)
        end

        private

        def handle_error(exception)
          case exception
          when ActiveRecord::RecordInvalid
            Rails.logger.error("Bad Request. exception: #{exception.full_message}")
            error_response("400", "Bad Request.")
          when ActiveRecord::RecordNotFound
            error_response("404", "Not Found.")
          when ActiveRecord::RecordNotUnique
            error_response("409", "Conflict.")
          else
            Rails.logger.error("Bad Request. exception: #{exception.full_message}")
            error_response("503", "Internal Server Error.")
          end
        end

        def error_response(code, message)
          {
            __typename: "CreateYahooAuctionCrawlSettingExcludeKeywordResultError",
            error: {
              __typename: "CreateYahooAuctionCrawlSettingExcludeKeywordResultValidationFailed",
              code:,
              message:,
              details: []
            },
            ok: false
          }
        end
      end
    end
  end
end
