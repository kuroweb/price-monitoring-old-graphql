module GraphqlSchema
  module Objects
    module Products
      module UsedSofmapCrawlSettings
        module UpdateUsedSofmapCrawlSettingExcludeProduct
          class UpdateUsedSofmapCrawlSettingExcludeProductResultError < Base
            implements Interfaces::ResultBase

            field :error,
                  Unions::Products::UsedSofmapCrawlSettings::UpdateUsedSofmapCrawlSettingExcludeProduct::
                  UpdateUsedSofmapCrawlSettingExcludeProductResultErrorsUnion,
                  null: false
          end
        end
      end
    end
  end
end
