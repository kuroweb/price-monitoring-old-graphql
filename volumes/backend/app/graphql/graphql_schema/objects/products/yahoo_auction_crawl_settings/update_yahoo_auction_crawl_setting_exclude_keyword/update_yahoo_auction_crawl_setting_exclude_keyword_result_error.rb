module GraphqlSchema
  module Objects
    module Products
      module YahooAuctionCrawlSettings
        module UpdateYahooAuctionCrawlSettingExcludeKeyword
          class UpdateYahooAuctionCrawlSettingExcludeKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::YahooAuctionCrawlSettings::UpdateYahooAuctionCrawlSettingExcludeKeyword::
                  UpdateYahooAuctionCrawlSettingExcludeKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
