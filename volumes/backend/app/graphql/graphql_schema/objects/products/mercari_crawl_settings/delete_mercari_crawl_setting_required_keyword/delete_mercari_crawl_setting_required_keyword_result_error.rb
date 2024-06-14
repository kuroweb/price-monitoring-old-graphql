module GraphqlSchema
  module Objects
    module Products
      module MercariCrawlSettings
        module DeleteMercariCrawlSettingRequiredKeyword
          class DeleteMercariCrawlSettingRequiredKeywordResultError < Base
            implements Interfaces::Products::ResultBase

            field :error,
                  Unions::Products::MercariCrawlSettings::DeleteMercariCrawlSettingRequiredKeyword::
                  DeleteMercariCrawlSettingRequiredKeywordResultErrors,
                  null: false
          end
        end
      end
    end
  end
end
