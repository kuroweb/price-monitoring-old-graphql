module GraphqlSchema
  module Objects
    module Products
      module MercariCrawlSettings
        module CreateMercariCrawlSettingExcludeKeyword
          class CreateMercariCrawlSettingExcludeKeywordResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::MercariCrawlSettings::CreateMercariCrawlSettingExcludeKeyword::
                  CreateMercariCrawlSettingExcludeKeywordResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
