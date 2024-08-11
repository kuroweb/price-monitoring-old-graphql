module GraphqlSchema
  module Objects
    module Products
      module YahooAuctionCrawlSettings
        module DeleteYahooAuctionCrawlSettingExcludeProduct
          class DeleteYahooAuctionCrawlSettingExcludeProductResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::YahooAuctionCrawlSettings::DeleteYahooAuctionCrawlSettingExcludeProduct::
                  DeleteYahooAuctionCrawlSettingExcludeProductResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
