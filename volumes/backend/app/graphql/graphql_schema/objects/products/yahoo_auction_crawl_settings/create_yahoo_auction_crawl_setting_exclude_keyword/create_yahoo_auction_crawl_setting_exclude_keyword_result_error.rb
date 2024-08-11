module GraphqlSchema
  module Objects
    module Products
      module YahooAuctionCrawlSettings
        module CreateYahooAuctionCrawlSettingExcludeKeyword
          class CreateYahooAuctionCrawlSettingExcludeKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::YahooAuctionCrawlSettings::CreateYahooAuctionCrawlSettingExcludeKeyword::
                  CreateYahooAuctionCrawlSettingExcludeKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
